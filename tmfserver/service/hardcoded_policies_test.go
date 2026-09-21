package service

import (
	"testing"

	"github.com/DOME-Marketplace/tmfgo/tmfserver/repository"
	repo "github.com/DOME-Marketplace/tmfgo/tmfserver/repository"
	"github.com/DOME-Marketplace/tmfgo/types"
)

const testSellerAccessToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICItckxwSkhNVkhCSUQ0Q2FRX0dsTjhFTEprQ0tYMUJWUzhMTzd6enU1cTFVIn0.eyJleHAiOjE3Nzc3MzI1OTIsImlhdCI6MTc3NTE0MDU5MiwiYXV0aF90aW1lIjoxNzc1MTQwNTkyLCJqdGkiOiJvbnJ0YWM6ZDU3NDk4ODItZWJhOC0wNmJhLTk0MjktNjk5NDBlNDE1ZGU4IiwiaXNzIjoiaHR0cHM6Ly9pZHAuZGV2LmNsb3VkLXcuZW52cy5yZWRpc2JlLmNvbS9hdXRoL3JlYWxtcy9kZXYtaXNiZSIsImF1ZCI6WyJpc2JlLXBvcnRhbC1kZXYiLCJhY2NvdW50Il0sInN1YiI6ImIyMzQ5ZTMxLWEyOTktNGU5Yi04ZGQxLWQxMWExZDFmNzBjMSIsInR5cCI6IkJlYXJlciIsImF6cCI6Imh0dHBzOi8vY2F0YWxvZy5pc2Jlb25ib2FyZC5jb20iLCJzaWQiOiIyZWFhMjMyZC03ZTI5LTA1ZmUtM2YzZi04MjYwYTYzZWVjZTAiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbImh0dHBzOi8vY2F0YWxvZy5kZXYuY2xvdWQtdy5lbnZzLnJlZGlzYmUuY29tIl0sInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJkZWZhdWx0LXJvbGVzLWRldi1pc2JlIiwib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIG9yZ2FuaXphdGlvbiBlbWFpbCBwcm9maWxlIiwidXNlcl9pZGVudGlmaWVyIjoiMTIzNDU2NzhBIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsIm9yZ2FuaXphdGlvbiI6IkFMQVNUUklBIiwibmFtZSI6IkpvaG4gRG9lIiwib3JnYW5pemF0aW9uX2lkZW50aWZpZXIiOiJWQVRFUy1HODc5MzYxNTkiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJqZXN1c0BhbGFzdHJpYS5pbyIsInBvd2VyIjpbeyJhY3Rpb24iOlsiKiJdLCJkb21haW4iOiJJU0JFIiwiZnVuY3Rpb24iOiJNYW5hZ2VtZW50IiwidHlwZSI6Im9yZ2FuaXphdGlvbiJ9LHsiYWN0aW9uIjpbIioiXSwiZG9tYWluIjoiSVNCRSIsImZ1bmN0aW9uIjoiSGVscGRlc2siLCJ0eXBlIjoib3JnYW5pemF0aW9uIn0seyJhY3Rpb24iOlsiKiJdLCJkb21haW4iOiJJU0JFIiwiZnVuY3Rpb24iOiJGYXVjZXQiLCJ0eXBlIjoib3JnYW5pemF0aW9uIn0seyJhY3Rpb24iOlsiKiJdLCJkb21haW4iOiJJU0JFIiwiZnVuY3Rpb24iOiJXaXphcmQiLCJ0eXBlIjoib3JnYW5pemF0aW9uIn0seyJhY3Rpb24iOlsiKiJdLCJkb21haW4iOiJJU0JFIiwiZnVuY3Rpb24iOiJOb3Rhcml6YXRpb24iLCJ0eXBlIjoib3JnYW5pemF0aW9uIn0seyJhY3Rpb24iOlsiKiJdLCJkb21haW4iOiJJU0JFIiwiZnVuY3Rpb24iOiJOb3RpZmljYXRpb25zIiwidHlwZSI6Im9yZ2FuaXphdGlvbiJ9LHsiYWN0aW9uIjpbIioiXSwiZG9tYWluIjoiSVNCRSIsImZ1bmN0aW9uIjoiSWRlbnRpdHkiLCJ0eXBlIjoib3JnYW5pemF0aW9uIn0seyJhY3Rpb24iOlsiKiJdLCJkb21haW4iOiJJU0JFIiwiZnVuY3Rpb24iOiJFbnJvbGxtZW50IiwidHlwZSI6Im9yZ2FuaXphdGlvbiJ9LHsiYWN0aW9uIjpbIkV4ZWN1dGUiXSwiZG9tYWluIjoiSVNCRSIsImZ1bmN0aW9uIjoiT25ib2FyZGluZyIsInR5cGUiOiJvcmdhbml6YXRpb24ifV0sImdpdmVuX25hbWUiOiJKb2huIiwidXNlciI6IkpvaG4gRG9lIiwiZmFtaWx5X25hbWUiOiJEb2UiLCJlbWFpbCI6Imhlc3VzLnJ1aXpAZ21haWwuY29tIn0.iUu88PYBYydC_YuYNfi9gma3jg3q1IPzys9Wfo_KvFoDatLhTshQ1b8gVOPB3Z4WFi_7TUcTmejtgYwYr79qjo1A2upSzL_IbGt20w4zUtuxJ4WA0anaD5cuyKBUieUNiEwR-_UKZyFj1dodIAtEIAU_D4qlxOgtZ2WKcxfT8RZajyUzHYCkZSiTLMlxkpXNciqaFOPD5Tjnu9N_G27WtknwBEIRRa3okM8K-mSn6zXpIVxE8kgi7Cvd0OJnOI6Jo4eWLtrYqw750CsjU8n6Y5L0Kl5M8bNOiKwMzVvChjPWcGRowxpJozLAZzhqQcmlhS8Sn79oSXWb67iC339Ogw"
const testServerOperatorAccessToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOi"

