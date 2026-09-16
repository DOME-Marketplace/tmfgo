package service

import (
	"context"
	"log/slog"
	"net/http"

	repo "github.com/hesusruiz/tmforum/tmfserver/repository"
	"github.com/hesusruiz/tmforum/types"
)

// ListTMFObjects retrieves all TMF objects of a given type.
func (svc *Service) ListTMFObjects(ctx context.Context, req *Request) *Response {
	if !req.HealthRequest {
		slog.Debug("ListGenericObjects called", slog.String("resourceName", req.ResourceName), slog.String("queryParams", req.QueryParams.Encode()))
	}

	// Make sure the resource is supported
	res := types.GetResourceDefinition(req.ResourceName)
	if res == nil {
		return ErrorResponsef(http.StatusBadRequest, "resource type %s not supported", req.ResourceName)
	}

	// Check if the user wants diagnostic information, which is specified in the query string as '?diagnostic=true'
	// This is not standard TMF, we use it to report on quality of data
	diagnostic := req.QueryParams.Has("diagnostic")

	// If the requested resource is public and the user is not authenticated, add the "lifecycleStatus=Launched" filter
	if !req.AuthUser.IsAuthenticated && res.Public {
		req.QueryParams.Set("lifecycleStatus", "Launched")
	}

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
	var resp *Response

	// Retrieve objects (Remote or Local)
	if svc.proxyEnabled {
		var err error
		var diagnosticObjects []repo.ValidationResult
		responseData, responseHeaders, diagnosticObjects, err = svc.listRemoteObjects(ctx, req, userLimit, userOffset, fieldSet)
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
	} else {
		responseData, responseHeaders, resp = svc.listLocalObjects(req, userLimit, userOffset, fieldSet)
		if resp != nil {
			return resp
		}
	}

	return &Response{
		StatusCode: http.StatusOK,
		Headers:    responseHeaders,
		Body:       responseData,
	}
}
