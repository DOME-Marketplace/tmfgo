package service

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/hesusruiz/tmforum/internal/errl"
)

// GetTMFObject retrieves a TMF object using generalized parameters.
//
// The function performs the following checks and operations:
//
// 1.  **Authentication**:
//   - It processes the Access Token to get the caller's information. Public (unauthenticated) access is allowed.
//
// 2.  **Object Retrieval**:
//   - It first tries to retrieve the object from the local database.
//   - If the object is not found and proxy mode is enabled, it tries to fetch it from the remote TMF server.
//   - If the object is found remotely, it is cached in the local database.
//   - If the object is not found locally or remotely, it returns a 404 Not Found error.
//
// 3.  **Authorization**:
//   - It calls the Policy Decision Point (PDP) to check if the user is authorized to read the object.
//   - If the user is not authorized, it returns a 403 Forbidden error.
//
// 4.  **Response**:
//   - It handles partial field selection using the "fields" query parameter.
//   - It returns a 200 OK response with the (potentially filtered) object in the body.
func (svc *Service) GetTMFObject(ctx context.Context, req *Request) *Response {
	slog.Debug("GetTMFObject called", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))

	// Authentication: anonymous access is allowed, so we do not check the existence of the access token

	// Retrieve the object from the database. If it is not found, we try to get it from the remote server (if proxy is enabled).
	// If the object is not found, we return a 404 error.
	existingStorageObject, err := svc.getLocalOrRemoteObject(ctx, req)
	if err != nil {
		// This is an unexpected error, so we return a server error
		return ErrorResponsef(http.StatusInternalServerError, "failed to get object from service: %w", errl.Error(err))
	}

	// Check if the object was found (which is not an error in the previous step)
	if existingStorageObject == nil {
		return ErrorResponsef(http.StatusNotFound, "object not found")
	}

	// Convert to a type-safe map representation to facilitate manipulation
	existingObjectMap, err := existingStorageObject.ToTMFObjectMap()
	if err != nil {
		if len(existingStorageObject.Validations.Errors) > 0 {
			return ErrorResponsef(http.StatusInternalServerError, "validation errors found: %s", existingStorageObject.Validations.String())
		} else {
			return ErrorResponsef(http.StatusInternalServerError, "failed to unmarshal existing object content: %w", errl.Error(err))
		}
	}

	// ************************************************************************************************
	// Before returning the object to the user, check if the user has access to the object,
	// based on the rules defined by the owner of the object in the policy engine.
	// ************************************************************************************************

	if authorized, err := svc.checkAuthorization(svc.ruleEngine, req, existingObjectMap); !authorized {
		return ErrorResponsef(http.StatusForbidden,
			"user %s is not authorized, object: %s, error: %w",
			req.AuthUser.OrganizationIdentifier,
			existingObjectMap,
			err,
		)
	} else {
		slog.Debug("caller is authorized to read object", "reason", err)
	}

	// ************************************************************************************************
	// Now we can proceed.
	// ************************************************************************************************

	// Handle partial field selection
	fields := req.QueryParams["fields"]
	if len(fields) > 0 {

		// The user can specify "none", but we do not have to care about it, because if he specifies a field that does not exist,
		// we will return the object with the mandatory fields anyway.

		// Create a set of fields for quick lookup
		fieldSet := make(map[string]bool, len(fields))
		for _, f := range fields {
			fieldSet[strings.TrimSpace(f)] = true
		}

		// Always include id, href, lastUpdate or lastModified, version, @type and lifecycleStatus
		fieldSet["id"] = true
		fieldSet["href"] = true
		fieldSet["lastUpdate"] = true
		fieldSet["lastModified"] = true
		fieldSet["version"] = true
		fieldSet["@type"] = true
		fieldSet["lifecycleStatus"] = true

		// Instead of deleting the fields we don't want, we replace the object with a new one containing only the fields we want.
		filteredObject := make(map[string]any)
		for key, value := range existingObjectMap {
			if fieldSet[key] {
				filteredObject[key] = value
			}
		}
		existingObjectMap = filteredObject
	}

	slog.Debug("Object retrieved successfully", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))
	return &Response{StatusCode: http.StatusOK, Body: existingObjectMap}
}
