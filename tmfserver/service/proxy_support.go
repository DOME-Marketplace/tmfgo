// Copyright 2023 Jesus Ruiz. All rights reserved.
// Use of this source code is governed by an Apache 2.0
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hesusruiz/tmforum/config"
	"github.com/hesusruiz/tmforum/internal/errl"
	repo "github.com/hesusruiz/tmforum/tmfserver/repository"
)

// createRemoteOrLocalObject creates an object in the remote server and then in the local database, if the proxy is enabled.
// Othewise, it just creates the object in the local database.
func (svc *Service) createRemoteOrLocalObject(ctx context.Context, req *Request, objMap repo.TMFObjectMap) *Response {

	// Create the object only in the local database if the proxy is not enabled
	if !svc.proxyEnabled {
		// Convert object to storage representation
		repoObject := objMap.ToTMFRecord(req.ResourceName)

		if err := svc.storage.CreateObject(req, repoObject); err != nil {
			if errors.Is(err, &ErrObjectExists{}) {
				return ErrorResponsef(http.StatusBadRequest, "object %s already exists: %w", objMap.ID(), err)
			} else {
				return ErrorResponsef(http.StatusInternalServerError, "failed to create object locally: %w", err)
			}
		}

		// Set the headers of the reply to the caller, as per TMF specs
		headers := map[string]string{
			"Location": objMap.Href(),
		}
		slog.Debug("Object created successfully", slog.String("id", objMap.ID()), slog.String("resourceName", req.ResourceName), slog.String("location", objMap.Href()))

		return &Response{StatusCode: http.StatusCreated, Headers: headers, Body: objMap}

	}

	// With the proxy enabled, first create the object in the remote server and then locally
	// We do not have to worry about transaction integrity, because if the remote server fails, we do not create the object locally
	// If the local server fails, the object will be eventually updated in our cache later, for other operations against the object.

	remoteObjectMap, err := svc.tmfClient.TMFPost(ctx, req, objMap)
	if err != nil {

		if apiErr, ok := errors.AsType[*ApiError](err); ok {
			return &Response{StatusCode: apiErr.StatusCode(), Body: apiErr}
		}

		return ErrorResponsef(http.StatusInternalServerError, "failed to proxy request: %w", err)
	}

	// Prepare the object for the database
	lastUpdate := remoteObjectMap.LastUpdate()

	// It is an error if the remote server does not return a 'lastUpdate', but we just log a warning
	if lastUpdate == "" {
		slog.Warn("remote server did not return lastUpdate, fixing it", slog.String("id", remoteObjectMap.ID()))
	}

	// Create the new object in the local database
	if err := svc.storage.CreateObject(req, remoteObjectMap.ToTMFRecord(req.ResourceName)); err != nil {
		// If we get an error, just log the error because we return the object created remotely
		slog.Error("failed to create local object", slog.String("id", remoteObjectMap.ID()), slog.String("resourceName", req.ResourceName), slog.String("location", remoteObjectMap.Href()))
	}

	// Set the headers of the reply to the caller, as per TMF specs
	headers := map[string]string{
		"Location": remoteObjectMap.Href(),
	}
	slog.Debug("Object created successfully", slog.String("id", remoteObjectMap.ID()), slog.String("resourceName", req.ResourceName), slog.String("location", remoteObjectMap.Href()))

	return &Response{StatusCode: http.StatusCreated, Headers: headers, Body: remoteObjectMap}

}