const (
	testServerOperator = "VATES-Foundation"
	testSeller         = "VATES-G87936159"
	testBuyer          = "did:elsi:VATES-B22222222"
	testOtherOrg       = "did:elsi:VATES-C33333333"
	testExtNodeOp      = "did:elsi:VATES-X99999999"
)

const (
	testSeller1      = "did:elsi:VATES-Seller00001"
	testSeller2      = "did:elsi:VATES-Seller00002"
	testMarketplace1 = "did:elsi:MP-MARKETPLACE-00001"
	testMarketplace2 = "did:elsi:MP-MARKETPLACE-00002"
)

func init() {
	types.ParseActionDefinitions()
}

func newTestServiceForPolicies() *Service {
	return &Service{
		ServerOperatorDid: testServerOperator,
	}
}

// -----------------------------------------------------------------------------
// Helper constructors for TMF test objects and AuthUsers
// -----------------------------------------------------------------------------

func makeUser(orgDid string, isAuthenticated, isLear, createPwr, updatePwr, deletePwr bool) types.AuthUser {
	return types.AuthUser{
		IsAuthenticated:        isAuthenticated,
		OrganizationIdentifier: orgDid,
		IsLEAR:                 isLear,
		ProductCreatePower:     createPwr,
		ProductUpdatePower:     updatePwr,
		ProductDeletePower:     deletePwr,
	}
}

func makePublicOffering(seller, sellerOp string) repo.TMFObjectMap {
	obj := repo.TMFObjectMap{
		"@type": "productOffering",
		"id":    "offering-100",
	}
	if seller != "" || sellerOp != "" {
		_ = obj.SetSellerInfo(sellerOp, seller, "v4")
	}
	return obj
}

func makePrivateOffering(seller, sellerOp, buyer, buyerOp string) repo.TMFObjectMap {
	return repo.TMFObjectMap{
		"@type": "productOffering",
		"id":    "offering-200",
		"relatedParty": []any{
			map[string]any{"role": "Seller", "name": seller},
			map[string]any{"role": "SellerOperator", "name": sellerOp},
			map[string]any{"role": "Buyer", "name": buyer},
			map[string]any{"role": "BuyerOperator", "name": buyerOp},
		},
	}
}

func makeSpecialObject(objType, orgOrIssuingId string) repo.TMFObjectMap {
	obj := repo.TMFObjectMap{
		"@type": objType,
		"id":    objType + "-1",
	}
	switch objType {
	case "Organization", "organization":
		obj["organizationIdentification"] = []any{
			map[string]any{
				"identificationId": orgOrIssuingId,
			},
		}
	case "Individual", "individual":
		obj["individualIdentification"] = []any{
			map[string]any{
				"identificationType": "learcredentialemployee",
				"issuingAuthority":   orgOrIssuingId,
			},
		}
	}
	return obj
}

