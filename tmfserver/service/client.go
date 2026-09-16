package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/hesusruiz/tmforum/config"
	"github.com/hesusruiz/tmforum/internal/errl"
	"github.com/hesusruiz/tmforum/tmfserver/repository"
)

// TMFClientConfig holds the configuration for the tmfclient service
type TMFClientConfig struct {
	// BaseURL of the remote TMForum server
	BaseURL string `json:"base_url" yaml:"base_url"`

	// Timeout in seconds for HTTP requests
	Timeout int `json:"timeout" yaml:"timeout"`

	// Default page size for requests to the remote server
	PageSize int `json:"page_size" yaml:"page_size"`

	// If we are running in internal mode we do not use the BaseURL
	InternalMode bool `json:"internal_mode" yaml:"internal_mode"`
}

// TMFClient is a client for the TMForum API.
type TMFClient struct {
	config *TMFClientConfig
	client *http.Client
}

// NewClient creates a new client.
func NewClient(config *TMFClientConfig) *TMFClient {
	if config.Timeout == 0 {
		config.Timeout = 10
	}
	if config.PageSize == 0 {
		config.PageSize = 100
	}
	return &TMFClient{
		config: config,
		client: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Second,
		},
	}
}

func (c *TMFClient) PageSize() int {
	return c.config.PageSize
}

func (c *TMFClient) TMFPost(ctx context.Context, req *Request, objMap repository.TMFObjectMap) (repository.TMFObjectMap, error) {

	requestBody, err := json.Marshal(objMap)
	if err != nil {
		return nil, errl.Errorf("failed to marshall object: %w", err)
	}

	path, err := config.ExternalUpstreamTMFPath(req.ResourceName)
	if err != nil {
		return nil, errl.Errorf("failed to get path prefix: %w", err)
	}

	headers := map[string]string{
		"Authorization": "Bearer " + req.AuthUser.AccessToken,
		"Content-Type":  "application/json",
	}

	resp, responseBody, err := c.Post(ctx, path, requestBody, headers)
	if err != nil {
		return nil, errl.Errorf("remote server returned error: %w", err)
	}

	if resp.StatusCode != http.StatusCreated {
		responseString := string(responseBody)
		slog.Error("unexpected status code from remote server",
			slog.Int("status_code", resp.StatusCode),
			slog.String("response_body", responseString),
			slog.String("path", path))

		reqBodyJSON, _ := json.MarshalIndent(objMap, "", "  ")
		slog.Error(string(reqBodyJSON))

		// Unmarshall response as APIError
		apiError := &ApiError{}
		if jsonErr := json.Unmarshal(responseBody, apiError); jsonErr != nil {
			_, code, reason := getHTTPStatusInfo(resp.StatusCode)
			return nil, NewApiError(resp.StatusCode, code, reason, string(responseBody), resp.Status, "")
		}

		apiError.statusCode = resp.StatusCode
		return nil, apiError

	}

	obj, err := repository.NewTMFObjectMapFromBytes(req.ResourceName, responseBody)
	if err != nil {
		return nil, errl.Errorf("failed to bind request body: %w", err)
	}

	validation := obj.Validate(req.ResourceName)
	if len(validation.Errors) > 0 {
		var errs []error
		for _, vErr := range validation.Errors {
			errs = append(errs, errl.Errorf("validation error on field '%s': %s (code: %s)", vErr.Field, vErr.Message, vErr.Code))
		}
		return nil, errors.Join(errs...)
	}

	return obj, nil

}

