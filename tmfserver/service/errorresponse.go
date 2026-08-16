package service

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/hesusruiz/tmforum/internal/errl"
)

// ApiError represents a standardized TMForum API error response.
type ApiError struct {
	statusCode     int    `json:"-"`
	Type           string `json:"@type,omitempty"`
	Code           string `json:"code"`
	Reason         string `json:"reason"`
	Message        string `json:"message,omitempty"`
	Status         string `json:"status,omitempty"`
	ReferenceError string `json:"referenceError,omitempty"`
}

// Implement the error interface
func (e *ApiError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// StatusCode returns the HTTP status code associated with the error.
func (e *ApiError) StatusCode() int {
	return e.statusCode
}

// NewApiError creates a new ApiError instance.
func NewApiError(statusCode int, code, reason, message, status, referenceError string) *ApiError {
	return &ApiError{
		statusCode:     statusCode,
		Code:           code,
		Reason:         reason,
		Message:        message,
		Status:         status,
		ReferenceError: referenceError,
		Type:           "Error", // Default @type for Error object
	}
}

// ErrorResponse creates a standardized error response using only the HTTP status code
func ErrorResponse(err error, statusCode int) *Response {
	_, code, reason := getHTTPStatusInfo(statusCode)

	message := ""
	if err != nil {
		wrappedErr := errl.Error2(err)
		message = fmt.Sprintf("%s %s: %s", code, reason, wrappedErr.Error())
	} else {
		message = fmt.Sprintf("%s %s", code, reason)
	}

	if statusCode >= 500 {
		slog.Error("server error", slog.String("error", message))
	} else {
		slog.Warn("client error", slog.String("error", message))
	}

	apiErr := NewApiError(statusCode, code, reason, message, "", "")
	return &Response{StatusCode: statusCode, Body: apiErr}
}

// ErrorResponsef creates a standardized error response with formatted error message using only the HTTP status code
func ErrorResponsef(statusCode int, format string, args ...any) *Response {
	_, code, reason := getHTTPStatusInfo(statusCode)

	wrappedErr := errl.Errorf2(format, args...)
	message := fmt.Sprintf("%s %s: %s", code, reason, wrappedErr.Error())

	if statusCode >= 500 {
		slog.Error("server error", slog.String("error", message))
	} else {
		slog.Warn("client error", slog.String("error", message))
	}

	apiErr := NewApiError(statusCode, code, reason, message, "", "")
	return &Response{StatusCode: statusCode, Body: apiErr}
}

// getHTTPStatusInfo returns the standard HTTP status code, code string, and status message
func getHTTPStatusInfo(statusCode int) (int, string, string) {
	switch statusCode {
	case http.StatusBadRequest:
		return statusCode, "400", "Bad Request"
	case http.StatusUnauthorized:
		return statusCode, "401", "Unauthorized"
	case http.StatusForbidden:
		return statusCode, "403", "Forbidden"
	case http.StatusNotFound:
		return statusCode, "404", "Not Found"
	case http.StatusMethodNotAllowed:
		return statusCode, "405", "Method Not Allowed"
	case http.StatusConflict:
		return statusCode, "409", "Conflict"
	case http.StatusInternalServerError:
		return statusCode, "500", "Internal Server Error"
	case http.StatusNotImplemented:
		return statusCode, "501", "Not Implemented"
	case http.StatusBadGateway:
		return statusCode, "502", "Bad Gateway"
	case http.StatusServiceUnavailable:
		return statusCode, "503", "Service Unavailable"
	default:
		// For any other status code, use the numeric value as both code and status
		return statusCode, fmt.Sprintf("%d", statusCode), http.StatusText(statusCode)
	}
}