// -----------------------------------------------------------------------------
// Test 1: Attribute Integrity Verification
// -----------------------------------------------------------------------------

func TestHardcodedPolicies_AttributeIntegrity(t *testing.T) {
	svc := newTestServiceForPolicies()
	req := &Request{
		Action:   ActionREAD,
		AuthUser: makeUser(testServerOperator, true, true, true, true, true),
	}

	// Expect error when partially setting Seller info
	badSellerObj := repo.TMFObjectMap{
		"@type": "productOffering",
		"relatedParty": []any{
			map[string]any{"role": "Seller", "name": testSeller},
		},
	}
	if _, err := svc.hardcodedPolicies(req, badSellerObj); err == nil {
		t.Errorf("expected attribute integrity error for partially set Seller, got nil")
	}

	// Expect error when partially setting Buyer info
	badBuyerObj := repo.TMFObjectMap{
		"@type": "productOffering",
		"relatedParty": []any{
			map[string]any{"role": "Seller", "name": testSeller},
			map[string]any{"role": "SellerOperator", "name": testServerOperator},
			map[string]any{"role": "Buyer", "name": testBuyer},
		},
	}
	if _, err := svc.hardcodedPolicies(req, badBuyerObj); err == nil {
		t.Errorf("expected attribute integrity error for partially set Buyer, got nil")
	}

	// Expect success when no related party is present
	somePartyObj := repo.TMFObjectMap{
		"@type": "productOffering",
	}
	if _, err := svc.hardcodedPolicies(req, somePartyObj); err != nil {
		t.Errorf("expected no error when no full seller info or buyer info are present, got %v", err)
	}

	// Expect success when full seller info set and buyer related party
	fullSellerObj := repo.TMFObjectMap{
		"@type": "productOffering",
		"relatedParty": []any{
			map[string]any{"role": "Seller", "name": testSeller},
			map[string]any{"role": "SellerOperator", "name": testServerOperator},
		},
	}
	if _, err := svc.hardcodedPolicies(req, fullSellerObj); err != nil {
		t.Errorf("expected no error when full seller info set and buyer related party is not present, got %v", err)
	}

	// Expect success when full seller info set and buyer related party set
	fullSellerAndBuyerObj := repo.TMFObjectMap{
		"@type": "productOffering",
		"relatedParty": []any{
			map[string]any{"role": "Seller", "name": testSeller},
			map[string]any{"role": "SellerOperator", "name": testServerOperator},
			map[string]any{"role": "Buyer", "name": testBuyer},
			map[string]any{"role": "BuyerOperator", "name": testServerOperator},
		},
	}
	if _, err := svc.hardcodedPolicies(req, fullSellerAndBuyerObj); err != nil {
		t.Errorf("expected no error when full seller info set and buyer related party set, got %v", err)
	}

	// Expect error when buyer info set but no seller info set
	buyerOnlyObj := repo.TMFObjectMap{
		"@type": "productOffering",
		"relatedParty": []any{
			map[string]any{"role": "Buyer", "name": testBuyer},
			map[string]any{"role": "BuyerOperator", "name": testServerOperator},
		},
	}
	if _, err := svc.hardcodedPolicies(req, buyerOnlyObj); err == nil {
		t.Errorf("expected error when buyer info set but no seller info set, got nil")
	}
}

// -----------------------------------------------------------------------------
// Test 2: CREATE Operations Compliance
// -----------------------------------------------------------------------------

