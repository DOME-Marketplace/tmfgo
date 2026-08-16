package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

// CreateTMFObject creates a new TMF object using generalized parameters.
func (svc *Service) CreateTMFObject(ctx context.Context, req *Request) *Response {
	slog.Debug("CreateTMFObject called", slog.String("apiFamily", req.APIfamily), slog.String("resourceName", req.ResourceName))

	// Authentication is required for create operations
	if errorResponse := svc.requiresAuthentication(req); errorResponse != nil {
		return errorResponse
	}

	// Parse request body
	incomingObjectMap, errorResponse := svc.parseRequestBodyForCreateAndReplace(req)
	if errorResponse != nil {
		slog.Error(string(req.Body))
		return errorResponse
	}

	if slog.Default().Enabled(context.Background(), slog.LevelDebug) {
		data, _ := json.MarshalIndent(incomingObjectMap, "", "  ")
		fmt.Println(string(data))
	}

	// Ensure TMF metadata (ID, version, href, etc.)
	if errorResponse := svc.verifyObjectOnPOST(req, incomingObjectMap); errorResponse != nil {
		bodyJson, _ := json.MarshalIndent(incomingObjectMap, "", "  ")
		slog.Error(string(bodyJson))
		return errorResponse
	}

	// Authorization
	if errorResponse := svc.authorizeAction(req, incomingObjectMap); errorResponse != nil {
		bodyJson, _ := json.MarshalIndent(incomingObjectMap, "", "  ")
		slog.Error(string(bodyJson))
		return errorResponse
	}

	// Object Creation (Remote or Local)
	response := svc.createRemoteOrLocalObject(ctx, req, incomingObjectMap)

	// Notification to subscribers
	if response.StatusCode == http.StatusCreated {
		eventType := toEventType(req.ResourceName, "CreateEvent")
		eventPayload := buildEventPayload(req, eventType, response.Body)
		svc.notif.PublishEvent(req.APIfamily, eventType, eventPayload)
	}

	return response
}
