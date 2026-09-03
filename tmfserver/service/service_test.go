package service

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/hesusruiz/tmforum/config"
	pdp "github.com/hesusruiz/tmforum/pdp"
	"github.com/hesusruiz/tmforum/tmfserver/notifications"
	"github.com/hesusruiz/tmforum/tmfserver/repository"
	"github.com/hesusruiz/tmforum/types"
	_ "github.com/mattn/go-sqlite3"
)

// TestLocalCRUDAndListGenericObject exercises the CRUD functions in a local database.
func TestLocalCRUDAndListGenericObject(t *testing.T) {
	s := newLocalTestService(t)

	// We will manage a productOffering object
	apiFamily := "productCatalogManagement"
	resourceName := "productOffering"

	// To authenticate as a seller
	sellerUser, err := s.ProcessAccessToken(testSellerAccessToken)
	if err != nil {
		t.Fatalf("failed to process access token: %v", err)
	}

	// To authenticate as another seller
	foreignUser := &types.AuthUser{
		OrganizationIdentifier: "VATES-11111111K",
		IsAuthenticated:        true,
		IsLEAR:                 true,
		ProductCreatePower:     true,
		ProductUpdatePower:     true,
		ProductDeletePower:     true,
	}

	// Define the object to create
	createObj := map[string]any{
		"@type": resourceName,
		"name":  "Test Product",
		"relatedParty": []map[string]any{
			{"role": "Seller", "name": testSeller},
			{"role": "SellerOperator", "name": testServerOperator},
		},
	}
	requestBody, _ := json.Marshal(createObj)

	// The request for the create service
	request := &Request{
		Method:       "POST",
		Action:       ActionCREATE,
		APIfamily:    apiFamily,
		APIVersion:   "v4",
		ResourceName: resourceName,
		ID:           "",
		Body:         requestBody,
		QueryParams:  nil,
		AuthUser:     *sellerUser,
	}

	// Create the object
	response := s.CreateTMFObject(context.Background(), request)
	if response.StatusCode != http.StatusCreated {
		t.Fatalf("create expected 201, got %d", response.StatusCode)
	}
	bodyMap := response.Body.(repository.TMFObjectMap)
	id, _ := bodyMap["id"].(string)
	if id == "" {
		t.Fatalf("no id returned")
	}

	// Retrieve the object, authenticating as the same seller

	request = &Request{
		Method:       "GET",
		Action:       ActionREAD,
		APIfamily:    apiFamily,
		APIVersion:   "v4",
		ResourceName: resourceName,
		ID:           id,
		Body:         nil,
		QueryParams:  nil,
		AuthUser:     *sellerUser,
	}

	response = s.GetTMFObject(context.Background(), request)
	if response.StatusCode != http.StatusOK {
		t.Fatalf("get expected 200, got %d", response.StatusCode)
	}

	// Update the object incrementing the version
	upd := map[string]any{
		"@type":       resourceName,
		"id":          id,
		"version":     "1.1",
		"description": "Updated description",
	}
	bUpd, _ := json.Marshal(upd)

	request = &Request{
		Method:       "PATCH",
		Action:       ActionUPDATE,
		APIfamily:    apiFamily,
		APIVersion:   "v4",
		ResourceName: resourceName,
		ID:           id,
		Body:         bUpd,
		QueryParams:  nil,
		AuthUser:     *sellerUser,
	}

	response = s.UpdateTMFObject(context.Background(), request)
	if response.StatusCode != http.StatusOK {
		t.Fatalf("update expected 200, got %d", response.StatusCode)
	}
	updated := response.Body.(repository.TMFObjectMap)
	if updated["version"].(string) != "1.1" {
		t.Fatalf("expected version 1.1, got %v", updated["version"])
	}
	if updated["description"].(string) != "Updated description" {
		t.Fatalf("expected description Updated description, got %v", updated["description"])
	}

	// List (all). First with a user who is not the seller.
	// The offereing is not yet launched, so it should not be visible.

	request = &Request{
		Method:       "GET",
		Action:       ActionLIST,
		APIfamily:    apiFamily,
		APIVersion:   "v4",
		ResourceName: resourceName,
		ID:           "",
		Body:         nil,
		QueryParams:  url.Values{},
		AuthUser:     *foreignUser,
	}

	response = s.ListTMFObjects(context.Background(), request)
	if response.StatusCode != http.StatusOK {
		t.Fatalf("list expected 200, got %d", response.StatusCode)
	}

	responseObjects, ok := response.Body.([]repository.TMFObjectMap)
	if !ok {
		t.Fatalf("expected list of objects, got %T", response.Body)
	}
	if len(responseObjects) != 0 {
		t.Fatalf("expected 0 objects, got %d", len(responseObjects))
	}

	if response.Headers["X-Total-Count"] != "0" {
		t.Fatalf("expected X-Total-Count=0, got %s", response.Headers["X-Total-Count"])
	}

	// Now, update the object modifying the lifecycleStatus to "Launched"
	upd = map[string]any{
		"@type":           resourceName,
		"version":         "1.2",
		"lifecycleStatus": "Launched",
	}
	bUpd, _ = json.Marshal(upd)
	request = &Request{
		Method:       "PATCH",
		Action:       ActionUPDATE,
		APIfamily:    apiFamily,
		APIVersion:   "v4",
		ResourceName: resourceName,
		ID:           id,
		Body:         bUpd,
		QueryParams:  nil,
		AuthUser:     *sellerUser,
	}

	response = s.UpdateTMFObject(context.Background(), request)
	if response.StatusCode != http.StatusOK {
		t.Fatalf("update expected 200, got %d", response.StatusCode)
	}

	// List again
	request = &Request{
		Method:       "GET",
		Action:       ActionLIST,
		APIfamily:    apiFamily,
		APIVersion:   "v4",
		ResourceName: resourceName,
		ID:           "",
		Body:         nil,
		QueryParams:  url.Values{},
		AuthUser:     *foreignUser,
	}

	response = s.ListTMFObjects(context.Background(), request)
	if response.StatusCode != http.StatusOK {
		t.Fatalf("list expected 200, got %d", response.StatusCode)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatalf("list expected 200, got %d", response.StatusCode)
	}

	responseObjects, ok = response.Body.([]repository.TMFObjectMap)
	if !ok {
		t.Fatalf("expected list of objects, got %T", response.Body)
	}
	if len(responseObjects) != 1 {
		t.Fatalf("expected 1 object, got %d", len(responseObjects))
	}

	if response.Headers["X-Total-Count"] == "" {
		t.Fatalf("missing X-Total-Count header")
	}
	if response.Headers["X-Total-Count"] != "1" {
		t.Fatalf("expected X-Total-Count=1, got %s", response.Headers["X-Total-Count"])
	}

	// List with fields=none (should reduce fields per item)

	fields := url.Values{"fields": []string{"none"}}
	requestQP := &Request{
		Method:       "GET",
		Action:       ActionLIST,
		APIfamily:    apiFamily,
		APIVersion:   "v4",
		ResourceName: resourceName,
		ID:           "",
		Body:         nil,
		QueryParams:  fields,
		AuthUser:     *sellerUser,
	}

	response = s.ListTMFObjects(context.Background(), requestQP)
	if response.StatusCode != http.StatusOK {
		t.Fatalf("list expected 200, got %d", response.StatusCode)
	}

	items, ok := response.Body.([]repository.TMFObjectMap)
	if !ok || len(items) == 0 {
		t.Fatalf("expected list of items, got %T", response.Body)
	}
	// Expect minimal keys present
	item := items[0]
	if item["id"] == nil || item["href"] == nil || item["version"] == nil || item["lastUpdate"] == nil || item["@type"] == nil {
		t.Fatalf("fields=none did not include minimal fields")
	}

	// Delete
	dReq := &Request{
		Method:       "DELETE",
		Action:       ActionDELETE,
		APIfamily:    apiFamily,
		APIVersion:   "v4",
		ResourceName: resourceName,
		ID:           id,
		Body:         nil,
		QueryParams:  nil,
		AuthUser:     *sellerUser,
	}

	response = s.DeleteTMFObject(context.Background(), dReq)
	if response.StatusCode != http.StatusNoContent {
		t.Fatalf("delete expected 204, got %d", response.StatusCode)
	}

	// Get after delete -> 404

	request = &Request{
		Method:       "GET",
		Action:       ActionREAD,
		APIfamily:    apiFamily,
		APIVersion:   "v4",
		ResourceName: resourceName,
		ID:           id,
		Body:         nil,
		QueryParams:  nil,
		AuthUser:     *sellerUser,
	}
	response = s.GetTMFObject(context.Background(), request)
	if response.StatusCode != http.StatusNotFound {
		t.Fatalf("get after delete expected 404, got %d", response.StatusCode)
	}
}