func TestHardcodedPolicies_Create(t *testing.T) {
	svc := newTestServiceForPolicies()

	// Creating a public offering managed by the Server Operator
	pubObj := makePublicOffering(testSeller, testServerOperator)

	// An unauthenticated user
	unauthUser := makeUser("", false, false, false, false, false)

	// A server operator with LEAR
	serverOperatorLear := makeUser(testServerOperator, true, true, true, true, true)

	// A seller without create power
	sellerNoCreatePwr := makeUser(testSeller, true, false, false, true, true)

	// A seller with create power
	sellerWithCreatePwr := makeUser(testSeller, true, false, true, true, true)

	// Tests with the public offering
	{
		// Unauthenticated users can not create objects
		unauthReq := &Request{Action: ActionCREATE, AuthUser: unauthUser}
		if _, err := svc.hardcodedPolicies(unauthReq, pubObj); err == nil {
			t.Errorf("unauthenticated CREATE should fail")
		}

		// A server operator can create any object, in particular a public offering
		soLearReq := &Request{Action: ActionCREATE, AuthUser: serverOperatorLear}
		if _, err := svc.hardcodedPolicies(soLearReq, pubObj); err != nil {
			t.Errorf("ServerOperator LEAR CREATE failed: %v", err)
		}

		// A seller without create power can not create objects
		noPwrSellerReq := &Request{Action: ActionCREATE, AuthUser: sellerNoCreatePwr}
		if _, err := svc.hardcodedPolicies(noPwrSellerReq, pubObj); err == nil {
			t.Errorf("CREATE without ProductCreatePower should fail")
		}

		// A seller with create power can create objects
		sellerReq := &Request{Action: ActionCREATE, AuthUser: sellerWithCreatePwr}
		if _, err := svc.hardcodedPolicies(sellerReq, pubObj); err != nil {
			t.Errorf("Seller CREATE failed: %v", err)
		}
	}

	// For a seller, specifying the seller info is optional, as it will be filled automatically
	emptySellerObj := makePublicOffering("", "")
	sellerReq := &Request{Action: ActionCREATE, AuthUser: sellerWithCreatePwr}
	if _, err := svc.hardcodedPolicies(sellerReq, emptySellerObj); err != nil {
		t.Errorf("Seller CREATE with omitted Seller info failed: %v", err)
	}
	s, so, _ := emptySellerObj.GetSellerInfo("")
	if !repository.SameOrganizations(s, testSeller) || !repository.SameOrganizations(so, testServerOperator) {
		t.Errorf("auto-assign Seller info failed: expected seller=%s so=%s, got seller=%s so=%s", testSeller, testServerOperator, s, so)
	}

	// A seller can not create a public offering managed by a different marketplace.
	offerMkt1 := makePublicOffering(testSeller, testMarketplace1)
	if _, err := svc.hardcodedPolicies(sellerReq, offerMkt1); err == nil {
		t.Errorf("public object CREATE with SellerOperator != ServerOperator should fail")
	}

	// Not even if the seller is the external marketplace
	extOfferMkt1 := makePublicOffering(testMarketplace1, testMarketplace1)
	if _, err := svc.hardcodedPolicies(sellerReq, extOfferMkt1); err == nil {
		t.Errorf("public object CREATE with seller being the external marketplace should fail")
	}

	// Tests with Special Objects
	{
		// 2f. Special objects CREATE restrictions
		catObj := makeSpecialObject("category", "")
		orgObj := makeSpecialObject("organization", testSeller)
		indObj := makeSpecialObject("individual", testSeller)

		if _, err := svc.hardcodedPolicies(sellerReq, catObj); err == nil {
			t.Errorf("non-ServerOperator CREATE Category should fail")
		}
		if _, err := svc.hardcodedPolicies(sellerReq, orgObj); err == nil {
			t.Errorf("non-ServerOperator CREATE Organization should fail")
		}

		soLearReq := &Request{Action: ActionCREATE, AuthUser: serverOperatorLear}

		// ServerOperator can create Category and Organization
		if _, err := svc.hardcodedPolicies(soLearReq, catObj); err != nil {
			t.Errorf("ServerOperator CREATE Category failed: %v", err)
		}
		if _, err := svc.hardcodedPolicies(soLearReq, orgObj); err != nil {
			t.Errorf("ServerOperator CREATE Organization failed: %v", err)
		}

		// Individual CREATE requires caller to match mandator (issuingAuthority)
		if _, err := svc.hardcodedPolicies(sellerReq, indObj); err != nil {
			t.Errorf("matching mandator CREATE Individual failed: %v", err)
		}
		otherOrgReq := &Request{Action: ActionCREATE, AuthUser: makeUser(testOtherOrg, true, false, true, true, true)}
		if _, err := svc.hardcodedPolicies(otherOrgReq, indObj); err == nil {
			t.Errorf("non-mandator CREATE Individual should fail")
		}

	}

}

// -----------------------------------------------------------------------------
// Test 3: READ / LIST Operations Compliance
// -----------------------------------------------------------------------------

