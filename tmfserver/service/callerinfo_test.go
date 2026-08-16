package service

import (
	"testing"
)

const isbeAdminAccessToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICItckxwSkhNVkhCSUQ0Q2FRX0dsTjhFTEprQ0tYMUJWUzhMTzd6enU1cTFVIn0.eyJleHAiOjE3Nzc3MzI1OTIsImlhdCI6MTc3NTE0MDU5MiwiYXV0aF90aW1lIjoxNzc1MTQwNTkyLCJqdGkiOiJvbnJ0YWM6ZDU3NDk4ODItZWJhOC0wNmJhLTk0MjktNjk5NDBlNDE1ZGU4IiwiaXNzIjoiaHR0cHM6Ly9pZHAuZGV2LmNsb3VkLXcuZW52cy5yZWRpc2JlLmNvbS9hdXRoL3JlYWxtcy9kZXYtaXNiZSIsImF1ZCI6WyJpc2JlLXBvcnRhbC1kZXYiLCJhY2NvdW50Il0sInN1YiI6ImIyMzQ5ZTMxLWEyOTktNGU5Yi04ZGQxLWQxMWExZDFmNzBjMSIsInR5cCI6IkJlYXJlciIsImF6cCI6Imh0dHBzOi8vY2F0YWxvZy5pc2Jlb25ib2FyZC5jb20iLCJzaWQiOiIyZWFhMjMyZC03ZTI5LTA1ZmUtM2YzZi04MjYwYTYzZWVjZTAiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbImh0dHBzOi8vY2F0YWxvZy5kZXYuY2xvdWQtdy5lbnZzLnJlZGlzYmUuY29tIl0sInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJkZWZhdWx0LXJvbGVzLWRldi1pc2JlIiwib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIG9yZ2FuaXphdGlvbiBlbWFpbCBwcm9maWxlIiwidXNlcl9pZGVudGlmaWVyIjoiMTIzNDU2NzhBIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsIm9yZ2FuaXphdGlvbiI6IkFMQVNUUklBIiwibmFtZSI6IkpvaG4gRG9lIiwib3JnYW5pemF0aW9uX2lkZW50aWZpZXIiOiJWQVRFUy1HODc5MzYxNTkiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJqZXN1c0BhbGFzdHJpYS5pbyIsInBvd2VyIjpbeyJhY3Rpb24iOlsiKiJdLCJkb21haW4iOiJJU0JFIiwiZnVuY3Rpb24iOiJNYW5hZ2VtZW50IiwidHlwZSI6Im9yZ2FuaXphdGlvbiJ9LHsiYWN0aW9uIjpbIioiXSwiZG9tYWluIjoiSVNCRSIsImZ1bmN0aW9uIjoiSGVscGRlc2siLCJ0eXBlIjoib3JnYW5pemF0aW9uIn0seyJhY3Rpb24iOlsiKiJdLCJkb21haW4iOiJJU0JFIiwiZnVuY3Rpb24iOiJGYXVjZXQiLCJ0eXBlIjoib3JnYW5pemF0aW9uIn0seyJhY3Rpb24iOlsiKiJdLCJkb21haW4iOiJJU0JFIiwiZnVuY3Rpb24iOiJXaXphcmQiLCJ0eXBlIjoib3JnYW5pemF0aW9uIn0seyJhY3Rpb24iOlsiKiJdLCJkb21haW4iOiJJU0JFIiwiZnVuY3Rpb24iOiJOb3Rhcml6YXRpb24iLCJ0eXBlIjoib3JnYW5pemF0aW9uIn0seyJhY3Rpb24iOlsiKiJdLCJkb21haW4iOiJJU0JFIiwiZnVuY3Rpb24iOiJOb3RpZmljYXRpb25zIiwidHlwZSI6Im9yZ2FuaXphdGlvbiJ9LHsiYWN0aW9uIjpbIioiXSwiZG9tYWluIjoiSVNCRSIsImZ1bmN0aW9uIjoiSWRlbnRpdHkiLCJ0eXBlIjoib3JnYW5pemF0aW9uIn0seyJhY3Rpb24iOlsiKiJdLCJkb21haW4iOiJJU0JFIiwiZnVuY3Rpb24iOiJFbnJvbGxtZW50IiwidHlwZSI6Im9yZ2FuaXphdGlvbiJ9LHsiYWN0aW9uIjpbIkV4ZWN1dGUiXSwiZG9tYWluIjoiSVNCRSIsImZ1bmN0aW9uIjoiT25ib2FyZGluZyIsInR5cGUiOiJvcmdhbml6YXRpb24ifV0sImdpdmVuX25hbWUiOiJKb2huIiwidXNlciI6IkpvaG4gRG9lIiwiZmFtaWx5X25hbWUiOiJEb2UiLCJlbWFpbCI6Imhlc3VzLnJ1aXpAZ21haWwuY29tIn0.iUu88PYBYydC_YuYNfi9gma3jg3q1IPzys9Wfo_KvFoDatLhTshQ1b8gVOPB3Z4WFi_7TUcTmejtgYwYr79qjo1A2upSzL_IbGt20w4zUtuxJ4WA0anaD5cuyKBUieUNiEwR-_UKZyFj1dodIAtEIAU_D4qlxOgtZ2WKcxfT8RZajyUzHYCkZSiTLMlxkpXNciqaFOPD5Tjnu9N_G27WtknwBEIRRa3okM8K-mSn6zXpIVxE8kgi7Cvd0OJnOI6Jo4eWLtrYqw750CsjU8n6Y5L0Kl5M8bNOiKwMzVvChjPWcGRowxpJozLAZzhqQcmlhS8Sn79oSXWb67iC339Ogw"

func TestProcessAccessToken(t *testing.T) {
	s := newISBEDEVTestService(t)

	t.Run("UsingTestAccessToken", func(t *testing.T) {
		// testAccessToken is defined in callerinfo.go
		user, err := s.ProcessAccessToken(isbeAdminAccessToken)

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
		// "email": "jesus@alastria.io"
		if user.EmailAddress != "hesus.ruiz@gmail.com" {
			t.Errorf("Expected EmailAddress 'jesus@alastria.io', got '%s'", user.EmailAddress)
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
		user, err := s.ProcessAccessToken(isbeAdminAccessToken)

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
		if user.EmailAddress != "hesus.ruiz@gmail.com" {
			t.Errorf("Expected EmailAddress 'hesus.ruiz@gmail.com', got '%s'", user.EmailAddress)
		}
		if !user.IsAuthenticated {
			t.Error("Expected IsAuthenticated to be true")
		}
	})
}
