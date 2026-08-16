package service

import (
	"testing"

	repo "github.com/hesusruiz/tmforum/tmfserver/repository"
	"github.com/hesusruiz/tmforum/types"
)

const (
	testServerOp  = "did:elsi:VATES-G87936159"
	testSeller    = "did:elsi:VATES-A11111111"
	testBuyer     = "did:elsi:VATES-B22222222"
	testOtherOrg  = "did:elsi:VATES-C33333333"
	testExtNodeOp = "did:elsi:VATES-X99999999"
)

func init() {
	types.ParseActionDefinitions()
}

func newTestServiceForPolicies() *Service {
	return &Service{
		ServerOperatorDid: testServerOp,
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
		AuthUser: makeUser(testServerOp, true, true, true, true, true),
	}

	// Partially set Seller info
	badSellerObj := repo.TMFObjectMap{
		"@type": "productOffering",
		"relatedParty": []any{
			map[string]any{"role": "Seller", "name": testSeller},
		},
	}
	if _, err := svc.hardcodedPolicies(req, badSellerObj); err == nil {
		t.Errorf("expected attribute integrity error for partially set Seller, got nil")
	}

	// Partially set Buyer info
	badBuyerObj := repo.TMFObjectMap{
		"@type": "productOffering",
		"relatedParty": []any{
			map[string]any{"role": "Seller", "name": testSeller},
			map[string]any{"role": "SellerOperator", "name": testServerOp},
			map[string]any{"role": "Buyer", "name": testBuyer},
		},
	}
	if _, err := svc.hardcodedPolicies(req, badBuyerObj); err == nil {
		t.Errorf("expected attribute integrity error for partially set Buyer, got nil")
	}
}

// -----------------------------------------------------------------------------
// Test 2: CREATE Operations Compliance
// -----------------------------------------------------------------------------