func TestHardcodedPolicies_ReadList(t *testing.T) {
	svc := newTestServiceForPolicies()

	unauthReq := &Request{Action: ActionREAD, Method: "GET", AuthUser: makeUser("", false, false, false, false, false)}
	authBuyerReq := &Request{Action: ActionREAD, Method: "GET", AuthUser: makeUser(testBuyer, true, false, false, false, false)}
	authSellerReq := &Request{Action: ActionREAD, Method: "GET", AuthUser: makeUser(testSeller, true, false, false, false, false)}
	authForeignReq := &Request{Action: ActionREAD, Method: "GET", AuthUser: makeUser(testOtherOrg, true, false, false, false, false)}
	serverOperatorRequest := &Request{Action: ActionREAD, Method: "GET", AuthUser: makeUser(testServerOperator, true, true, true, true, true)}

	pubObj := makePublicOffering(testSeller, testServerOperator)
	privObj := makePrivateOffering(testSeller, testServerOperator, testBuyer, testServerOperator)
	catObj := makeSpecialObject("category", "")
	orgObj := makeSpecialObject("organization", testOtherOrg)
	indObj := makeSpecialObject("individual", testSeller)

	// Public offerings not Launched can not be read by unauthenticated users
	if _, err := svc.hardcodedPolicies(unauthReq, pubObj); err == nil {
		t.Errorf("unauthenticated READ public offering failed: %v", err)
	}

	// Once an offering is Launched, it can be read by anyone
	pubObj.SetLifecycleStatus("Launched")
	if _, err := svc.hardcodedPolicies(unauthReq, pubObj); err != nil {
		t.Errorf("unauthenticated READ public offering failed: %v", err)
	}

	// Categories and Organizations are always readable by anyone
	if _, err := svc.hardcodedPolicies(unauthReq, catObj); err != nil {
		t.Errorf("unauthenticated READ Category failed: %v", err)
	}
	if _, err := svc.hardcodedPolicies(unauthReq, orgObj); err != nil {
		t.Errorf("unauthenticated READ Organization failed: %v", err)
	}

	// Authenticated user can read any Category and Organization
	if _, err := svc.hardcodedPolicies(authForeignReq, catObj); err != nil {
		t.Errorf("authenticated READ Category failed: %v", err)
	}
	if _, err := svc.hardcodedPolicies(authForeignReq, orgObj); err != nil {
		t.Errorf("authenticated READ Organization failed: %v", err)
	}

	// Private objects can not be read by unauthenticated users
	if _, err := svc.hardcodedPolicies(unauthReq, privObj); err == nil {
		t.Errorf("unauthenticated READ private object should fail")
	}

	// Private object READ by involved party (Seller or Buyer) must succeed
	// Also the server operator can read any private object
	if _, err := svc.hardcodedPolicies(authBuyerReq, privObj); err != nil {
		t.Errorf("involved party READ private object failed: %v", err)
	}
	if _, err := svc.hardcodedPolicies(authBuyerReq, privObj); err != nil {
		t.Errorf("involved party READ private object failed: %v", err)
	}
	if _, err := svc.hardcodedPolicies(serverOperatorRequest, privObj); err != nil {
		t.Errorf("ServerOperator READ private object failed: %v", err)
	}

	// Private object READ by uninvolved third party must fail
	if _, err := svc.hardcodedPolicies(authForeignReq, privObj); err == nil {
		t.Errorf("uninvolved party READ private object should fail")
	}

	// Individual objects can only be read by the employer or by the Seller operator
	if _, err := svc.hardcodedPolicies(authSellerReq, indObj); err != nil {
		t.Errorf("mandator READ Individual failed: %v", err)
	}
	if _, err := svc.hardcodedPolicies(serverOperatorRequest, indObj); err != nil {
		t.Errorf("ServerOperator READ Individual failed: %v", err)
	}
	if _, err := svc.hardcodedPolicies(authForeignReq, indObj); err == nil {
		t.Errorf("uninvolved party READ Individual should fail")
	}

}

// -----------------------------------------------------------------------------
// Test 4: UPDATE Operations Compliance
// -----------------------------------------------------------------------------