// updateRemoteOrLocalObject updates an existing object in the remote server and then in the local database,
// if the proxy is enabled.
// existingRecord is only used if the proxy is not enabled.
// Otherwise, it just updates the object in the local database after merging with the RFC7396 patch.
func (svc *Service) updateRemoteOrLocalObject(ctx context.Context, req *Request, existingRecord *repo.TMFRecord, patch repo.TMFObjectMap) *Response {
	var existingObjectMap repo.TMFObjectMap
	var err error

	existingObjectMap, err = existingRecord.ToTMFObjectMap()
	if err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to unmarshal existing object content: %w", err)
	}

	if svc.proxyEnabled {
		remoteObjectMap, err := svc.tmfClient.TMFPatch(ctx, req, patch)
		if err != nil {
			if svc.LogLevel() >= 3 {
				// Pretty-pring the path object
				patchJSON, _ := json.MarshalIndent(patch, "", "  ")
				fmt.Println("Path object:", string(patchJSON))
			}
			return ErrorResponsef(http.StatusInternalServerError, "failed to proxy request: %w", err)
		}

		// Set the existingObjectMap to the remoteObjectMap, so we store the object as it was received from the remote server
		existingObjectMap = remoteObjectMap

	} else {
		// Merge patch into existing object using RFC7396
		svc.mergeRFC7396(existingObjectMap, patch)
	}

	objectID := req.ID

	// For organization resources in local database, the objectID is the organization identifier.
	if svc.IsISBE() {
		if req.ResourceName == "organization" {
			if strings.HasPrefix(req.ID, "urn:ngsi-ld:organization:") && !strings.HasPrefix(req.ID, "urn:ngsi-ld:organization:did:elsi:") {
				objectID = "urn:ngsi-ld:organization:did:elsi:" + strings.TrimPrefix(req.ID, "urn:ngsi-ld:organization:")
			}
		}
	}

	// Prepare the object for the database
	existingVersion := existingObjectMap.Version()
	existingLastUpdate := existingObjectMap.LastUpdate()

	// It is an error if the remote server does not return a 'lastUpdate', but we fix it and log a warning
	if existingLastUpdate == "" {
		slog.Warn("remote server did not return lastUpdate, fixing it", slog.String("id", objectID))

		now := time.Now()
		existingLastUpdate = now.Format(time.RFC3339Nano)
		existingObjectMap["lastUpdate"] = existingLastUpdate
	}

	// We will store the serialized object in the database
	existingObjectContent, err := json.Marshal(existingObjectMap)
	if err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to marshal object content for update: %w", err)
	}

	existingObject := &repo.TMFRecord{
		ID:         objectID,
		Type:       req.ResourceName,
		Version:    existingVersion,
		APIVersion: req.APIVersion,
		LastUpdate: existingLastUpdate,
		Content:    existingObjectContent,
		CreatedAt:  existingRecord.CreatedAt,
		UpdatedAt:  time.Now().Unix(),
	}

	if err := svc.storage.UpdateObject(req, existingObject); err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to update object in service: %w", err)
	}

	return &Response{StatusCode: http.StatusOK, Body: existingObjectMap}
}