// TMFPut is used to forward PUT requests to the remote server.
// req is the incoming request and objMap is the object to be sent to the remote server.
func (c *TMFClient) TMFPut(ctx context.Context, req *Request, objMap repository.TMFObjectMap) (repository.TMFObjectMap, error) {

	requestBody, err := json.Marshal(objMap)
	if err != nil {
		return nil, errl.Errorf("failed to marshall object: %w", err)
	}

	path, err := config.ExternalUpstreamTMFPath(req.ResourceName)
	if err != nil {
		return nil, errl.Errorf("failed to get path prefix: %w", err)
	}

	headers := map[string]string{
		"Authorization": "Bearer " + req.AuthUser.AccessToken,
		"Content-Type":  "application/json",
	}

	resp, responseBody, err := c.Put(ctx, path, requestBody, headers)
	if err != nil {
		return nil, errl.Errorf("remote server returned error: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		responseString := string(responseBody)
		slog.Error("unexpected status code from remote server",
			slog.Int("status_code", resp.StatusCode),
			slog.String("response_body", responseString),
			slog.String("path", path))

		reqBodyJSON, _ := json.MarshalIndent(objMap, "", "  ")
		slog.Error(string(reqBodyJSON))

		// Unmarshall response as APIError
		apiError := &ApiError{}
		if jsonErr := json.Unmarshal(responseBody, apiError); jsonErr != nil {
			_, code, reason := getHTTPStatusInfo(resp.StatusCode)
			return nil, NewApiError(resp.StatusCode, code, reason, string(responseBody), resp.Status, "")
		}

		apiError.statusCode = resp.StatusCode
		return nil, apiError
	}

	obj, err := repository.NewTMFObjectMapFromBytes(req.ResourceName, responseBody)
	if err != nil {
		return nil, errl.Errorf("failed to bind request body: %w", err)
	}

	validation := obj.Validate(req.ResourceName)
	if len(validation.Errors) > 0 {
		var errs []error
		for _, vErr := range validation.Errors {
			errs = append(errs, errl.Errorf("validation error on field '%s': %s (code: %s)", vErr.Field, vErr.Message, vErr.Code))
		}
		return nil, errors.Join(errs...)
	}

	return obj, nil

}

// TMFPatch is used to forward PATCH requests to the remote server.
// req is the incoming request and patchMap is the object to be sent to the remote server.
func (c *TMFClient) TMFPatch(ctx context.Context, req *Request, patchMap repository.TMFObjectMap) (repository.TMFObjectMap, error) {

	requestBody, err := json.Marshal(patchMap)
	if err != nil {
		return nil, errl.Errorf("failed to marshall object: %w", err)
	}

	// Get the resource path to the remote server
	// TODO: add support for requests internal to the DOME server, which go to the kubernetes pods
	pathPrefix, err := config.ExternalUpstreamTMFPath(req.ResourceName)
	if err != nil {
		return nil, errl.Errorf("failed to get path prefix: %w", err)
	}

	path := fmt.Sprintf("%s/%s", pathPrefix, req.ID)

	headers := map[string]string{
		"Authorization": "Bearer " + req.AuthUser.AccessToken,
		"Content-Type":  "application/json",
	}

	resp, responseBody, err := c.Patch(ctx, path, requestBody, headers)
	if err != nil {
		return nil, errl.Errorf("remote server returned error: %w", err)
	}

	if resp.StatusCode >= 300 {
		responseString := string(responseBody)
		slog.Error("unexpected status code from remote server",
			slog.Int("status_code", resp.StatusCode),
			slog.String("response_body", responseString),
			slog.String("path", path))

		reqBodyJSON, _ := json.MarshalIndent(patchMap, "", "  ")
		slog.Error(string(reqBodyJSON))

		// Unmarshall response as APIError
		apiError := &ApiError{}
		if jsonErr := json.Unmarshal(responseBody, apiError); jsonErr != nil {
			_, code, reason := getHTTPStatusInfo(resp.StatusCode)
			return nil, NewApiError(resp.StatusCode, code, reason, string(responseBody), resp.Status, "")
		}

		apiError.statusCode = resp.StatusCode
		return nil, apiError
	}

	// Build an object from the reply
	obj, err := repository.NewTMFObjectMapFromBytes(req.ResourceName, responseBody)
	if err != nil {
		return nil, errl.Errorf("failed to bind request body: %w", err)
	}

	// And validate the object
	validation := obj.Validate(req.ResourceName)
	if len(validation.Errors) > 0 {
		var errs []error
		for _, vErr := range validation.Errors {
			errs = append(errs, errl.Errorf("validation error on field '%s': %s (code: %s)", vErr.Field, vErr.Message, vErr.Code))
		}
		return nil, errors.Join(errs...)
	}

	return obj, nil

}

// processObject is a function provided by the caller of GetAllObjectsOfType which processes an object of a specific type.
// It takes the object type and the object being processed as input.
// It returns the processed object, a boolean indicating whether to continue processing, and an error if any.
// If processObject returns false, it means that the caller wants to stop processing the objects.
type processObject func(obj repository.TMFObjectMap) (repository.TMFObjectMap, bool, error)

// TMFGetList retrieves a list of TMF objects from the remote server.
// It does not perform any validation of the objects, but delegates it to the processObject callback provided by the caller.
func (c *TMFClient) TMFGetList(
	ctx context.Context,
	resourceName string,
	queryParams url.Values,
	pageSize int,
	pageOffset int,
	headers map[string]string,
	processObject processObject,
	healthRequest bool,
) ([]repository.TMFObjectMap, int, error) {

	remoteTotalObjects := -1

	// Build the parameters to send to the remote server
	baseParams := queryParams.Encode()

	// Build the base path including parameters for the request to the remote server
	// The path is terminated with '&' or '?' because we will add the paging parameters later
	basePath, err := config.ExternalUpstreamTMFPath(resourceName)
	if err != nil {
		return nil, 0, errl.Errorf("failed to get path prefix: %w", err)
	}
	if baseParams != "" {
		basePath += "?" + baseParams + "&"
	} else {
		basePath += "?"
	}

	// Build the full path and ask the server for the specified page of objects
	path := basePath + "limit=" + strconv.Itoa(pageSize) + "&offset=" + strconv.Itoa(pageOffset)

	if !healthRequest {
		slog.Debug("sending request to remote", "path", path)
	}

	resp, body, err := c.Get(ctx, path, headers)
	if err != nil {
		return nil, 0, errl.Errorf("remote server returned error: %w", err)
	}

	// Check the content type of the response and return an error if it is not JSON
	if !strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		return nil, 0, errl.Errorf("remote server returned non-JSON content type: %s", resp.Header.Get("Content-Type"))
	}

	if resp.StatusCode >= 300 {
		slog.Error("unexpected status code from remote server", slog.Int("status_code", resp.StatusCode), slog.String("response_body", string(body)), slog.String("path", path))

		var errMsgs []string
		var errorResponse map[string]any
		if jsonErr := json.Unmarshal(body, &errorResponse); jsonErr == nil {
			if reason, ok := errorResponse["reason"].(string); ok {
				errMsgs = append(errMsgs, reason)
			}
			if message, ok := errorResponse["message"].(string); ok {
				errMsgs = append(errMsgs, message)
			}
		}

		if len(errMsgs) > 0 {
			return nil, 0, errl.Errorf("remote server returned status %d: %s", resp.StatusCode, strings.Join(errMsgs, ", "))
		}

		return nil, 0, errl.Errorf("remote server returned status %d: %s", resp.StatusCode, string(body))
	}

	var objects []repository.TMFObjectMap
	if err := json.Unmarshal(body, &objects); err != nil {
		return nil, 0, errl.Errorf("remote server returned invalid JSON: %w", err)
	}

	if totalHeader := resp.Header.Get("X-Total-Count"); totalHeader != "" {
		if val, err := strconv.Atoi(totalHeader); err == nil {
			remoteTotalObjects = val
		} else {
			slog.Warn("Failed to parse X-Total-Count header", "value", totalHeader, "error", err)
		}
	}

	// Process each object with the user-supplied logic
	var cont bool
	if processObject != nil {
		for i := range objects {
			objects[i], cont, err = processObject(objects[i])
			// In case of error we just log it and continue with the next object
			if err != nil {
				err = errl.Error(err)
				slog.Error("processing object", "object_id", objects[i].ID(), "error", err)
			}
			// If the user wants to stop processing, we return the objects retrieved so far
			if !cont {
				return objects, remoteTotalObjects, nil
			}
		}
	}

	return objects, remoteTotalObjects, nil
}

