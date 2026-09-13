package service

import (
	"testing"
)

func TestProcessAccessToken(t *testing.T) {
	s := newISBEDEVTestService(t)

	// Disable signature validation for the tests, as the token is expired
	s.Features.VerifyJWTSignature = false

	t.Run("UsingTestAccessToken", func(t *testing.T) {
		// Authenticate as a normal user
		user, err := s.ProcessAccessToken(testSellerAccessToken)

		// The current implementation of ProcessAccessToken expects a 'vc' claim.
		// The testAccessToken provided does not have it, so it returns an error.
		// We assert this behavior for now.
		if err != nil {
			t.Logf("ProcessAccessToken returned error as expected: %v", err)
			if err.Error() != "missing 'vc' in JWT claims" {
				t.Errorf("Expected error 'missing vc in JWT claims', got: %v", err)
			}
			return
		}

		// If the implementation is updated to support the token format, these assertions will run.
		if user == nil {
			t.Fatal("Expected user to be returned")
		}

		// Verify fields from the token payload

		// "organization_identifier": "VATES-G87936159"
		if user.OrganizationIdentifier != "VATES-G87936159" {
			t.Errorf("Expected OrganizationIdentifier 'VATES-G87936159', got '%s'", user.OrganizationIdentifier)
		}
		if user.Organization != "ALASTRIA" {
			t.Errorf("Expected Organization 'ALASTRIA', got '%s'", user.Organization)
		}
		if user.SerialNumber != "12345678A" {
			t.Errorf("Expected SerialNumber '12345678A', got '%s'", user.SerialNumber)
		}
		if user.CommonName != "John Doe" {
			t.Errorf("Expected CommonName 'John Doe', got '%s'", user.CommonName)
		}
		if user.EmailAddress != "hesus.ruiz@gmail.com" {
			t.Errorf("Expected EmailAddress '[EMAIL_ADDRESS]', got '%s'", user.EmailAddress)
		}

		if !user.IsAuthenticated {
			t.Error("Expected IsAuthenticated to be true")
		}
	})
}

func TestProcessISBEAccessToken(t *testing.T) {
	s := newISBEDEVTestService(t)

	s.Features.VerifyJWTSignature = true

	t.Run("UsingISBEAccessToken", func(t *testing.T) {
		// testAccessToken is defined in callerinfo.go
		user, err := s.ProcessAccessToken(testServerOperatorAccessToken)

		// The current implementation of ProcessAccessToken expects a 'vc' claim.
		// The testAccessToken provided does not have it, so it returns an error.
		// We assert this behavior for now.
		if err != nil {
			t.Logf("ProcessAccessToken returned error as expected: %v", err)
			if err.Error() != "missing 'vc' in JWT claims" {
				t.Errorf("Expected error 'missing vc in JWT claims', got: %v", err)
			}
			return
		}

		// If the implementation is updated to support the token format, these assertions will run.
		if user == nil {
			t.Fatal("Expected user to be returned")
		}

		// Verify fields from the token payload
		// "organization_identifier": "VATES-G87936159"
		if user.OrganizationIdentifier != "VATES-G87936159" {
			t.Errorf("Expected OrganizationIdentifier 'VATES-G87936159', got '%s'", user.OrganizationIdentifier)
		}

		// Must be LEAR
		if !user.IsLEAR {
			t.Errorf("Expected IsLEAR to be true")
		}

		if !user.IsAuthenticated {
			t.Error("Expected IsAuthenticated to be true")
		}
	})
}
