package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hesusruiz/tmforum/tmfserver/notifications"
)

// CreateHubSubscription creates a new notification subscription (hub) for an API family.
func (svc *Service) CreateHubSubscription(req *Request) *Response {
	// Authentication is required for create operations
	if errorResponse := svc.requiresAuthentication(req); errorResponse != nil {
		return errorResponse
	}

	// Parse incoming body
	var body map[string]any
	if err := json.Unmarshal(req.Body, &body); err != nil {
		return ErrorResponsef(http.StatusBadRequest, "failed to bind request body: %w", err)
	}

	callback, _ := body["callback"].(string)
	if callback == "" {
		return ErrorResponsef(http.StatusBadRequest, "callback is required")
	}

	var eventTypes []string
	if raw, ok := body["eventTypes"].([]any); ok {
		for _, v := range raw {
			if s, ok := v.(string); ok {
				eventTypes = append(eventTypes, s)
			}
		}
	}

	headers := make(map[string]string)
	if hmap, ok := body["headers"].(map[string]any); ok {
		for k, v := range hmap {
			if sv, ok := v.(string); ok {
				headers[strings.ToLower(k)] = sv
			}
		}
	}

	query, _ := body["query"].(string)

	// Build subscription
	id := uuid.NewString()
	sub := &notifications.Subscription{
		ID:         id,
		APIFamily:  req.APIfamily,
		Callback:   callback,
		EventTypes: eventTypes,
		Headers:    headers,
		Query:      query,
	}

	_, err := svc.notif.CreateSubscription(req.APIfamily, sub)
	if err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to create subscription: %w", err)
	}

	resp := map[string]any{
		"id":         sub.ID,
		"callback":   sub.Callback,
		"eventTypes": sub.EventTypes,
		"query":      sub.Query,
		"href":       fmt.Sprintf("/tmf-api/%s/%s/hub/%s", req.APIfamily, req.APIVersion, sub.ID),
	}
	if token, ok := sub.Headers["x-auth-token"]; ok && token != "" {
		resp["headers"] = map[string]any{"x-auth-token": token}
	}

	return &Response{StatusCode: http.StatusCreated, Body: resp}
}

// DeleteHubSubscription deletes a subscription by id for an API family.
func (svc *Service) DeleteHubSubscription(req *Request) *Response {
	// Authenticate like write operations
	if errorResponse := svc.requiresAuthentication(req); errorResponse != nil {
		return errorResponse
	}

	if req.ID == "" {
		return ErrorResponsef(http.StatusBadRequest, "id is required")
	}

	if err := svc.notif.DeleteSubscription(req.APIfamily, req.ID); err != nil {
		return ErrorResponsef(http.StatusNotFound, "failed to delete subscription: %w", err)
	}

	return &Response{StatusCode: http.StatusNoContent}
}

// toEventType composes the TMF event type name from a resource name and a suffix.
func toEventType(resourceName, suffix string) string {
	// Convert resourceName to PascalCase
	parts := strings.Split(resourceName, "-")
	if len(parts) > 1 {
		for i, p := range parts {
			if len(p) == 0 {
				continue
			}
			parts[i] = strings.ToUpper(p[:1]) + strings.ToLower(p[1:])
		}
		return strings.Join(parts, "") + suffix
	}
	if len(resourceName) == 0 {
		return suffix
	}
	return strings.ToUpper(resourceName[:1]) + resourceName[1:] + suffix
}

// buildEventPayload builds a generic TMF event envelope.
func buildEventPayload(req *Request, eventType string, resource any) map[string]any {
	return map[string]any{
		"eventId":      uuid.NewString(),
		"eventTime":    time.Now().Format(time.RFC3339Nano),
		"eventType":    eventType,
		"apiFamily":    req.APIfamily,
		"resourceName": req.ResourceName,
		"resourceId":   req.ID,
		"resourcePath": fmt.Sprintf("/tmf-api/%s/%s/%s", req.APIfamily, req.APIVersion, req.ResourceName),
		"event": map[string]any{
			"resource": resource,
		},
	}
}
