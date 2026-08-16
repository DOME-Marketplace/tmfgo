package fiber

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gofiber/fiber/v2"
	svc "github.com/hesusruiz/tmforum/tmfserver/service"
	"github.com/hesusruiz/tmforum/types"
)

// ServiceInterface defines the methods used by the handler
type ServiceInterface interface {
	CreateGenericObject(req *svc.Request) *svc.Response
	GetGenericObject(req *svc.Request) *svc.Response
	UpdateGenericObject(req *svc.Request) *svc.Response
	DeleteGenericObject(req *svc.Request) *svc.Response
	ListGenericObjects(req *svc.Request) *svc.Response
}

// testHandler wraps the original handler to use an interface
type testHandler struct {
	service ServiceInterface
}

func (h *testHandler) ListGenericObjects(c *fiber.Ctx) error {
	jwtToken := ExtractJWTToken(c.Get("Authorization"))

	queryParams, _ := url.ParseQuery(string(c.Request().URI().QueryString()))
	req := &svc.Request{
		Method:       c.Method(),
		Action:       svc.ActionLIST,
		ResourceName: c.Params("resourceName"),
		QueryParams:  queryParams,
		AuthUser: types.AuthUser{
			IsAuthenticated:        true,
			IsLEAR:                 true,
			ProductCreatePower:     true,
			ProductUpdatePower:     true,
			ProductDeletePower:     true,
			OrganizationIdentifier: "1234567890",
			Organization:           "Test Organization",
			CommonName:             "Test User",
			Country:                "US",
			EmailAddress:           "[EMAIL_ADDRESS]",
			SerialNumber:           "1234567890",
			AccessToken:            jwtToken,
		},
	}

	resp := h.service.ListGenericObjects(req)
	return SendResponse(c, resp)
}

// mockService is a mock service for testing
type mockService struct{}

func (m *mockService) CreateGenericObject(req *svc.Request) *svc.Response {
	return &svc.Response{StatusCode: http.StatusNotImplemented}
}

func (m *mockService) GetGenericObject(req *svc.Request) *svc.Response {
	return &svc.Response{StatusCode: http.StatusNotImplemented}
}

func (m *mockService) UpdateGenericObject(req *svc.Request) *svc.Response {
	return &svc.Response{StatusCode: http.StatusNotImplemented}
}

func (m *mockService) DeleteGenericObject(req *svc.Request) *svc.Response {
	return &svc.Response{StatusCode: http.StatusNotImplemented}
}

func (m *mockService) ListGenericObjects(req *svc.Request) *svc.Response {
	// Return an empty list response
	headers := make(map[string]string)
	headers["X-Total-Count"] = "0"

	return &svc.Response{
		StatusCode: http.StatusOK,
		Headers:    headers,
		Body:       []map[string]any{}, // Empty array
	}
}

// TestListGenericObjectsEmptyList tests that the handler returns an empty JSON array and proper X-Total-Count header
func TestListGenericObjectsEmptyList(t *testing.T) {
	// Create handler with mock service
	mockSvc := &mockService{}
	handler := &testHandler{service: mockSvc}

	// Create fiber app
	app := fiber.New()

	// Setup route
	app.Get("/:resourceName", handler.ListGenericObjects)

	// Create request
	req := httptest.NewRequest("GET", "/NonExistentResource", nil)

	// Execute request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("fiber app test failed: %v", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	// Check X-Total-Count header
	totalCount := resp.Header.Get("X-Total-Count")
	if totalCount != "0" {
		t.Fatalf("expected X-Total-Count=0, got %s", totalCount)
	}

	// Check content type
	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf("expected application/json content type, got %s", contentType)
	}

	// Check body is an empty JSON array
	var body []map[string]any
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if len(body) != 0 {
		t.Fatalf("expected empty array, got %d items", len(body))
	}
}
