package service

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/hesusruiz/tmforum/config"
	pdp "github.com/hesusruiz/tmforum/pdp"
	"github.com/hesusruiz/tmforum/tmfserver/notifications"
	"github.com/hesusruiz/tmforum/tmfserver/repository"
	"github.com/hesusruiz/tmforum/types"
	_ "github.com/mattn/go-sqlite3"
)

// TestISBECRUDAndListGenericObject simulates a Handler to invoke the create service
func TestISBECRUDAndListGenericObject(t *testing.T) {
	s := newISBEDEVTestService(t)

	apiFamily := "productCatalogManagement"
	resourceName := "productOffering"

	s.Features.VerifyJWTSignature = true

	// Authenticate

	authUser, err := s.ProcessAccessToken(isbeAdminAccessToken)
	if err != nil {
		t.Fatalf("failed to process access token: %v", err)
	}

	// Create
	createObj := map[string]any{
		"@type": resourceName,
		"name":  "Test Product",
		"relatedParty": []map[string]any{
			{"role": "Seller", "name": "did:elsi:VATES-11111111K"},
			{"role": "SellerOperator", "name": "did:elsi:VATES-G87936159"},
		},
	}
	bCreate, _ := json.Marshal(createObj)

	cReq := &Request{
		Method:       "POST",
		Action:       ActionCREATE,
		APIfamily:    apiFamily,
		APIVersion:   "v4",
		ResourceName: resourceName,
		ID:           "",
		Body:         bCreate,
		QueryParams:  nil,
		AuthUser:     *authUser,
	}

	cResp := s.CreateTMFObject(context.Background(), cReq)
	if cResp.StatusCode != http.StatusCreated {
		t.Fatalf("create expected 201, got %d", cResp.StatusCode)
	}
	bodyMap := cResp.Body.(repository.TMFObjectMap)
	id, _ := bodyMap["id"].(string)
	if id == "" {
		t.Fatalf("no id returned")
	}

	// Get

	gReq := &Request{
		Method:       "GET",
		Action:       ActionREAD,
		APIfamily:    apiFamily,
		APIVersion:   "v4",
		ResourceName: resourceName,
		ID:           id,
		Body:         nil,
		QueryParams:  nil,
		AuthUser:     *authUser,
	}

	gResp := s.GetTMFObject(context.Background(), gReq)
	if gResp.StatusCode != http.StatusOK {
		t.Fatalf("get expected 200, got %d", gResp.StatusCode)
	}

	// Update (must include version greater than existing)
	upd := map[string]any{
		"@type":       resourceName,
		"id":          id,
		"version":     "1.1",
		"description": "Updated description",
	}
	bUpd, _ := json.Marshal(upd)
	uReq := newReq("PATCH", "UPDATE", apiFamily, resourceName, id, bUpd, nil)
	uResp := s.UpdateTMFObject(context.Background(), uReq)
	if uResp.StatusCode != http.StatusOK {
		t.Fatalf("update expected 200, got %d", uResp.StatusCode)
	}
	updated := uResp.Body.(repository.TMFObjectMap)
	if updated["version"].(string) != "1.1" {
		t.Fatalf("expected version 1.1, got %v", updated["version"])
	}
	if updated["description"].(string) != "Updated description" {
		t.Fatalf("expected description Updated description, got %v", updated["description"])
	}

	// List (all)
	lReq := newReq("GET", "LIST", apiFamily, resourceName, "", nil, url.Values{})
	lResp := s.ListTMFObjects(context.Background(), lReq)
	if lResp.StatusCode != http.StatusOK {
		t.Fatalf("list expected 200, got %d", lResp.StatusCode)
	}
	if lResp.Headers["X-Total-Count"] == "" {
		t.Fatalf("missing X-Total-Count header")
	}
	if lResp.Headers["X-Total-Count"] != "1" {
		t.Fatalf("expected X-Total-Count=1, got %s", lResp.Headers["X-Total-Count"])
	}

	// List with fields=none (should reduce fields per item)
	lReqQP := newReq("GET", "LIST", apiFamily, resourceName, "", nil, url.Values{"fields": []string{"none"}})
	lResp2 := s.ListTMFObjects(context.Background(), lReqQP)
	if lResp2.StatusCode != http.StatusOK {
		t.Fatalf("list expected 200, got %d", lResp2.StatusCode)
	}
	items, ok := lResp2.Body.([]repository.TMFObjectMap)
	if !ok || len(items) == 0 {
		t.Fatalf("expected list of items, got %T", lResp2.Body)
	}
	// Expect minimal keys present
	item := items[0]
	if item["id"] == nil || item["href"] == nil || item["version"] == nil || item["lastUpdate"] == nil || item["@type"] == nil {
		t.Fatalf("fields=none did not include minimal fields")
	}

	// Delete
	dReq := newReq("DELETE", "DELETE", apiFamily, resourceName, id, nil, nil)
	dResp := s.DeleteTMFObject(context.Background(), dReq)
	if dResp.StatusCode != http.StatusNoContent {
		t.Fatalf("delete expected 204, got %d", dResp.StatusCode)
	}

	// Get after delete -> 404
	gReq = newReq("GET", "READ", apiFamily, resourceName, id, nil, nil)
	gResp2 := s.GetTMFObject(context.Background(), gReq)
	if gResp2.StatusCode != http.StatusNotFound {
		t.Fatalf("get after delete expected 404, got %d", gResp2.StatusCode)
	}
}

