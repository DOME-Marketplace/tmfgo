package service

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

// UpdateTMFObject updates an existing TMF object using generalized parameters.
func (svc *Service) UpdateTMFObject(ctx context.Context, req *Request) *Response {
	slog.Debug("UpdateGenericObject called", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))

	// Authentication is required for update operations
	if errorResponse := svc.requiresAuthentication(req); errorResponse != nil {
		return errorResponse
	}

	// Parse request body
	incomingObjectMap, errorResponse := svc.parseRequestBodyForUpdate(req)
	if errorResponse != nil {
		slog.Error(string(req.Body))
		return errorResponse
	}

	if slog.Default().Enabled(ctx, slog.LevelDebug) {
		data, _ := json.MarshalIndent(incomingObjectMap, "", "  ")
		slog.Debug(string(data))
	}

	// Ensure update metadata
	if errorResponse := svc.verifyObjectOnUpdate(req, incomingObjectMap); errorResponse != nil {
		data, _ := json.MarshalIndent(incomingObjectMap, "", "  ")
		slog.Error(string(data))
		return errorResponse
	}

	// Retrieve existing object
	existingRecord, err := svc.getLocalOrRemoteObject(ctx, req)
	if err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "error retrieving existing object %s for update: %w", req.ID, err)
	}
	if existingRecord == nil {
		return ErrorResponsef(http.StatusNotFound, "object %s not found", req.ID)
	}

	// Convert from storage representation to TMF object map
	existingObjectMap, err := existingRecord.ToTMFObjectMap()
	if err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "error unmarshalling existing object record: %w", err)
	}

	// Authorization (using existing object)
	if errorResponse := svc.authorizeAction(req, existingObjectMap); errorResponse != nil {
		return errorResponse
	}

	// Object Update (Local or Remote)
	response := svc.updateRemoteOrLocalObject(ctx, req, existingRecord, incomingObjectMap)

	// Notification
	if response.StatusCode == http.StatusOK {
		eventType := toEventType(req.ResourceName, "AttributeValueChangeEvent")
		eventPayload := buildEventPayload(req, eventType, response.Body)
		svc.notif.PublishEvent(req.APIfamily, eventType, eventPayload)
	}

	return response
}