// listRemoteObjects retrieves objects from the remote TMF server, filters them based on authorization,
// caches them locally, and returns the requested page.
func (svc *Service) listRemoteObjects(ctx context.Context, req *Request, userLimit, userOffset int, fieldSet map[string]bool) (
	responseObjects []repo.TMFObjectMap, responseHeaders map[string]string, diagnosticObjects []repo.ValidationResult, err error) {

	// Delete the attribute selection for the query to the upstream server. We will receive full objects and
	// perform attribute selection ourselves. This is because we want to store the full objects in our local cache.
	req.QueryParams.Del("fields")

	// We forward the same access token that we received from the user
	upstreamHeaders := map[string]string{
		"Authorization": "Bearer " + req.AuthUser.AccessToken,
		"Accept":        "application/json",
		"Content-Type":  "application/json",
	}

	// Check if the user wants diagnostic information, which is specified in the query string as '?diagnostic=true'
	// This is not standard TMF, we use it to report on quality of data
	diagnostic := req.QueryParams.Has("diagnostic")
	if diagnostic {
		// We do not want to forward this to the remote server
		req.QueryParams.Del("diagnostic")
	}

	responseObjects = make([]repo.TMFObjectMap, 0)
	diagnosticObjects = make([]repo.ValidationResult, 0)
	var offsetCounter int
	invalidObjects := 0

	// To be able to perform authorization and return the number of objects requested from the user, we must request objects from the server
	// starting from the beginning, i.e. offset=0, and continue until we have enough objects to satisfy the user's request.
	// This is because the objects are not ordered in any particular way, so we cannot just request the objects we need.
	// We request objects in pages of 100 objects at a time.

	// Delete the paging parameters as we will handle them ourselves
	req.QueryParams.Del("offset")
	req.QueryParams.Del("limit")

	remoteTotalObjects := 0
	pageSize := svc.tmfClient.PageSize()
	pageOffset := 0
	for {

		// Get one page of objects from the remote server
		receivedObjects, totalObjects, err := svc.tmfClient.TMFGetList(ctx, req.ResourceName, req.QueryParams, pageSize, pageOffset, upstreamHeaders, nil, req.HealthRequest)
		if err != nil {
			return nil, nil, nil, errl.Errorf("upstream server failed with error: %w", err)
		}

		remoteTotalObjects = totalObjects

		if !req.HealthRequest {
			slog.Debug("received objects from remote", "num_objects", len(receivedObjects))
		}

		// Stop requesting pages if we received zero objects.
		if len(receivedObjects) == 0 {
			break
		}

		// We check each object to see if the user can access it.
		// Additionally, we cache all the objects received independently of the user's access.
		for _, receivedObject := range receivedObjects {

			// Perform validations on the received object
			validations := receivedObject.Validate(req.ResourceName)

			// If the object is invalid, add it to the diagnostic objects and delete it
			if len(validations.Errors) > 0 {
				invalidObjects++
				diagnosticObjects = append(diagnosticObjects, validations)
				if diagnostic {
					receivedObject["validationErrors"] = validations.Errors
					responseObjects = append(responseObjects, receivedObject)
				}

				// Delete the offending object if we are not in production
				if svc.environment != config.DOME_PRO {
					pathPrefix, err := config.ExternalUpstreamTMFPath(req.ResourceName)
					if err != nil {
						slog.Error("failed to get path prefix", "error", err, "resourceName", req.ResourceName)
						continue
					}
					path := fmt.Sprintf("%s/%s", pathPrefix, receivedObject.ID())

					resp, _, err := svc.tmfClient.Delete(ctx, path, upstreamHeaders)
					if err != nil || resp.StatusCode >= 300 {
						slog.Error("failed to delete invalid object", "error", err, "status_code", resp.StatusCode, "path", path)
						continue
					}

					slog.Info("Invalid object deleted", "resourceName", req.ResourceName, "id", receivedObject.ID())
				}

				continue
			}

			// Convert object to storage representation to save it in the local database
			storageObject := receivedObject.ToTMFRecord(req.ResourceName)
			if err := svc.storage.UpsertObject(req, storageObject); err != nil {
				if !errors.Is(err, &ErrObjectExists{}) {
					invalidObjects++
					slog.Error("error saving object in local database", "error", err)
					continue
				}
			}

			// Check if the user is authorized to access the object
			authorized, err := svc.checkAuthorization(svc.ruleEngine, req, receivedObject)
			if !authorized {
				slog.Debug("object not authorized", "id", receivedObject.ID(), "error", err)
				// Add diagnostic info if not authorized
				invalidObjects++
				diagnosticObjects = append(diagnosticObjects, repo.ValidationResult{
					ObjectID:   receivedObject.ID(),
					ObjectType: req.ResourceName,
					Valid:      false,
					Errors: []repo.ValidationError{
						{
							Field:   receivedObject.ID(),
							Message: fmt.Sprintf("object %s not authorized: %s", receivedObject.ID(), errl.Error(err)),
							Code:    "NOT_AUTHORIZED",
						},
					},
				})
				continue
			}

			if offsetCounter < userOffset {
				offsetCounter++
				continue
			}

			// Apply attribute selection, according to what the user specified in the query
			receivedObject = svc.applyAttributeSelection(receivedObject, fieldSet)
			responseObjects = append(responseObjects, receivedObject)

			// Stop validating objects if we have enough objects to satisfy the user's request
			if userLimit >= 0 && len(responseObjects) >= userLimit {
				break
			}
		}

		// Stop requesting objects from the remote server if we have enough objects to satisfy the user's request.
		if userLimit >= 0 && len(responseObjects) >= userLimit {
			break
		}
		pageOffset += pageSize
	}

	responseHeaders = map[string]string{
		"X-Result-Count": strconv.Itoa(len(responseObjects)),
		"X-Total-Count":  strconv.Itoa(remoteTotalObjects),
	}

	if !req.HealthRequest {
		slog.Debug("Remote objects listed", slog.Int("valid", len(responseObjects)), slog.Int("invalid", invalidObjects), slog.String("resourceName", req.ResourceName))
	}

	return responseObjects, responseHeaders, diagnosticObjects, nil
}