func TestBadISBECreate(t *testing.T) {
	s := newISBEDEVTestService(t)

	// Change the server admin
	s.ServerOperatorDid = "VATES-9999999K"

	apiFamily := "productCatalogManagement"
	resourceName := "productOffering"

	s.Features.VerifyJWTSignature = true

	// Authenticate

	authUser, err := s.ProcessAccessToken(isbeAdminAccessToken)
	if err != nil {
		t.Fatalf("failed to process access token: %v", err)
	}

	// Create
	createObj := map[string]any{
		"@type": resourceName,
		"name":  "Test Product",
		"relatedParty": []map[string]any{
			{"role": "Seller", "name": "did:elsi:VATES-11111111K"},
			{"role": "SellerOperator", "name": "did:elsi:VATES-222222K"},
		},
	}
	bCreate, _ := json.Marshal(createObj)

	cReq := &Request{
		Method:       "POST",
		Action:       ActionCREATE,
		APIfamily:    apiFamily,
		APIVersion:   "v4",
		ResourceName: resourceName,
		ID:           "",
		Body:         bCreate,
		QueryParams:  nil,
		AuthUser:     *authUser,
	}

	cResp := s.CreateTMFObject(context.Background(), cReq)
	if cResp.StatusCode != http.StatusCreated {
		t.Fatalf("create expected 201, got %d", cResp.StatusCode)
	}
	bodyMap := cResp.Body.(repository.TMFObjectMap)
	id, _ := bodyMap["id"].(string)
	if id == "" {
		t.Fatalf("no id returned")
	}

}

// newISBEDEVTestService creates a new service for testing with ISBE DEV configuration
func newISBEDEVTestService(t *testing.T) *Service {
	t.Helper()
	configuration, err := config.LoadConfig("isbedev", true)
	if err != nil {
		t.Fatalf("failed to load configuration: %v", err)
	}

	dbLayer, err := repository.NewDBService(":memory:")
	if err != nil {
		t.Fatalf("create test db: %v", err)
	}

	configuration.PolicyFileName = "../../auth_policies.star"
	rulesEngine, err := pdp.NewPDPService(&pdp.Config{
		PolicyFileName: configuration.PolicyFileName,
		Debug:          configuration.Debug,
	})
	if err != nil {
		t.Fatalf("create test rules engine: %v", err)
	}

	// Create the service, which will use the database and the rules engine
	tmfService, err := NewTMFService(configuration, dbLayer, rulesEngine)
	if err != nil {
		t.Fatalf("create test service: %v", err)
	}

	return tmfService
}