// Try to create an object which is not managed by the server operator
func TestLocalForeignerCreate(t *testing.T) {
	s := newLocalTestService(t)

	apiFamily := "productCatalogManagement"
	resourceName := "productOffering"

	s.Features.VerifyJWTSignature = true

	// Authenticate as the server operator, who can do everything
	authUser, err := s.ProcessAccessToken(testServerOperatorAccessToken)
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

	// But now try to create a foreign object authenticating as a normal user
	// To authenticate as another seller
	foreignUser := &types.AuthUser{
		OrganizationIdentifier: "VATES-11111111K",
		IsAuthenticated:        true,
		IsLEAR:                 true,
		ProductCreatePower:     true,
		ProductUpdatePower:     true,
		ProductDeletePower:     true,
	}

	cReq = &Request{
		Method:       "POST",
		Action:       ActionCREATE,
		APIfamily:    apiFamily,
		APIVersion:   "v4",
		ResourceName: resourceName,
		ID:           "",
		Body:         bCreate,
		QueryParams:  nil,
		AuthUser:     *foreignUser,
	}

	cResp = s.CreateTMFObject(context.Background(), cReq)
	if cResp.StatusCode != http.StatusForbidden {
		t.Fatalf("create expected 403, got %d", cResp.StatusCode)
	}

}

