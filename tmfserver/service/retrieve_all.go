package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"strconv"
	"time"

	"github.com/DOME-Marketplace/tmfgo/config"
	"github.com/DOME-Marketplace/tmfgo/internal/errl"
	repo "github.com/DOME-Marketplace/tmfgo/tmfserver/repository"
	"github.com/DOME-Marketplace/tmfgo/types"
)

const defaultDeleteInvalid = true

func (svc *Service) ScheduleRetrieveAll() {
	if !svc.proxyEnabled || svc.tmfClient == nil {
		return
	}
	go func() {
		for {
			// Retrieve all public resources at the beginning and every 15 minutes
			err := svc.RetrieveAll(context.Background(), defaultDeleteInvalid)
			if err != nil {
				slog.Error("failed to retrieve all product offerings", "error", err)
			}
			time.Sleep(15 * time.Minute)
		}
	}()
}

// RetrieveAll retrieves all TMF objects of a given type.
func (svc *Service) RetrieveAll(ctx context.Context, deleteInvalid bool) error {
	if !svc.proxyEnabled || svc.tmfClient == nil {
		return nil
	}

	totalNumber := 0
	totalInvalidObjects := 0

	type counts struct {
		valid   int
		invalid int
	}

	summaryCounts := make(map[string]counts)

	publicResources := types.GetPublicResources()
	for _, resource := range publicResources {

		// Create a Request for listing all ProductOfferings
		req := &Request{
			Method:       "GET",
			Action:       ActionLIST,
			APIfamily:    "productCatalogManagement",
			APIVersion:   "v4",
			ResourceName: resource,
			QueryParams: url.Values{
				"limit": []string{"10000"},
			},
		}

		slog.Info("Retrieve All", "resource", resource)

		// Parse pagination parameters
		userLimit := 10000
		userOffset := 0

		// Retrieve objects
		receivedObjects, _, invalidObjects, err := svc.listRemoteObjectsRobust(ctx, req, userLimit, userOffset)
		if err != nil {
			slog.Error("Failed to retrieve objects", "error", err, "resource", resource)
			continue
		}

		if len(invalidObjects) == 0 {
			slog.Info("Retrieved objects", "resourceName", resource, "valid", len(receivedObjects))
		} else {
			slog.Error("Retrieved objects", "resourceName", resource, "valid", len(receivedObjects), "invalid", len(invalidObjects))
		}

		summaryCounts[resource] = counts{
			valid:   len(receivedObjects),
			invalid: len(invalidObjects),
		}
		totalNumber += len(receivedObjects)
		totalInvalidObjects += len(invalidObjects)

		if len(invalidObjects) > 0 {

			// Iterate through validation results and print the errors
			for _, vr := range invalidObjects {

				// Delete the offending object if deleteInvalid is true
				if deleteInvalid && resource != types.Category {
					pathPrefix, err := config.ExternalUpstreamTMFPath(req.ResourceName)
					if err != nil {
						slog.Error("failed to get path prefix", "error", err, "resourceName", req.ResourceName)
						continue
					}
					path := fmt.Sprintf("%s/%s", pathPrefix, vr.ObjectID)

					upstreamHeaders := map[string]string{
						"Authorization": "Bearer " + req.AuthUser.AccessToken,
						"Accept":        "application/json",
						"Content-Type":  "application/json",
					}

					fmt.Printf("##### Deleting invalid object: %s %s\n", req.ResourceName, vr.ObjectID)

					resp, _, err := svc.tmfClient.Delete(ctx, path, upstreamHeaders)
					if err != nil || resp.StatusCode >= 300 {
						slog.Error("failed to delete invalid object", "error", err, "status_code", resp.StatusCode, "path", path)
					}

					// Print a useful message for each validation error
					for _, valError := range vr.Errors {
						fmt.Printf("    Reason: %s %s %s\n", valError.Field, valError.Message, valError.Code)
					}

				} else {

					fmt.Printf("##### Invalid object not deleted: %s %s\n", req.ResourceName, vr.ObjectID)

					// Print a useful message for each validation error
					for _, valError := range vr.Errors {
						fmt.Printf("    Reason: %s %s %s\n", valError.Field, valError.Message, valError.Code)
					}
				}
			}
		}

	}

	slog.Info("Retrieved All", "valid", totalNumber, "invalid", totalInvalidObjects)
	for resource, c := range summaryCounts {
		fmt.Printf("%s valid: %d, invalid: %d\n", resource, c.valid, c.invalid)
	}

	return nil
}

func (svc *Service) listRemoteObjectsRobust(ctx context.Context, req *Request, userLimit, userOffset int) (
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
	diagnostic := false

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
			slog.Info("received objects from remote", "num_objects", len(receivedObjects))
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

			// If errors, record them and continue with next object
			if len(validations.Errors) > 0 {
				invalidObjects++
				diagnosticObjects = append(diagnosticObjects, validations)
				if diagnostic {
					receivedObject["validationErrors"] = validations.Errors
					responseObjects = append(responseObjects, receivedObject)
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

	return responseObjects, responseHeaders, diagnosticObjects, nil
}
