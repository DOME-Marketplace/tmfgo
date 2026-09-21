// Package fiber implements the Fiber web server handlers for the TMF API.
// It is a very thin layer on top of the service package, which is the real implementation of the TM Forum APIs.
// In this way, we can easily switch to another web framework by just implementing a new handler package,
// or even supporting more than one web framework at the same time.

// This also allows to have a different transport mechanism (like grpc), or request mechanism (like JSON-RPC),
// without changing the service implementation.
package fiber

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"encoding/json"
	"log/slog"

	"github.com/DOME-Marketplace/tmfgo/internal/errl"
	svc "github.com/DOME-Marketplace/tmfgo/tmfserver/service"
	"github.com/gofiber/fiber/v2"
)

// Handler is the handler for the TMF API (both V4 and V5).
type Handler struct {
	service *svc.Service
}

// NewHandler creates a new handler.
func NewHandler(app *fiber.App, s *svc.Service) *Handler {

	h := &Handler{service: s}

	h.registerRoutes(app)

	return h

}

// registerRoutes registers the routes for the TMF API.
func (h *Handler) registerRoutes(app *fiber.App) {

	// Mock listener for local testing (accepts any listener path)
	app.Post("/listener/*", h.MockListener)

	// Health check
	app.Get("/health", h.Health)

	// Group routes for TMF API (both V4 and V5)
	tmfApi := app.Group("/tmf-api/:apiFamily/:apiVersion")

	// Notifications Hub routes
	tmfApi.Post("/hub", h.CreateHubSubscription)
	tmfApi.Delete("/hub/:id", h.DeleteHubSubscription)

	// Generalized routes for TMF API resources
	// Collection operations (List and Create)
	tmfApi.Get("/:resourceName", h.ListTMFObjects)
	tmfApi.Post("/:resourceName", h.CreateTMFObject)
	tmfApi.Put("/:resourceName", h.ReplaceTMFObject)

	// Individual resource operations (Get, Update, Delete)
	tmfApi.Get("/:resourceName/:id", h.GetTMFObject)
	tmfApi.Patch("/:resourceName/:id", h.UpdateTMFObject)
	tmfApi.Delete("/:resourceName/:id", h.DeleteGenericObject)

}

// Health simulates a LIST request with limit = 1 to exercise the full path as much as possible
func (h *Handler) Health(c *fiber.Ctx) error {

	// Create a request to list a catalog with limit=1
	queryParams, _ := url.ParseQuery("limit=1")
	req := &svc.Request{
		HealthRequest: true,
		Method:        "GET",
		Action:        svc.ActionLIST,
		APIfamily:     "productCatalogManagement",
		APIVersion:    "v4",
		ResourceName:  "catalog",
		QueryParams:   queryParams,
	}

	// Create a context with a timeout of 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp := h.service.ListTMFObjects(ctx, req)
	return SendResponse(c, resp)

}

// CreateHubSubscription creates a new notification subscription (hub)
func (h *Handler) CreateHubSubscription(c *fiber.Ctx) error {

	req, err := h.parseRequest(c)
	if err != nil {
		resp := svc.ErrorResponsef(http.StatusBadRequest, "error parsing request: %w", errl.Error(err))
		return SendResponse(c, resp)
	}

	resp := h.service.CreateHubSubscription(req)
	return SendResponse(c, resp)
}

// DeleteHubSubscription deletes an existing notification subscription (hub)
func (h *Handler) DeleteHubSubscription(c *fiber.Ctx) error {

	req, err := h.parseRequest(c)
	if err != nil {
		resp := svc.ErrorResponsef(http.StatusBadRequest, "error parsing request: %w", errl.Error(err))
		return SendResponse(c, resp)
	}

	resp := h.service.DeleteHubSubscription(req)
	return SendResponse(c, resp)
}

