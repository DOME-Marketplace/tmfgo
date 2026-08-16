package service

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/hesusruiz/tmforum/config"
	repo "github.com/hesusruiz/tmforum/tmfserver/repository"
)

// DeleteTMFObject deletes a TMF object, first in the remote server and then locally.
func (svc *Service) DeleteTMFObject(ctx context.Context, req *Request) *Response {
	var err error
	slog.Debug("DeleteGenericObject called", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))

	// Authentication is required for delete operations
	if errorResponse := svc.requiresAuthentication(req); errorResponse != nil {
		return errorResponse
	}

	// ************************************************************************************************
	// We need the existing object to see if the user is authorised to delete it
	// ************************************************************************************************

	var existingObject *repo.TMFRecord

	// Retrieve existing object, locally or remotely
	existingObject, err = svc.getLocalOrRemoteObject(ctx, req)
	if err != nil {
		// TODO: check the return code from remote server and reply accordingly
		return ErrorResponsef(http.StatusBadRequest, "failed to get existing object for update: %w", err)
	}

	// If nothing to delete, return 404
	if existingObject == nil {
		return ErrorResponsef(http.StatusNotFound, "object %s not found", req.ID)
	}

	// Convert to a type-safe map representation to facilitate manipulation
	existingObjectMap, err := existingObject.ToTMFObjectMap()
	if err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to unmarshal existing object content: %w", err)
	}

	// ************************************************************************************************
	// Before performing the action, check if the user can perform the operation on the object,
	// based on the rules defined by the user in the policy engine.
	// ************************************************************************************************

	if authorized, err := svc.checkAuthorization(svc.ruleEngine, req, existingObjectMap); !authorized {
		return ErrorResponsef(http.StatusForbidden,
			"user %s is not authorized, object: %s, error: %w",
			req.AuthUser.OrganizationIdentifier,
			existingObjectMap,
			err,
		)
	} else {
		slog.Debug("caller is authorized to delete object", "reason", err)
	}

	// ##########################################################
	// Delete the object
	// ##########################################################

	// Delete the object in the remote server, if the proxy is enabled
	if svc.proxyEnabled {
		// Send the authentication header
		headers := map[string]string{
			"Authorization": "Bearer " + req.AuthUser.AccessToken,
		}

		pathPrefix, err := config.ExternalUpstreamTMFPath(req.ResourceName)
		if err != nil {
			return ErrorResponsef(http.StatusInternalServerError, "failed to get path prefix: %w", err)
		}
		path := fmt.Sprintf("%s/%s", pathPrefix, req.ID)

		resp, body, err := svc.tmfClient.Delete(ctx, path, headers)
		if err != nil {
			return ErrorResponsef(http.StatusInternalServerError, "failed to proxy request: %w", err)
		}

		if resp.StatusCode >= 300 {
			return &Response{
				StatusCode: resp.StatusCode,
				Body:       body,
			}
		}

		slog.Debug("Object deleted from remote server successfully", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))

	}

	// Delete the object in the local database
	if err := svc.DeleteObject(req.ID, req.ResourceName); err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to delete object %s from service: %w", req.ID, err)
	}

	slog.Debug("Object deleted successfully from local database", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))

	// Send TMForum notification
	eventType := toEventType(req.ResourceName, "DeleteEvent")
	minimal := map[string]any{
		"id":    req.ID,
		"@type": req.ResourceName,
		"href":  fmt.Sprintf("/tmf-api/%s/%s/%s/%s", req.APIfamily, req.APIVersion, req.ResourceName, req.ID),
	}
	eventPayload := buildEventPayload(req, eventType, minimal)
	svc.notif.PublishEvent(req.APIfamily, eventType, eventPayload)

	return &Response{StatusCode: http.StatusNoContent}

}