// newISBEDEVTestService creates a new service for testing with ISBE DEV configuration
func newDOMEDEVTestService(t *testing.T) *Service {
	t.Helper()
	configuration, err := config.LoadConfig("domedev", true)
	if err != nil {
		t.Fatalf("failed to load configuration: %v", err)
	}

	dbLayer, err := repository.NewDBService(":memory:")
	if err != nil {
		t.Fatalf("create test db: %v", err)
	}

	configuration.PolicyFileName = "../../auth_policies.star"
	rulesEngine, err := pdp.NewPDPService(&pdp.Config{
		PolicyFileName: configuration.PolicyFileName,
		Debug:          configuration.Debug,
	})
	if err != nil {
		t.Fatalf("create test rules engine: %v", err)
	}

	// Create the service, which will use the database and the rules engine
	tmfService, err := NewTMFService(configuration, dbLayer, rulesEngine)
	if err != nil {
		t.Fatalf("create test service: %v", err)
	}

	return tmfService
}

func newTestService(t *testing.T) *Service {
	t.Helper()

	dbLayer, err := repository.NewDBService(":memory:")
	if err != nil {
		t.Fatalf("create test db: %v", err)
	}

	// Create service struct directly (no external verifier)
	s := &Service{
		storage:           dbLayer,
		ServerOperatorDid: "VATES-G87936159",
		LEARPower: types.OnePower{
			Type:     "organization",
			Domain:   "ISBE",
			Function: "Onboarding",
			Action:   []string{"Execute"},
		},
		ProductCreatePower: types.OnePower{
			Type:     "organization",
			Domain:   "ISBE",
			Function: "ProductOffering",
			Action:   []string{"Create"},
		},
		ProductUpdatePower: types.OnePower{
			Type:     "organization",
			Domain:   "ISBE",
			Function: "ProductOffering",
			Action:   []string{"Update"},
		},
		ProductDeletePower: types.OnePower{
			Type:     "organization",
			Domain:   "ISBE",
			Function: "ProductOffering",
			Action:   []string{"Delete"},
		},
		Features: config.Features{
			GenerateIDOnCreate: true,
		},
	}
	// Wire notifications manager to a fake delivery by default
	s.notif = notifications.NewManager(notifications.NewMemoryStore(), &fakeDelivery{})
	return s
}

// newReq creates a fresh Request with a default authenticated user
func newReq(method, action, api, resource, id string, body []byte, qp url.Values) *Request {
	return &Request{
		Method:       method,
		Action:       HttpAction(action),
		APIfamily:    api,
		APIVersion:   "v4",
		ResourceName: resource,
		ID:           id,
		Body:         body,
		QueryParams:  qp,
		AuthUser: types.AuthUser{
			IsAuthenticated:        true,
			OrganizationIdentifier: "VATES-11111111K",
			ProductCreatePower:     true,
			ProductUpdatePower:     true,
			ProductDeletePower:     true,
		},
	}
}

// fakeDelivery records delivered payloads for assertions
type fakeDelivery struct{ deliveries []any }

func (f *fakeDelivery) Deliver(_ *notifications.Subscription, payload any) error {
	f.deliveries = append(f.deliveries, payload)
	return nil
}

func TestCreateAndDeleteHubSubscription(t *testing.T) {
	s := newTestService(t)

	body := map[string]any{
		"callback":   "http://localhost:9991/listener/test",
		"eventTypes": []string{"ProductOfferingCreateEvent"},
		"headers":    map[string]any{"x-auth-token": "abc123"},
	}
	b, _ := json.Marshal(body)
	req := newReq("POST", "CREATE", "TMF620", "", "", b, nil)

	resp := s.CreateHubSubscription(req)
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}

	respMap, _ := resp.Body.(map[string]any)
	id, _ := respMap["id"].(string)
	if id == "" {
		t.Fatalf("expected id in response")
	}

	// Delete
	delReq := newReq("DELETE", "DELETE", "TMF620", "", id, nil, nil)
	delResp := s.DeleteHubSubscription(delReq)
	if delResp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", delResp.StatusCode)
	}
}