// MockListener is a minimal endpoint to receive notifications locally for testing
func (h *Handler) MockListener(c *fiber.Ctx) error {
	path := string(c.Request().URI().Path())
	body := c.Body()
	if len(body) > 0 {
		var payload any
		if err := json.Unmarshal(body, &payload); err == nil {
			slog.Info("listener received event", slog.String("path", path), slog.Int("bytes", len(body)), slog.Any("body", payload))
		} else {
			slog.Info("listener received event", slog.String("path", path), slog.Int("bytes", len(body)), slog.String("bodyRaw", string(body)))
		}
	} else {
		slog.Info("listener received event", slog.String("path", path), slog.Int("bytes", 0))
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// CreateTMFObject creates a new TMF object.
func (h *Handler) CreateTMFObject(c *fiber.Ctx) error {

	resourceName := c.Params("resourceName")
	if resourceName == "hub" {
		slog.Debug("handling hub creation")
		return h.CreateHubSubscription(c)
	}

	req, err := h.parseRequest(c)
	if err != nil {
		resp := svc.ErrorResponsef(http.StatusBadRequest, "error parsing request: %w", errl.Error(err))
		return SendResponse(c, resp)
	}

	// An id in the request is forbidden
	if req.ID != "" {
		resp := svc.ErrorResponsef(http.StatusBadRequest, "id is not allowed in the request")
		return SendResponse(c, resp)
	}

	// Create a context with a timeout of 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp := h.service.CreateTMFObject(ctx, req)
	return SendResponse(c, resp)
}

// ReplaceTMFObject is like a CreateTMFObject but allows an id to be specified by the caller
// It creates or replaces the object if it exists.
func (h *Handler) ReplaceTMFObject(c *fiber.Ctx) error {

	req, err := h.parseRequest(c)
	if err != nil {
		resp := svc.ErrorResponsef(http.StatusBadRequest, "error parsing request: %w", errl.Error(err))
		return SendResponse(c, resp)
	}

	// The request MUST have an id
	if req.ID == "" {
		resp := svc.ErrorResponsef(http.StatusBadRequest, "id is required in the request")
		return SendResponse(c, resp)
	}

	// Create a context with a timeout of 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp := h.service.ReplaceTMFObject(ctx, req)
	return SendResponse(c, resp)
}

// UpdateTMFObject updates an existing TMF object.
func (h *Handler) UpdateTMFObject(c *fiber.Ctx) error {

	req, err := h.parseRequest(c)
	if err != nil {
		resp := svc.ErrorResponsef(http.StatusBadRequest, "error parsing request: %w", errl.Error(err))
		return SendResponse(c, resp)
	}

	// The request MUST have an id
	if req.ID == "" {
		resp := svc.ErrorResponsef(http.StatusBadRequest, "id is required in the request")
		return SendResponse(c, resp)
	}

	// Create a context with a timeout of 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp := h.service.UpdateTMFObject(ctx, req)
	return SendResponse(c, resp)
}

// GetTMFObject retrieves a TMF object.
func (h *Handler) GetTMFObject(c *fiber.Ctx) error {

	req, err := h.parseRequest(c)
	if err != nil {
		resp := svc.ErrorResponsef(http.StatusBadRequest, "error parsing request: %w", errl.Error(err))
		return SendResponse(c, resp)
	}

	// The request MUST have an id
	if req.ID == "" {
		resp := svc.ErrorResponsef(http.StatusBadRequest, "id is required in the request")
		return SendResponse(c, resp)
	}

	// Create a context with a timeout of 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp := h.service.GetTMFObject(ctx, req)
	return SendResponse(c, resp)
}

// DeleteGenericObject deletes a TMF object.
func (h *Handler) DeleteGenericObject(c *fiber.Ctx) error {
	resourceName := c.Params("resourceName")
	if resourceName == "hub" {
		slog.Debug("handling hub deletion")
		return h.DeleteHubSubscription(c)
	}

	req, err := h.parseRequest(c)
	if err != nil {
		resp := svc.ErrorResponsef(http.StatusBadRequest, "error parsing request: %w", errl.Error(err))
		return SendResponse(c, resp)
	}

	// The request MUST have an id
	if req.ID == "" {
		resp := svc.ErrorResponsef(http.StatusBadRequest, "id is required in the request")
		return SendResponse(c, resp)
	}

	// Create a context with a timeout of 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp := h.service.DeleteTMFObject(ctx, req)
	return SendResponse(c, resp)
}

// ListTMFObjects retrieves all TMF objects.
func (h *Handler) ListTMFObjects(c *fiber.Ctx) error {

	req, err := h.parseRequest(c)
	if err != nil {
		resp := svc.ErrorResponsef(http.StatusBadRequest, "error parsing request: %w", errl.Error(err))
		return SendResponse(c, resp)
	}

	// An id in the request is forbidden
	if req.ID != "" {
		resp := svc.ErrorResponsef(http.StatusBadRequest, "id is not allowed in the request")
		return SendResponse(c, resp)
	}

	// Create a context with a timeout of 1 minute
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	resp := h.service.ListTMFObjects(ctx, req)
	return SendResponse(c, resp)
}

// parseRequest parses the Fiber request and returns a framework-agnostic and transport-agnostic service request.
// The service Request allows the service layer to be used for any framework (http, grpc, etc) and any transport.
func (h *Handler) parseRequest(c *fiber.Ctx) (*svc.Request, error) {

	// resourceName is compulsory in all requests
	resourceName := c.Params("resourceName")
	if resourceName == "" {
		return nil, errl.Errorf("resourceName is compulsory")
	}

	// Extract API version from the path parameter
	apiVersion := strings.ToLower(c.Params("apiVersion"))
	if apiVersion == "" {
		return nil, errl.Errorf("apiVersion is compulsory")
	}

	// Extract the JWT token from the Authorization header
	jwtToken := ExtractJWTToken(c.Get("Authorization"))

	// Extract caller info from the token claims in the payload, verifying the signature.
	// If no token is present, authUser will be theGuestUser.
	// We do not perform authentication at this step, it is done in the service layer.
	authUser, err := h.service.ProcessAccessToken(jwtToken)
	if err != nil {
		return nil, errl.Errorf("invalid access token: %w", err)
	}

	// Acondition the query parameters, according to the TMF specs
	queryParams, err := ParseTMFRequestQuery(c)
	if err != nil {
		return nil, errl.Errorf("error parsing the request query: %w", err)
	}

	// Parses the 'id' in the path of the TMF requests which have it (PATCH, GET, PUT, DELETE)
	// If the id is not present in the path, 'idParam' is the empty string
	idParam, err := url.QueryUnescape(c.Params("id"))
	if err != nil {
		return nil, errl.Errorf("error parsing the id parameter: %w", err)
	}

	method := c.Method()
	action := svc.HttpActionFromMethod(method, idParam)

	req := &svc.Request{
		Method:       method,
		Action:       action,
		APIfamily:    c.Params("apiFamily"),
		APIVersion:   apiVersion,
		ResourceName: resourceName,
		ID:           idParam,
		QueryParams:  queryParams,
		AuthUser:     *authUser,
	}

	// Set the Body if the request potentially has one (POST, PATCH, PUT)
	if c.Method() == fiber.MethodPost || c.Method() == fiber.MethodPatch || c.Method() == fiber.MethodPut {
		bodyBytes := c.Body()
		// Make an immutable copy of the body to be used during processing of the request.
		immutableBody := make([]byte, len(bodyBytes))
		copy(immutableBody, bodyBytes)
		req.Body = immutableBody
	}

	return req, nil

}

func SendResponse(c *fiber.Ctx, resp *svc.Response) error {
	for key, value := range resp.Headers {
		c.Set(key, value)
	}
	if resp.Body != nil {
		return c.Status(resp.StatusCode).JSON(resp.Body)
	}
	return c.SendStatus(resp.StatusCode)
}

// ParseTMFParams takes url.Values and ensures every key has a slice
// where comma-separated values are broken into individual elements.
// It supports queries like these:
//
//	GET /api/troubleTicket/42/?fields=description,status&pepe=1&fields=version
//
// where the 'fields' key will contain after parsing: ['description'. 'status', 'version']
// It does not attempt to de-duplicate values, which is not a problem for TMF logic.
func ParseTMFParams(v url.Values) {
	for key, slices := range v {
		var flattened []string
		for _, s := range slices {
			// Split by comma
			parts := strings.SplitSeq(s, ",")
			for p := range parts {
				trimmed := strings.TrimSpace(p)
				if trimmed != "" {
					flattened = append(flattened, trimmed)
				}
			}
		}
		// Overwrite the existing slice with the flattened one
		v[key] = flattened
	}
}

// ParseTMFRequestQuery processes the URL query as per TMF requirements,
// where comma-separated values are broken into individual elements.
// It supports queries like these:
//
//	GET /api/troubleTicket/42/?fields=description,status&pepe=1&fields=version
//
// where the 'fields' key will contain after parsing: ['description'. 'status', 'version']
// It does not attempt to de-duplicate values, which is not a problem for TMF logic.
func ParseTMFRequestQuery(c *fiber.Ctx) (url.Values, error) {
	queryParams, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return nil, errl.Errorf("parsing the request query: %w", err)
	}

	ParseTMFParams(queryParams)

	return queryParams, nil
}

// ExtractJWTToken extracts the JWT token from the Authorization header.
// It handles both "Bearer <token>" and raw token formats.
// If the token is not found, it returns an empty string.
func ExtractJWTToken(authHeader string) string {
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	return tokenString
}
