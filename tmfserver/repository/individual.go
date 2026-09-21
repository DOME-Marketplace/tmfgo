package repository

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/DOME-Marketplace/tmfgo/internal/errl"
	"github.com/DOME-Marketplace/tmfgo/types"
)

func TMFIndividualFromCredential(verifiableCredential map[string]any, user *Organization) (*TMFRecord, error) {

	// TODO: what happens if we receive a LEARCredentialMachine?
	lc, err := types.LEARCredentialFromMap(verifiableCredential)
	if err != nil {
		return nil, errl.Errorf("error creating LEARCredentialEmployee struct: %w", err)
	}

	now := time.Now()
	lastUpdate := now.Format(time.RFC3339Nano)

	userId := lc.CredentialSubject.Mandate.Mandatee.Id
	if userId == "" {
		return nil, errl.Errorf("id field for mandatee is empty")
	}

	// The format is "urn:ngsi-ld:individual:{organizationIdentifier}:{uuid}"
	id := fmt.Sprintf("urn:ngsi-ld:individual:%s:%s", user.OrganizationIdentifier, userId)

	objectType := types.Individual
	version := "1.0"

	// Prepare contactMedium
	var contactMedium []any
	if lc.CredentialSubject.Mandate.Mandatee.Email != "" {
		contactMedium = append(contactMedium, map[string]any{
			"@type":        "EmailContactMedium",
			"preferred":    true,
			"emailAddress": user.EmailAddress,
		})
	}

	// Check if the powers are "Onboarding"
	hasOnboardingPower := false
	powers := lc.CredentialSubject.Mandate.Power
	for _, p := range powers {
		if strings.EqualFold(p.Tmf_type, "Domain") && strings.EqualFold(p.Tmf_domain, "DOME") {

			if strings.EqualFold(p.Tmf_function, "Onboarding") {

				for _, action := range p.Tmf_action {
					if strings.EqualFold(action, "execute") {
						hasOnboardingPower = true
						break
					}
				}

			}

		}
	}

	relatedParty := []any{}
	if hasOnboardingPower {
		relatedParty = append(relatedParty, map[string]any{
			"@referredType": nil,
			"href":          "urn:ngsi-ld:organization:" + user.OrganizationIdentifier,
			"id":            "urn:ngsi-ld:organization:" + user.OrganizationIdentifier,
			"name":          user.Organization,
			"role":          "orgAdmin",
		},
		)
	}

	indMap := map[string]any{
		"@type":      types.Individual,
		"id":         id,
		"href":       id,
		"version":    version,
		"lastUpdate": lastUpdate,
		"externalReference": []any{
			map[string]any{
				"externalReferenceType": "idm_id",
				"name":                  lc.CredentialSubject.Mandate.Mandatee.Id,
			},
		},
		"familyName": lc.CredentialSubject.Mandate.Mandatee.LastName,
		"givenName":  lc.CredentialSubject.Mandate.Mandatee.FirstName,
		"gender":     lc.CredentialSubject.Mandate.Mandatee.Gender,

		"contactMedium": contactMedium,

		"partyCharacteristic": []any{
			map[string]any{
				"name":  "country",
				"value": user.Country,
			},
		},

		"relatedParty": relatedParty,
	}

	theIdentification := map[string]any{
		"@type":              "individualIdentification",
		"identificationId":   lc.CredentialSubject.Mandate.Mandatee.Id,
		"identificationType": "learcredentialemployee",
		"issuingAuthority":   lc.CredentialSubject.Mandate.Mandator.OrganizationIdentifier,
	}

	indMap["individualIdentification"] = []any{theIdentification}

	content, err := json.Marshal(indMap)
	if err != nil {
		return nil, errl.Errorf("error marshalling organization: %w", err)
	}

	org := &TMFRecord{
		ID:         id,
		Type:       objectType,
		Version:    version,
		APIVersion: "v4",
		LastUpdate: lastUpdate,
		Content:    content,
		CreatedAt:  now.Unix(),
		UpdatedAt:  now.Unix(),
	}

	return org, nil
}
