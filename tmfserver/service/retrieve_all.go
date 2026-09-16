package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"

	"github.com/hesusruiz/tmforum/config"
	"github.com/hesusruiz/tmforum/internal/errl"
	repo "github.com/hesusruiz/tmforum/tmfserver/repository"
	"github.com/hesusruiz/tmforum/types"
)

// RetrieveAll retrieves all TMF objects of a given type.
func (svc *Service) RetrieveAll(ctx context.Context) *Response {

	// Create a Request for listing all ProductOfferings
	req := &Request{
		Method:       "GET",
		Action:       ActionLIST,
		APIfamily:    "productCatalogManagement",
		APIVersion:   "v4",
		ResourceName: types.ProductOffering,
		QueryParams: url.Values{
			"limit": []string{"10000"},
		},
	}

	// Make sure the resource is supported
	res := types.GetResourceDefinition(req.ResourceName)
	if res == nil {
		return ErrorResponsef(http.StatusBadRequest, "resource type %s not supported", req.ResourceName)
	}

	diagnostic := true

	// Parse pagination parameters
	userLimit, userOffset, err := svc.parsePaginationParams(req)
	if err != nil {
		return ErrorResponsef(http.StatusBadRequest, "failed to parse pagination parameters: %w", err)
	}

	// If the user specified explicitly limit=0, return an empty list. This may be used to test the API without returning all the objects.
	if userLimit == 0 {
		return &Response{StatusCode: http.StatusOK, Body: []repo.TMFObjectMap{}}
	}

	// Parse field selection parameters, which are the fields to be returned in the response.
	fieldsParam, _ := req.QueryParams["fields"]
	fieldSet := svc.parseFieldsParam(fieldsParam)

	var responseData []repo.TMFObjectMap
	var responseHeaders map[string]string

	// Retrieve objects
	var diagnosticObjects []repo.ValidationResult
	responseData, responseHeaders, diagnosticObjects, err = svc.listRemoteObjectsRobust(ctx, req, userLimit, userOffset, fieldSet)
	if err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to proxy request: %w", err)
	}
	if diagnostic || len(diagnosticObjects) > 0 {
		// return &Response{StatusCode: http.StatusOK, Headers: responseHeaders, Body: diagnosticObjects}
		return &Response{
			StatusCode: http.StatusOK,
			Headers:    responseHeaders,
			Body:       responseData,
		}
	}

	return &Response{
		StatusCode: http.StatusOK,
		Headers:    responseHeaders,
		Body:       responseData,
	}
}

func (svc *Service) listRemoteObjectsRobust(ctx context.Context, req *Request, userLimit, userOffset int, fieldSet map[string]bool) (
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
	diagnostic := true

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

			// We continue if we did not yet retrieve enough objects to be at the requested offset
			if offsetCounter < userOffset {
				offsetCounter++
				continue
			}

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
