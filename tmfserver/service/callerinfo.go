package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"

	"github.com/DOME-Marketplace/tmfgo/internal/errl"
	"github.com/DOME-Marketplace/tmfgo/internal/jpath"
	"github.com/DOME-Marketplace/tmfgo/types"
	"github.com/go-viper/mapstructure/v2"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/greyxor/slogor"
)

// ProcessAccessToken verifies the Access Token received from the caller and
// creates a map ready to be passed to the rules engine.
//
// The access token may not exist, but if it does then it must be valid.
// For convenience of the policies, some calculated fields are created and returned in the 'user' object.
func (svc *Service) ProcessAccessToken(accessToken string) (user *types.AuthUser, err error) {

	// The zero AuthUser is a valid guest user.
	authUser := &types.AuthUser{}

	// An empty token is not considered an error, and the caller should enforce its existence if needed
	if len(accessToken) == 0 {
		return authUser, nil
	}

	// If we receive a superadmin token, create an AuthUser will all powers
	if accessToken == svc.adminToken {

		// The user is the server operator
		authUser.CommonName = svc.ServerOperatorName
		authUser.Country = svc.ServerOperatorCountry
		authUser.EmailAddress = svc.ServerEmailAddress
		authUser.Organization = svc.ServerOperatorName
		authUser.OrganizationIdentifier = svc.ServerOperatorOrganizationIdentifier
		authUser.SerialNumber = svc.ServerOperatorOrganizationIdentifier

		// With all powers and owning all objects
		authUser.IsAuthenticated = true
		authUser.IsLEAR = true
		authUser.IsOwner = true
		authUser.ProductCreatePower = true
		authUser.ProductUpdatePower = true
		authUser.ProductDeletePower = true

		authUser.TokenMap = make(map[string]any)
		authUser.TokenMap["tokenType"] = ISBEAccessToken
		authUser.TokenMap["user_identifier"] = authUser.SerialNumber
		authUser.TokenMap["organization"] = authUser.Organization
		authUser.TokenMap["organization_identifier"] = authUser.OrganizationIdentifier
		authUser.TokenMap["name"] = authUser.CommonName
		authUser.TokenMap["country"] = authUser.Country
		authUser.TokenMap["email"] = authUser.EmailAddress
		authUser.TokenMap["serial_number"] = authUser.SerialNumber

		return authUser, nil

	}

	// It is an error to send an invaild token with the request, so we have to verify it.
	// We verify the token and extract the claims, a verification error stops processing.

	var token *jwt.Token
	var theClaims = jwt.MapClaims{}

	if svc.Features.VerifyJWTSignature {

		// This is called by ParseWithClaims to retrieve the verification key
		verifierPublicKeyFunc := func(tok *jwt.Token) (any, error) {

			// Check that the configuration for retrieving the JWK is present.
			if svc.oid == nil {
				return nil, errl.Errorf("openid support not initialized")
			}
			// The key ID is used to retrieve the verification key from the OpenID Provider
			keyID, ok := tok.Header["kid"].(string)
			if !ok {
				return nil, errl.Errorf("invalid access token: kid not found in header")
			}
			// Get the verification key from the OpenID Provider (it is cached locally)
			vk, err := svc.oid.VerificationJWKKey(keyID)
			if err != nil {
				return nil, errl.Error(err)
			}
			slog.Debug("publicKeyFunc", "key", vk)
			return vk.Key, nil
		}

		// Validate and verify the token
		token, err = jwt.NewParser().ParseWithClaims(accessToken, theClaims, verifierPublicKeyFunc)
		if err != nil {
			slog.Error("invalid access token", slogor.Err(err), "token", accessToken)
			return nil, errl.Errorf("invalid access token: %w, token: %s", err, accessToken)
		}

	} else {

		// Parse the token without signature verification
		token, _, err = new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})
		if err != nil {
			slog.Error("invalid access token", slogor.Err(err), "token", accessToken)
			return nil, errl.Errorf("invalid access token: %w, token: %s", err, accessToken)
		}

	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		slog.Error("JWT claims are not of type MapClaims")
		return nil, errors.New("invalid JWT claims format")
	}

	// The actual claims depends on the caller: the standard DOME or the ISBE one, which does not send the "vc" claim.
	if svc.IsDOME() {
		return svc.processDOMEAccessToken(claims, accessToken)
	} else {
		return svc.processISBEAccessToken(claims, accessToken)
	}

}

type accessTokenType string

const (
	DOMEAccessToken accessTokenType = "DOME"
	ISBEAccessToken accessTokenType = "ISBE"
)