func TestHardcodedPolicies_Update(t *testing.T) {
	svc := newTestServiceForPolicies()

	unauthReq := &Request{Action: ActionUPDATE, AuthUser: makeUser("", false, false, false, false, false)}
	sellerReq := &Request{Action: ActionUPDATE, AuthUser: makeUser(testSeller, true, false, false, true, true)}
	sellerNoPwrReq := &Request{Action: ActionUPDATE, AuthUser: makeUser(testSeller, true, false, true, false, true)}

	pubObj := makePublicOffering(testSeller, testServerOperator)
	catObj := makeSpecialObject("category", "")
	orgObj := makeSpecialObject("organization", testSeller)
	indObj := makeSpecialObject("individual", testSeller)

	// 4a. Unauthenticated UPDATE must fail
	if _, err := svc.hardcodedPolicies(unauthReq, pubObj); err == nil {
		t.Errorf("unauthenticated UPDATE should fail")
	}

	// 4b. UPDATE without ProductUpdatePower must fail
	if _, err := svc.hardcodedPolicies(sellerNoPwrReq, pubObj); err == nil {
		t.Errorf("UPDATE without ProductUpdatePower should fail")
	}

	// 4c. Seller UPDATE public object with power must succeed
	if _, err := svc.hardcodedPolicies(sellerReq, pubObj); err != nil {
		t.Errorf("Seller UPDATE public object failed: %v", err)
	}

	// 4d. Special objects UPDATE restrictions
	if _, err := svc.hardcodedPolicies(sellerReq, catObj); err == nil {
		t.Errorf("non-ServerOperator UPDATE Category should fail")
	}
	if _, err := svc.hardcodedPolicies(sellerReq, orgObj); err == nil {
		t.Errorf("non-ServerOperator UPDATE Organization should fail")
	}
	if _, err := svc.hardcodedPolicies(sellerReq, indObj); err != nil {
		t.Errorf("mandator UPDATE Individual with power failed: %v", err)
	}
}

// -----------------------------------------------------------------------------
// Test 5: DELETE Operations Compliance
// -----------------------------------------------------------------------------

func TestHardcodedPolicies_Delete(t *testing.T) {
	svc := newTestServiceForPolicies()

	soLearReq := &Request{Action: ActionDELETE, AuthUser: makeUser(testServerOperator, true, true, true, true, true)}
	sellerReq := &Request{Action: ActionDELETE, AuthUser: makeUser(testSeller, true, false, true, true, true)}
	buyerReq := &Request{Action: ActionDELETE, AuthUser: makeUser(testBuyer, true, false, true, true, true)}
	buyerOpReq := &Request{Action: ActionDELETE, AuthUser: makeUser(testServerOperator, true, false, true, true, true)}

	pubObj := makePublicOffering(testSeller, testServerOperator)
	privObj := makePrivateOffering(testSeller, testServerOperator, testBuyer, testServerOperator)
	catObj := makeSpecialObject("category", "")
	orgObj := makeSpecialObject("organization", testSeller)

	// 5a. Public object DELETE by Seller must succeed
	if _, err := svc.hardcodedPolicies(sellerReq, pubObj); err != nil {
		t.Errorf("Seller DELETE public object failed: %v", err)
	}

	// 5b. Private object DELETE by Seller must succeed
	if _, err := svc.hardcodedPolicies(sellerReq, privObj); err != nil {
		t.Errorf("Seller DELETE private object failed: %v", err)
	}

	// 5c. Private object DELETE by Buyer directly MUST FAIL (per hardcodedPolicies.md line 281)
	if _, err := svc.hardcodedPolicies(buyerReq, privObj); err == nil {
		t.Errorf("direct Buyer DELETE on private object MUST fail per hardcodedPolicies.md")
	}

	// 5d. Private object DELETE by BuyerOperator must succeed
	if _, err := svc.hardcodedPolicies(buyerOpReq, privObj); err != nil {
		t.Errorf("BuyerOperator DELETE private object failed: %v", err)
	}

	// 5e. Special objects DELETE restrictions
	if _, err := svc.hardcodedPolicies(sellerReq, catObj); err == nil {
		t.Errorf("non-ServerOperator DELETE Category should fail")
	}
	if _, err := svc.hardcodedPolicies(sellerReq, orgObj); err == nil {
		t.Errorf("non-ServerOperator DELETE Organization should fail")
	}

	// ServerOperator can DELETE Category and Organization
	if _, err := svc.hardcodedPolicies(soLearReq, catObj); err != nil {
		t.Errorf("ServerOperator DELETE Category failed: %v", err)
	}
	if _, err := svc.hardcodedPolicies(soLearReq, orgObj); err != nil {
		t.Errorf("ServerOperator DELETE Organization failed: %v", err)
	}
}
