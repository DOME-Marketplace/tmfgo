package types

import (
	"encoding/json"
	"log/slog"
	"slices"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hesusruiz/tmforum/internal/errl"
)

type Mandate struct {
	Id       string `json:"id,omitempty"`
	Mandator struct {
		OrganizationIdentifier string `json:"organizationIdentifier,omitempty"` // OID 2.5.4.97
		CommonName             string `json:"commonName,omitempty"`             // OID 2.5.4.3
		GivenName              string `json:"givenName,omitempty"`
		Surname                string `json:"surname,omitempty"`
		EmailAddress           string `json:"emailAddress,omitempty"`
		SerialNumber           string `json:"serialNumber,omitempty"`
		Organization           string `json:"organization,omitempty"`
		Country                string `json:"country,omitempty"`
	} `json:"mandator"`
	Mandatee struct {
		Id           string `json:"id,omitempty"`
		FirstName    string `json:"firstName,omitempty"`
		LastName     string `json:"lastName,omitempty"`
		Gender       string `json:"gender,omitempty"`
		Email        string `json:"email,omitempty"`
		Mobile_phone string `json:"mobile_phone,omitempty"`
	} `json:"mandatee"`
	Power []struct {
		Id           string   `json:"id,omitempty"`
		Tmf_type     string   `json:"type,omitempty"`
		Tmf_domain   string   `json:"domain,omitempty"`
		Tmf_function string   `json:"function,omitempty"`
		Tmf_action   []string `json:"action,omitempty"`
	} `json:"power"`
}

type CredentialIssuer struct {
	Id                     string `json:"id,omitempty"`
	OrganizationIdentifier string `json:"organizationIdentifier,omitempty"`
	Organization           string `json:"organization,omitempty"`
	Country                string `json:"country,omitempty"`
	SerialNumber           string `json:"serialNumber,omitempty"`
	CommonName             string `json:"commonName,omitempty"`
}

type OnePower struct {
	Id       string   `json:"id,omitempty" yaml:"id,omitempty" mapstructure:"id"`
	Type     string   `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type"`
	Domain   string   `json:"domain,omitempty" yaml:"domain,omitempty" mapstructure:"domain"`
	Function string   `json:"function,omitempty" yaml:"function,omitempty" mapstructure:"function"`
	Action   []string `json:"action,omitempty" yaml:"action,omitempty" mapstructure:"action"`
}

func (p *OnePower) SameAs(other *OnePower) bool {
	if !strings.EqualFold(p.Type, other.Type) {
		return false
	}
	if !strings.EqualFold(p.Domain, other.Domain) {
		return false
	}
	if !strings.EqualFold(p.Function, other.Function) {
		return false
	}
	for i, action := range p.Action {
		if !strings.EqualFold(action, other.Action[i]) {
			return false
		}
	}
	return true
}

// Includes reports whether a given power p "includes" the supplied other power.
//
// It returns true iff all of the following conditions hold:
//   - p.type equals other.type (case-insensitive, via strings.EqualFold)
//   - p.domain equals other.domain (case-insensitive)
//   - p.function equals other.function (case-insensitive)
//   - every action in other power is present in the actions of p
//
// The method performs early returns and returns false on the first mismatch. It treats p.Action as a superset:
// extra actions present in p but not in other do not prevent inclusion. Note that the implementation assumes both
// p and other are non-nil; calling Includes with a nil receiver or nil other will result in a runtime panic.
func (p *OnePower) Includes(other OnePower) bool {
	// Check the lengths of the action arrays for a maximum length of 10
	if len(p.Action) > 10 || len(other.Action) > 10 {
		slog.Error("lenghts of action arrays are greater than 10", "p", len(p.Action), "other", len(other.Action))
		return false
	}

	// Type, Domain and Function must be the same
	if !strings.EqualFold(p.Type, other.Type) {
		return false
	}
	if !strings.EqualFold(p.Domain, other.Domain) {
		return false
	}
	if !strings.EqualFold(p.Function, other.Function) {
		return false
	}

	// Now we check the Action array.
	// If p.Action has an asterisc '*', this includes any action, so we return true
	if slices.Contains(p.Action, "*") {
		return true
	}

	// Check that each element of other.action is included in p.action
	// The comparison of individual elements is case-insensitive
	for _, otherAction := range other.Action {
		found := false
		for _, pAction := range p.Action {
			if strings.EqualFold(pAction, otherAction) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true

}

// Example of credentialStatus
//
//	"credentialStatus": {
//		"id": "https://issuer.dome-marketplace-sbx.org/backoffice/v1/credentials/status/1#SYC908RIQQqeUXR19nh3EQ",
//		"statusListCredential": "https://issuer.dome-marketplace-sbx.org/backoffice/v1/credentials/status/1",
//		"statusListIndex": "SYC908RIQQqeUXR19nh3EQ",
//		"statusPurpose": "revocation",
//		"type": "PlainListEntity"
//	},
type CredentialStatus struct {
	Id                   string `json:"id,omitempty"`
	StatusListCredential string `json:"statusListCredential,omitempty"`
	StatusListIndex      string `json:"statusListIndex,omitempty"`
	StatusPurpose        string `json:"statusPurpose,omitempty"`
	Type                 string `json:"type,omitempty"`
}

type LEARCredentialEmployee struct {
	Context []string `json:"@context,omitempty"`
	// CredentialStatus  CredentialStatus `json:"credentialStatus,omitempty"`
	Id             string   `json:"id,omitempty"`
	TypeCredential []string `json:"type,omitempty"`
	// Issuer            CredentialIssuer `json:"issuer,omitempty"`
	ValidFrom         string `json:"validFrom,omitempty"`
	ValidUntil        string `json:"validUntil,omitempty"`
	CredentialSubject struct {
		Mandate Mandate `json:"mandate"`
	} `json:"credentialSubject"`
}

type LEARCredentialEmployeeJWTClaims struct {
	LEARCredentialEmployee
	jwt.RegisteredClaims
}

func LEARCredentialFromMap(s map[string]any) (*LEARCredentialEmployee, error) {

	// Marshall to a byte array
	out, err := json.Marshal(s)
	if err != nil {
		return nil, errl.Errorf("error marshalling credential: %w", err)
	}

	// Unmarshal to a struct
	lc := &LEARCredentialEmployee{}
	err = json.Unmarshal(out, &lc)
	if err != nil {
		return nil, errl.Errorf("error unmarshalling credential: %w", err)
	}

	return lc, nil

}