func TestHardcodedPolicies_Create(t *testing.T) {
	svc := newTestServiceForPolicies()

	// 2a. Unauthenticated CREATE must fail
	unauthReq := &Request{Action: ActionCREATE, AuthUser: makeUser("", false, false, false, false, false)}
	pubObj := makePublicOffering(testSeller, testServerOp)
	if _, err := svc.hardcodedPolicies(unauthReq, pubObj); err == nil {
		t.Errorf("unauthenticated CREATE should fail")
	}

	// 2b. ServerOperator CREATE (with power & LEAR)
	soLearReq := &Request{Action: ActionCREATE, AuthUser: makeUser(testServerOp, true, true, true, true, true)}
	if _, err := svc.hardcodedPolicies(soLearReq, pubObj); err != nil {
		t.Errorf("ServerOperator LEAR CREATE failed: %v", err)
	}

	// 2c. Non-ServerOperator CREATE without ProductCreatePower must fail
	noPwrSellerReq := &Request{Action: ActionCREATE, AuthUser: makeUser(testSeller, true, false, false, true, true)}
	if _, err := svc.hardcodedPolicies(noPwrSellerReq, pubObj); err == nil {
		t.Errorf("CREATE without ProductCreatePower should fail")
	}

	// 2d. Seller CREATE public object auto-populates Seller/SellerOperator when missing
	emptySellerObj := makePublicOffering("", "")
	sellerReq := &Request{Action: ActionCREATE, AuthUser: makeUser(testSeller, true, false, true, true, true)}
	if _, err := svc.hardcodedPolicies(sellerReq, emptySellerObj); err != nil {
		t.Errorf("Seller CREATE with omitted Seller info failed: %v", err)
	}
	s, so, _ := emptySellerObj.GetSellerInfo("")
	if s != testSeller || so != testServerOp {
		t.Errorf("auto-assign Seller info failed: expected seller=%s so=%s, got seller=%s so=%s", testSeller, testServerOp, s, so)
	}

	// 2e. Public object CREATE with SellerOperator != ServerOperator must fail
	extSoObj := makePublicOffering(testSeller, testExtNodeOp)
	if _, err := svc.hardcodedPolicies(sellerReq, extSoObj); err == nil {
		t.Errorf("public object CREATE with SellerOperator != ServerOperator should fail")
	}

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

// -----------------------------------------------------------------------------
// Test 3: READ / LIST Operations Compliance
// -----------------------------------------------------------------------------

func TestHardcodedPolicies_ReadList(t *testing.T) {
	svc := newTestServiceForPolicies()

	unauthReq := &Request{Action: ActionREAD, Method: "GET", AuthUser: makeUser("", false, false, false, false, false)}
	authBuyerReq := &Request{Action: ActionREAD, Method: "GET", AuthUser: makeUser(testBuyer, true, false, false, false, false)}
	authOtherReq := &Request{Action: ActionREAD, Method: "GET", AuthUser: makeUser(testOtherOrg, true, false, false, false, false)}

	pubObj := makePublicOffering(testSeller, testServerOp)
	privObj := makePrivateOffering(testSeller, testServerOp, testBuyer, testServerOp)
	catObj := makeSpecialObject("category", "")
	orgObj := makeSpecialObject("organization", testOtherOrg)
	indObj := makeSpecialObject("individual", testSeller)

	// 3a. Unauthenticated READ on public objects (offering, category, organization) must succeed
	if _, err := svc.hardcodedPolicies(unauthReq, pubObj); err != nil {
		t.Errorf("unauthenticated READ public offering failed: %v", err)
	}
	if _, err := svc.hardcodedPolicies(unauthReq, catObj); err != nil {
		t.Errorf("unauthenticated READ Category failed: %v", err)
	}
	if _, err := svc.hardcodedPolicies(unauthReq, orgObj); err != nil {
		t.Errorf("unauthenticated READ Organization failed: %v", err)
	}

	// 3b. Authenticated READ on Category and Organization by any user must succeed
	if _, err := svc.hardcodedPolicies(authOtherReq, catObj); err != nil {
		t.Errorf("authenticated READ Category failed: %v", err)
	}
	if _, err := svc.hardcodedPolicies(authOtherReq, orgObj); err != nil {
		t.Errorf("authenticated READ Organization failed: %v", err)
	}

	// 3c. Unauthenticated READ on private object must fail
	if _, err := svc.hardcodedPolicies(unauthReq, privObj); err == nil {
		t.Errorf("unauthenticated READ private object should fail")
	}

	// 3d. Private object READ by involved party (Buyer) must succeed
	if _, err := svc.hardcodedPolicies(authBuyerReq, privObj); err != nil {
		t.Errorf("involved party READ private object failed: %v", err)
	}

	// 3e. Private object READ by uninvolved third party must fail
	if _, err := svc.hardcodedPolicies(authOtherReq, privObj); err == nil {
		t.Errorf("uninvolved party READ private object should fail")
	}

	// 3f. Individual READ by mandator vs uninvolved party
	authSellerReq := &Request{Action: ActionREAD, Method: "GET", AuthUser: makeUser(testSeller, true, false, false, false, false)}
	if _, err := svc.hardcodedPolicies(authSellerReq, indObj); err != nil {
		t.Errorf("mandator READ Individual failed: %v", err)
	}
	if _, err := svc.hardcodedPolicies(authOtherReq, indObj); err == nil {
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

	pubObj := makePublicOffering(testSeller, testServerOp)
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

	soLearReq := &Request{Action: ActionDELETE, AuthUser: makeUser(testServerOp, true, true, true, true, true)}
	sellerReq := &Request{Action: ActionDELETE, AuthUser: makeUser(testSeller, true, false, true, true, true)}
	buyerReq := &Request{Action: ActionDELETE, AuthUser: makeUser(testBuyer, true, false, true, true, true)}
	buyerOpReq := &Request{Action: ActionDELETE, AuthUser: makeUser(testServerOp, true, false, true, true, true)}

	pubObj := makePublicOffering(testSeller, testServerOp)
	privObj := makePrivateOffering(testSeller, testServerOp, testBuyer, testServerOp)
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
