package types

import (
	"net/http"
	"net/url"
	"strings"
)

// Request represents a generic HTTP request. Handlers must convert to this representation.
// In this way, we support easily any HTTP framework (currently Fiber), but also other
// future channels like JSON-RPC or even non-HTTP channels like GRPC.
type Request struct {
	Method        string     // The HTTP method (GET, POST, PUT, PATCH, DELETE)
	Action        HttpAction // The action to perform (READ, CREATE, PUT, UPDATE, DELETE, LIST)
	APIfamily     string     // The API family (e.g., "productCatalogManagement")
	APIVersion    string     // The API version (e.g., "v4", "v5")
	ResourceName  string     // The resource name (e.g., "productOffering", "catalog")
	ID            string     // The ID of the resource (empty for CREATE requests)
	QueryParams   url.Values // The query parameters (e.g., "limit=1", "offset=0")
	Body          []byte     // The body of the request (empty for READ, LIST, DELETE)
	AuthUser      AuthUser   // The authenticated user, or zero value if non authenticated
	HealthRequest bool       // Whether this is a health request
}

func (r *Request) ToMap() map[string]any {
	return map[string]any{
		"method":   r.Method,
		"action":   r.Action,
		"api":      r.APIfamily,
		"version":  r.APIVersion,
		"resource": r.ResourceName,
		"id":       r.ID,
	}
}

type HttpAction string

// HttpActionFromMethod converts an HTTP request to an HttpAction
//
//	@param httpMethod the HTTP method (e.g., "GET", "POST", "PUT", "PATCH", "DELETE")
//	@param idParam the ID of the resource (empty for CREATE requests)
func HttpActionFromMethod(httpMethod string, idParam string) HttpAction {
	httpMethod = strings.ToUpper(httpMethod)
	action := HttpActions[httpMethod]
	if idParam == "" && httpMethod == http.MethodGet {
		action = ActionLIST
	}
	return action
}

// These are the possible values for Action
const (
	ActionREAD    HttpAction = "READ"
	ActionCREATE  HttpAction = "CREATE"
	ActionREPLACE HttpAction = "REPLACE"
	ActionUPDATE  HttpAction = "UPDATE"
	ActionDELETE  HttpAction = "DELETE"
	ActionLIST    HttpAction = "LIST"
)

// These are more friendly names for the writers of policy rules and can be used interchangeably
var HttpActions = map[string]HttpAction{
	"GET":    ActionREAD,
	"POST":   ActionCREATE,
	"PUT":    ActionREPLACE,
	"PATCH":  ActionUPDATE,
	"DELETE": ActionDELETE,
	"LIST":   ActionLIST,
}