// Get sends a GET request to the remote server.
func (c *TMFClient) Get(ctx context.Context, path string, headers map[string]string) (*http.Response, []byte, error) {
	return c.do(ctx, "GET", path, nil, headers)
}

// Post sends a POST request to the remote server.
func (c *TMFClient) Post(ctx context.Context, path string, body []byte, headers map[string]string) (*http.Response, []byte, error) {
	return c.do(ctx, "POST", path, body, headers)
}

// Put sends a PUT request to the remote server.
func (c *TMFClient) Put(ctx context.Context, path string, body []byte, headers map[string]string) (*http.Response, []byte, error) {
	return c.do(ctx, "PUT", path, body, headers)
}

// Patch sends a PATCH request to the remote server.
func (c *TMFClient) Patch(ctx context.Context, path string, body []byte, headers map[string]string) (*http.Response, []byte, error) {
	return c.do(ctx, "PATCH", path, body, headers)
}

// Delete sends a DELETE request to the remote server.
func (c *TMFClient) Delete(ctx context.Context, path string, headers map[string]string) (*http.Response, []byte, error) {
	return c.do(ctx, "DELETE", path, nil, headers)
}

// do sends an HTTP request to the remote server.
// It uses the BaseURL and PathPrefix for the server from the configuration.
func (c *TMFClient) do(ctx context.Context, method, path string, body []byte, headers map[string]string) (*http.Response, []byte, error) {

	var url string

	if c.config.InternalMode {
		// Get the URL from the config
		origin, err := config.InternalUpstreamURL(path)
		if err != nil {
			return nil, nil, errl.Errorf("failed to get upstream URL for %s: %w", path, err)
		}
		url = fmt.Sprintf("%s%s", origin, path)
	} else {
		url = fmt.Sprintf("%s%s", c.config.BaseURL, path)
	}

	var req *http.Request
	var err error

	if body != nil {
		req, err = http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
	} else {
		req, err = http.NewRequestWithContext(ctx, method, url, nil)
	}

	if err != nil {
		return nil, nil, errl.Errorf("failed to create request for %s: %w", url, err)
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, errl.Errorf("error sending %s request to %s: %w", method, url, err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, errl.Errorf("failed to read response body: %w", err)
	}

	return resp, responseBody, nil
}