func (svc *Service) processDOMEAccessToken(claims jwt.MapClaims, accessToken string) (user *types.AuthUser, err error) {

	authUser := &types.AuthUser{}
	authUser.IsAuthenticated = true

	// Extract the Verifiable Credential from the claims
	verifiableCredential := jpath.GetMap(claims, "vc")
	if len(verifiableCredential) == 0 {
		return nil, errl.Errorf("access token without 'vc': %s", accessToken)
	}

	credentialSubject := jpath.GetMap(verifiableCredential, "credentialSubject")
	if len(credentialSubject) == 0 {
		return nil, errl.Errorf("access token without 'credentialSubject': %s", accessToken)
	}

	mandateData := jpath.GetMap(credentialSubject, "mandate")
	if len(mandateData) == 0 {
		return nil, errl.Errorf("access token without 'mandate': %s", accessToken)
	}

	mandatorData := jpath.GetMap(mandateData, "mandator")
	if len(mandatorData) == 0 {
		return nil, errl.Errorf("access token without 'mandator': %s", accessToken)
	}

	// Marshal and unmarshal to AuthUser struct for type safety and JSON tag mapping
	mandatorJSON, err := json.Marshal(mandatorData)
	if err != nil {
		slog.Error("Failed to marshal mandator data", slog.Any("error", err))
		return nil, fmt.Errorf("failed to process mandator data: %w", err)
	}

	if err := json.Unmarshal(mandatorJSON, authUser); err != nil {
		slog.Error("Failed to unmarshal mandator data to AuthUser", slog.Any("error", err))
		return nil, fmt.Errorf("failed to process mandator data: %w", err)
	}

	// Verify that the critical fields exist (Organization, Organization Identifier and Country)
	// These are the minimum fields required to identify the caller for H2M and M2M flows.
	if len(authUser.Organization) == 0 {
		return nil, errl.Errorf("access token without 'organization': %s", accessToken)
	}
	if len(authUser.OrganizationIdentifier) == 0 {
		return nil, errl.Errorf("access token without 'organization_identifier': %s", accessToken)
	}
	if len(authUser.Country) == 0 {
		return nil, errl.Errorf("access token without 'country': %s", accessToken)
	}

	slog.Debug("Successfully parsed AuthUser from JWT",
		slog.String("organizationIdentifier", authUser.OrganizationIdentifier),
		slog.String("country", authUser.Country))

	claims["tokenType"] = DOMEAccessToken

	authUser.AccessToken = accessToken
	authUser.TokenMap = claims

	// Parse the user powers to look for the ones we only care about
	authUserPowers := jpath.GetList(verifiableCredential, "credentialSubject.mandate.power")

	return svc.procesPowers(authUserPowers, authUser), nil

}

func (svc *Service) processISBEAccessToken(claims jwt.MapClaims, accessToken string) (user *types.AuthUser, err error) {
	authUser := &types.AuthUser{}
	authUser.IsAuthenticated = true

	// Get the organization data

	authUser.Organization = jpath.GetString(claims, "organization")
	if len(authUser.Organization) == 0 {
		return nil, errl.Errorf("access token without 'organization': %s", accessToken)
	}

	authUser.OrganizationIdentifier = jpath.GetString(claims, "organization_identifier")
	if len(authUser.OrganizationIdentifier) == 0 {
		return nil, errl.Errorf("access token without 'organization_identifier': %s", accessToken)
	}

	authUser.CommonName = jpath.GetString(claims, "name")
	if len(authUser.CommonName) == 0 {
		return nil, errl.Errorf("access token without 'name': %s", accessToken)
	}

	authUser.SerialNumber = jpath.GetString(claims, "user_identifier")
	if len(authUser.SerialNumber) == 0 {
		return nil, errl.Errorf("access token without 'user_identifier': %s", accessToken)
	}

	// TODO: the token from ISBE should contain the country. Until they fix it upstream, we use the default value "ES".
	authUser.Country = jpath.GetString(claims, "country")
	if len(authUser.Country) == 0 {
		slog.Debug("access token does not contain 'country' or it's not a string", "token", accessToken)
		authUser.Country = "ES"
	}

	authUser.EmailAddress = jpath.GetString(claims, "email")
	if len(authUser.EmailAddress) == 0 {
		return nil, errl.Errorf("access token without 'email': %s", accessToken)
	}

	claims["tokenType"] = ISBEAccessToken

	authUser.AccessToken = accessToken
	authUser.TokenMap = claims

	authUserPowers := jpath.GetList(claims, "power")
	if len(authUserPowers) == 0 {
		return nil, errl.Errorf("access token without 'power': %s", accessToken)
	}

	return svc.procesPowers(authUserPowers, authUser), nil
}

func (svc *Service) procesPowers(authUserPowers []any, authUser *types.AuthUser) *types.AuthUser {

	// Parse the user powers for the ones we only care about
	for _, p := range authUserPowers {
		var userPower types.OnePower
		if err := mapstructure.Decode(p, &userPower); err != nil {
			slog.Error("error decoding power", slogor.Err(err))
			continue
		}

		if userPower.Includes(svc.LEARPower) {
			authUser.IsLEAR = true
			// A LEAR can create, update and delete product offerings
			authUser.ProductCreatePower = true
			authUser.ProductUpdatePower = true
			authUser.ProductDeletePower = true
		} else {
			if userPower.Includes(svc.ProductCreatePower) {
				authUser.ProductCreatePower = true
			}
			if userPower.Includes(svc.ProductUpdatePower) {
				authUser.ProductUpdatePower = true
			}
			if userPower.Includes(svc.ProductDeletePower) {
				authUser.ProductDeletePower = true
			}
		}

	}

	return authUser
}