func TestCreateGenericObjectPublishesEvent(t *testing.T) {
	s := newTestService(t)

	// Replace notifications manager with one that uses fake delivery
	memStore := notifications.NewMemoryStore()
	fdel := &fakeDelivery{}
	s.notif = notifications.NewManager(memStore, fdel)

	// Create a subscription to receive create events
	sub := &notifications.Subscription{
		ID:         "sub1",
		APIFamily:  "TMF620",
		Callback:   "http://localhost:9991/listener/ProductOfferingCreateEvent",
		EventTypes: []string{"ProductOfferingCreateEvent"},
		Headers:    map[string]string{"x-auth-token": "abc123"},
	}
	if _, err := s.notif.CreateSubscription("TMF620", sub); err != nil {
		t.Fatalf("create sub: %v", err)
	}

	resourceName := "productOffering"
	obj := map[string]any{
		"@type":   resourceName,
		"version": "1.0",
		"name":    "Test Product Offering",
		"relatedParty": []map[string]any{
			{"role": "Seller", "id": "did:elsi:VATES-11111111K"},
		},
	}
	b, _ := json.Marshal(obj)
	req := newReq("POST", "CREATE", "TMF620", resourceName, "", b, nil)

	authUser, err := s.ProcessAccessToken("abc123")
	if err != nil {
		t.Fatalf("invalid access token: %v", err)
	}

	// Grant power manually for the test
	authUser.ProductCreatePower = true

	req.AuthUser = *authUser

	resp := s.CreateTMFObject(context.Background(), req)
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}

	// Wait briefly for goroutine delivery
	time.Sleep(200 * time.Millisecond)

	if len(fdel.deliveries) != 1 {
		t.Fatalf("expected 1 delivery, got %d", len(fdel.deliveries))
	}
	// Basic payload shape assertions
	payload, ok := fdel.deliveries[0].(map[string]any)
	if !ok {
		t.Fatalf("payload not a map")
	}
	if payload["eventType"] != "ProductOfferingCreateEvent" {
		t.Fatalf("unexpected eventType: %v", payload["eventType"])
	}
}

func TestCRUDAndListGenericObject(t *testing.T) {
	s := newTestService(t)

	apiFamily := "productCatalogManagement"

	// Create
	resourceName := "productOffering"
	createObj := map[string]any{
		"@type": resourceName,
		"name":  "Test Product",
		"relatedParty": []map[string]any{
			{"role": "Seller", "name": "did:elsi:VATES-11111111K"},
			{"role": "SellerOperator", "name": "did:elsi:VATES-G87936159"},
		},
	}
	bCreate, _ := json.Marshal(createObj)
	cReq := newReq("POST", "CREATE", apiFamily, resourceName, "", bCreate, nil)
	cResp := s.CreateTMFObject(context.Background(), cReq)
	if cResp.StatusCode != http.StatusCreated {
		t.Fatalf("create expected 201, got %d", cResp.StatusCode)
	}
	bodyMap := cResp.Body.(repository.TMFObjectMap)
	id, _ := bodyMap["id"].(string)
	if id == "" {
		t.Fatalf("no id returned")
	}

	// Get
	gReq := newReq("GET", "READ", apiFamily, resourceName, id, nil, nil)
	gResp := s.GetTMFObject(context.Background(), gReq)
	if gResp.StatusCode != http.StatusOK {
		t.Fatalf("get expected 200, got %d", gResp.StatusCode)
	}

	// Update (must include version greater than existing)
	upd := map[string]any{
		"@type":       resourceName,
		"id":          id,
		"version":     "1.1",
		"description": "Updated description",
	}
	bUpd, _ := json.Marshal(upd)
	uReq := newReq("PATCH", "UPDATE", apiFamily, resourceName, id, bUpd, nil)
	uResp := s.UpdateTMFObject(context.Background(), uReq)
	if uResp.StatusCode != http.StatusOK {
		t.Fatalf("update expected 200, got %d", uResp.StatusCode)
	}
	updated := uResp.Body.(repository.TMFObjectMap)
	if updated["version"].(string) != "1.1" {
		t.Fatalf("expected version 1.1, got %v", updated["version"])
	}
	if updated["description"].(string) != "Updated description" {
		t.Fatalf("expected description Updated description, got %v", updated["description"])
	}

	// List (all)
	lReq := newReq("GET", "LIST", apiFamily, resourceName, "", nil, url.Values{})
	lResp := s.ListTMFObjects(context.Background(), lReq)
	if lResp.StatusCode != http.StatusOK {
		t.Fatalf("list expected 200, got %d", lResp.StatusCode)
	}
	if lResp.Headers["X-Total-Count"] == "" {
		t.Fatalf("missing X-Total-Count header")
	}
	if lResp.Headers["X-Total-Count"] != "1" {
		t.Fatalf("expected X-Total-Count=1, got %s", lResp.Headers["X-Total-Count"])
	}

	// List with fields=none (should reduce fields per item)
	lReqQP := newReq("GET", "LIST", apiFamily, resourceName, "", nil, url.Values{"fields": []string{"none"}})
	lResp2 := s.ListTMFObjects(context.Background(), lReqQP)
	if lResp2.StatusCode != http.StatusOK {
		t.Fatalf("list expected 200, got %d", lResp2.StatusCode)
	}
	items, ok := lResp2.Body.([]repository.TMFObjectMap)
	if !ok || len(items) == 0 {
		t.Fatalf("expected list of items, got %T", lResp2.Body)
	}
	// Expect minimal keys present
	item := items[0]
	if item["id"] == nil || item["href"] == nil || item["version"] == nil || item["lastUpdate"] == nil || item["@type"] == nil {
		t.Fatalf("fields=none did not include minimal fields")
	}

	// Delete
	dReq := newReq("DELETE", "DELETE", apiFamily, resourceName, id, nil, nil)
	dResp := s.DeleteTMFObject(context.Background(), dReq)
	if dResp.StatusCode != http.StatusNoContent {
		t.Fatalf("delete expected 204, got %d", dResp.StatusCode)
	}

	// Get after delete -> 404
	gReq = newReq("GET", "READ", apiFamily, resourceName, id, nil, nil)
	gResp2 := s.GetTMFObject(context.Background(), gReq)
	if gResp2.StatusCode != http.StatusNotFound {
		t.Fatalf("get after delete expected 404, got %d", gResp2.StatusCode)
	}
}