// listLocalObjects retrieves TMF objects from the local database, filters them based on authorization,
// handles attribute selection, and returns the requested page.
func (svc *Service) listLocalObjects(req *Request, userLimit, userOffset int, fieldSet map[string]bool) ([]repo.TMFObjectMap, map[string]string, *Response) {
	// Set the offset and limit for the database query
	req.QueryParams.Set("offset", strconv.Itoa(userOffset))
	req.QueryParams.Set("limit", strconv.Itoa(userLimit))

	// storageObjects will contain the objects retrieved from the local database, filtered by access control
	// They contain all attributes, which will be filtered later
	storageObjects, total, err := svc.storage.ListObjects(req, func(storageObject *repo.TMFRecord) bool {
		// Convert to internal object representation
		objMap, err := storageObject.ToTMFObjectMap()
		if err != nil {
			slog.Error("failed to unmarshal object content for listing", "error", err)
			return false
		}

		// Check if the user is authorized to access the object
		authorized, err := svc.checkAuthorization(svc.ruleEngine, req, objMap)
		if !authorized {
			slog.Debug("object not authorized", "id", storageObject.ID, "error", err)
			return false
		}
		return true
	})

	if err != nil {
		return nil, nil, ErrorResponsef(http.StatusInternalServerError, "failed to list objects from local database: %w", err)
	}

	// Now we create the actual objects containing the attributes selected by the user (plus the compulsory ones)
	responseObjects := make([]repo.TMFObjectMap, 0, len(storageObjects))
	for _, storageObject := range storageObjects {
		// Convert to internal object representation
		objectMap, err := storageObject.ToTMFObjectMap()
		if err != nil {
			return nil, nil, ErrorResponsef(http.StatusInternalServerError, "failed to unmarshal object content for listing: %w", err)
		}

		// Apply attribute selection
		objectMap = svc.applyAttributeSelection(objectMap, fieldSet)
		responseObjects = append(responseObjects, objectMap)
	}

	responseHeaders := map[string]string{
		"X-Result-Count": strconv.Itoa(len(responseObjects)),
		"X-Total-Count":  strconv.Itoa(total),
	}

	if !req.HealthRequest {
		slog.Debug("Objects listed successfully", slog.Int("count", len(responseObjects)), slog.String("resourceName", req.ResourceName))
	}

	return responseObjects, responseHeaders, nil
}