// newISBEDEVTestService creates a new service for testing with ISBE DEV configuration
func newISBEDEVTestService(t *testing.T) *Service {
	t.Helper()

	// Set the environment variable ISBETMF_ADMIN_TOKEN to the admin token
	os.Setenv("ISBETMF_ADMIN_TOKEN", testServerOperatorAccessToken)

	configuration, err := config.LoadConfig("isbedev", true)
	if err != nil {
		t.Fatalf("failed to load configuration: %v", err)
	}

	dbLayer, err := repository.NewDBService(":memory:", configuration.ServerOperatorOrganizationIdentifier)
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

func newLocalTestService(t *testing.T) *Service {
	t.Helper()

	dbLayer, err := repository.NewDBService(":memory:", testServerOperator)
	if err != nil {
		t.Fatalf("create test db: %v", err)
	}

	// Create service struct directly (no external verifier)
	s := &Service{
		adminToken:                           testServerOperatorAccessToken,
		ServerOperatorDid:                    testServerOperator,
		ServerOperatorName:                   "Foundation Operator",
		ServerOperatorOrganizationIdentifier: testServerOperator,
		ServerOperatorCountry:                "SPAIN",
		proxyEnabled:                         false,

		storage: dbLayer,
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
			VerifyJWTSignature: false,
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

func TestCreateGenericObjectPublishesEvent(t *testing.T) {
	s := newLocalTestService(t)

	// To authenticate as a seller
	sellerUser, err := s.ProcessAccessToken(testSellerAccessToken)
	if err != nil {
		t.Fatalf("failed to process access token: %v", err)
	}

	// Replace notifications manager with one that uses fake delivery
	memStore := notifications.NewMemoryStore()
	fdel := &fakeDelivery{}
	s.notif = notifications.NewManager(memStore, fdel)

	// Create a subscription to receive create events
	sub := &notifications.Subscription{
		ID:         "sub1",
		APIFamily:  "productCatalogManagement",
		Callback:   "http://localhost:9991/listener/ProductOfferingCreateEvent",
		EventTypes: []string{"ProductOfferingCreateEvent"},
		Headers:    map[string]string{"x-auth-token": "abc123"},
	}
	if _, err := s.notif.CreateSubscription("productCatalogManagement", sub); err != nil {
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

	req := &Request{
		Method:       "POST",
		Action:       ActionCREATE,
		APIfamily:    "productCatalogManagement",
		APIVersion:   "v4",
		ResourceName: resourceName,
		ID:           "",
		Body:         b,
		QueryParams:  nil,
		AuthUser:     *sellerUser,
	}

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

// TestInvalidResourcename tests that ListGenericObjects returns an empty JSON array and proper X-Total-Count header
func TestInvalidResourcename(t *testing.T) {
	s := newLocalTestService(t)
	resourceName := "TestResource"
	apiFamily := "productCatalogManagement"

	// List objects for a resource that doesn't exist (should return 400)
	request := newReq("GET", "LIST", apiFamily, resourceName, "", nil, url.Values{})
	lResp := s.ListTMFObjects(context.Background(), request)

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

	dbLayer, err := repository.NewDBService(":memory:", configuration.ServerOperatorOrganizationIdentifier)
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