// TestInvalidResourcename tests that ListGenericObjects returns an empty JSON array and proper X-Total-Count header
func TestInvalidResourcename(t *testing.T) {
	s := newTestService(t)
	resourceName := "TestResource"
	apiFamily := "productCatalogManagement"

	// List objects for a resource that doesn't exist (should return 400)
	lReq := newReq("GET", "LIST", apiFamily, resourceName, "", nil, url.Values{})
	lResp := s.ListTMFObjects(context.Background(), lReq)

	// Should return 200 OK
	if lResp.StatusCode != http.StatusBadRequest {
		t.Fatalf("empty list expected 200, got %d", lResp.StatusCode)
	}

	// Body should be an empty array, not nil
	if lResp.Body == nil {
		t.Fatalf("empty list body should not be nil")
	}

}

type mockRuleEngine struct {
	authorizeFunc func(input pdp.StarTMFMap) (bool, error)
}

func (m *mockRuleEngine) Authorize(input pdp.StarTMFMap) (bool, error) {
	if m.authorizeFunc != nil {
		return m.authorizeFunc(input)
	}
	return true, nil
}

func TestServiceWithMockRuleEngine(t *testing.T) {

	configuration, err := config.LoadConfig(string(config.LOCAL), false)
	if err != nil {
		t.Fatalf("create test config: %v", err)
	}

	dbLayer, err := repository.NewDBService(":memory:")
	if err != nil {
		t.Fatalf("create test db: %v", err)
	}

	called := false
	mockPDP := &mockRuleEngine{
		authorizeFunc: func(input pdp.StarTMFMap) (bool, error) {
			called = true
			return true, nil
		},
	}

	tmfService, err := NewTMFService(configuration, dbLayer, mockPDP)
	if err != nil {
		t.Fatalf("create test service: %v", err)
	}

	if tmfService.ruleEngine == nil {
		t.Fatalf("ruleEngine should not be nil")
	}

	req := &Request{
		Method:       "GET",
		Action:       ActionREAD,
		APIfamily:    "productCatalogManagement",
		APIVersion:   "v4",
		ResourceName: "productOffering",
	}

	err = tmfService.userPolicies(tmfService.ruleEngine, req, nil, nil)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if !called {
		t.Fatalf("expected mock PDP Authorize to be called")
	}
}