// getLocalOrRemoteObject retrieves the object from the database.
// If the proxy is enabled and the object is not found locally or is stale, we try to get it from the remote server.
// If the object is not found anywhere, it returns a nil object and no error.
// There is no way to force the retrieval from the remote server if the object exists locally and is fresh enough.
func (svc *Service) getLocalOrRemoteObject(ctx context.Context, req *Request) (*repo.TMFRecord, error) {

	objectID := req.ID

	// For organization resources in local database, the objectID is the organization identifier.
	if svc.IsISBE() {
		if req.ResourceName == "organization" {
			if strings.HasPrefix(req.ID, "urn:ngsi-ld:organization:") && !strings.HasPrefix(req.ID, "urn:ngsi-ld:organization:did:elsi:") {
				objectID = "urn:ngsi-ld:organization:did:elsi:" + strings.TrimPrefix(req.ID, "urn:ngsi-ld:organization:")
			}
		}
	}

	// Check if we have the object locally
	obj, err := svc.storage.GetObject(req, objectID, req.ResourceName)
	if err != nil {
		return nil, errl.Errorf("failed to get object %s from local service: %w", objectID, err)
	}

	// If proxy is not enabled, we return whatever we found, which is nil if object is not found.
	// Not finding an object is not an error at this level but the caller is responsible for returning a 404 error to the user.
	if !svc.proxyEnabled {
		return obj, nil
	}

	// If we found the object and the database record is not stale, return it
	if obj != nil {
		if time.Since(obj.GetUpdatedAt()) < svc.fressness {
			slog.Debug("Object found in local database and fresh")
			return obj, nil
		}
		slog.Debug("Object found in local database but stale, retrieving from remote")
	} else {
		slog.Debug("Object not found in local database, retrieving from remote")
	}

	// The object was not found or is stale, so we have to retrieve remotely and update the local database
	remoteObj, err := svc.getRemoteObject(ctx, req)
	if err != nil {
		return nil, errl.Errorf("failed to get object %s from remote service: %w", req.ID, err)
	}

	if remoteObj == nil {
		slog.Debug("object not found in remote service", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))
		return nil, nil
	}

	// Store the object locally and return it to caller
	if err := svc.storage.CreateObject(req, remoteObj); err != nil {
		slog.Error("failed to cache object", slog.Any("error", err))
		// Return the stale object or nil
		return remoteObj, nil
	}

	slog.Debug("Object retrieved from remote and cached successfully", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))
	return remoteObj, nil

}

func (svc *Service) getRemoteObject(ctx context.Context, req *Request) (*repo.TMFRecord, error) {
	slog.Debug("retrieving object from remote", slog.String("id", req.ID))

	// Send the access token
	headers := map[string]string{
		"Authorization": "Bearer " + req.AuthUser.AccessToken,
	}

	// Build the path for the request according to TMForum specs
	pathPrefix, err := config.ExternalUpstreamTMFPath(req.ResourceName)
	if err != nil {
		return nil, errl.Errorf("failed to get path prefix: %w", err)
	}

	path := fmt.Sprintf("%s/%s", pathPrefix, req.ID)

	// Send the request to the remote with our HTTP Client
	resp, body, err := svc.tmfClient.Get(ctx, path, headers)
	if err != nil {
		return nil, errl.Errorf("failed to proxy request: %w", err)
	}

	// Not found is not an error at this level, but the caller must check for a nil object
	if resp.StatusCode == 404 {
		return nil, nil
	}

	// TODO: handle caching replies, with status codes in the 3xx range
	if resp.StatusCode >= 300 {
		return nil, errl.Errorf("remote server returned error: %s", string(body))
	}

	// Build the object from the reply
	receivedObjectMap, err := repo.NewTMFObjectMapFromBytes(req.ResourceName, body)
	if err != nil {
		return nil, errl.Errorf("failed to bind request body: %w", err)
	}

	// Prepare the object for the database
	id := receivedObjectMap.ID()
	version := receivedObjectMap.Version()
	lastUpdate := receivedObjectMap.LastUpdate()

	// It is an error if the remote server does not return a 'lastUpdate', but we fix it and log a warning
	if lastUpdate == "" {
		slog.Warn("remote server did not return lastUpdate, fixing it", slog.String("id", id))

		now := time.Now()
		lastUpdate = now.Format(time.RFC3339Nano)
		receivedObjectMap["lastUpdate"] = lastUpdate
	}

	// Instead of storing what we receive, we store a compact serialization of the JSON object with possible updated fields
	receivedContent, err := json.Marshal(receivedObjectMap)
	if err != nil {
		return nil, errl.Errorf("failed to marshal object content: %w", err)
	}

	obj := repo.NewTMFRecord(id, req.ResourceName, version, req.APIVersion, lastUpdate, receivedContent)

	return obj, nil
}
