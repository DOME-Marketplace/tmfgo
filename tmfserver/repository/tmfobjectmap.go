package repository

import (
	"encoding/json"
	"fmt"
	"maps"
	"slices"
	"strings"
	"time"

	"github.com/hesusruiz/tmforum/internal/errl"
	"github.com/hesusruiz/tmforum/internal/jpath"
	"github.com/hesusruiz/tmforum/internal/jsone"
	"github.com/hesusruiz/tmforum/types"
	"golang.org/x/exp/slog"
)

// TMFObjectMap represents a TMForum object as a map with utility methods
// This is designed to make fixing validation errors simple and efficient
// Using a map makes it easier to support the many types of different objects of TM Forum,
// and be more robust to slight variations in the JSON data received from the TM Forum server.
// But to make manipulations as type-safe as possible, the map is surrounded by a set of methods.
type TMFObjectMap map[string]any

// NewTMFObjectMap creates a new TMFObjectMap from a JSON byte slice.
// This is intended to be used to instantiate a TMFObjectMap from the content of the database record,
// and does not perform any validations, so any JSON object will be accepted.
func NewTMFObjectMap(data []byte) (TMFObjectMap, error) {
	var obj TMFObjectMap
	err := jsone.Unmarshal(data, &obj)
	if err != nil {
		return nil, errl.Errorf("failed to unmarshal TMF object: %w", err)
	}
	return obj, nil
}

// NewTMFObjectMapFromBytes creates a new TMFObject from a JSON byte slice.
// It is intended to be used with data received from a remote TMF server.
// The `@type` field is checked to ensure it matches the resourceName passed by the caller.
// If the `@type` field is not present, it is added to the object.
func NewTMFObjectMapFromBytes(resourceName string, data []byte) (TMFObjectMap, error) {
	var obj TMFObjectMap
	err := jsone.Unmarshal(data, &obj)
	if err != nil {
		return nil, errl.Errorf("failed to unmarshal TMF object: %w", err)
	}

	// Check that the object matches the resourceName
	objType, _ := obj["@type"].(string)
	if objType != "" {
		// If the object type is specified, it must match the resourceName passed by the caller
		if !strings.EqualFold(resourceName, objType) {
			return nil, errl.Errorf("invalid object type, expected %s but received %s", resourceName, objType)
		}

	} else {
		// No object type in the object, set it to the resourceName
		obj["@type"] = resourceName
	}

	return obj, nil

}

// Validate checks if the object is valid, and returns as many errors as it can find.
// It does not stop after finding the first error, and results are accumulated in the ValidationResult.
func (obj TMFObjectMap) Validate(resourceName string) ValidationResult {
	result := ValidationResult{
		ObjectID:   obj.ID(),
		ObjectType: resourceName,
		Valid:      true,
		Timestamp:  time.Now(),
	}

	// Validate required fields
	obj.validateRequiredFields(resourceName, &result)

	// Validate related party requirements
	obj.validateRelatedParty(&result)

	// Determine overall validity (object is valid if it has zero validation errors)
	result.Valid = len(result.Errors) == 0

	return result
}

func (obj TMFObjectMap) ValidateCreate(resourceName string) ValidationResult {
	result := ValidationResult{
		ObjectID:   obj.ID(),
		ObjectType: resourceName,
		Valid:      true,
		Timestamp:  time.Now(),
	}

	// Validate required fields
	obj.validateRequiredFieldsCreate(resourceName, &result)

	// Validate related party requirements
	obj.validateRelatedParty(&result)

	// Determine overall validity (object is valid if it has zero validation errors)
	result.Valid = len(result.Errors) == 0

	return result
}

// validateRequiredFields checks if all required fields are present and optionally fixes them
func (obj TMFObjectMap) validateRequiredFieldsCreate(resourceName string, result *ValidationResult) {

	// Special processing for the object type: check that the object matches the resourceName and fix the object if needed.
	objType := obj.Type()
	if objType != "" {
		// If the object type is specified, it must match the resourceName passed by the caller
		field := "@type"
		if !strings.EqualFold(resourceName, objType) {
			result.Errors = append(result.Errors, ValidationError{
				Field:   field,
				Message: fmt.Sprintf("Object type field '%s' does not match resource type '%s'", objType, resourceName),
				Code:    "MISSING_INVALID_VALUE",
			})
		}

	} else {
		// No object type in the object, set it to the resourceName
		field := "@type"
		result.Warnings = append(result.Warnings, ValidationWarning{
			Field:   field,
			Message: fmt.Sprintf("Recommended field '%s' is missing, setting it to %s", field, resourceName),
			Code:    "MISSING_RECOMMENDED_FIELD",
		})
		obj.SetType(resourceName)
	}

	// Generate the warnings for the recommended fields
	for _, field := range types.RecommendedFieldsForAllObjects {
		if !obj.HasField(field) {
			result.Warnings = append(result.Warnings, ValidationWarning{
				Field:   field,
				Message: fmt.Sprintf("Recommended field '%s' is missing", field),
				Code:    "MISSING_RECOMMENDED_FIELD",
			})
		}
	}

}

func (obj TMFObjectMap) validateRequiredFields(resourceName string, result *ValidationResult) {

	// Special processiong for the object type: check that the object matches the resourceName and fix the object if needed.
	objType := obj.Type()
	if objType != "" {
		// If the object type is specified, it must match the resourceName passed by the caller
		field := "@type"
		if !strings.EqualFold(resourceName, objType) {
			result.Errors = append(result.Errors, ValidationError{
				Field:   field,
				Message: fmt.Sprintf("Object type field '%s' does not match resource type '%s'", objType, resourceName),
				Code:    "MISSING_INVALID_VALUE",
			})
		}

	} else {
		// No object type in the object, set it to the resourceName
		field := "@type"
		result.Warnings = append(result.Warnings, ValidationWarning{
			Field:   field,
			Message: fmt.Sprintf("Recommended field '%s' is missing, setting it to %s", field, resourceName),
			Code:    "MISSING_RECOMMENDED_FIELD",
		})
		obj.SetType(resourceName)
	}

	// This checks the fields that are required for all objects
	for _, field := range types.RequiredFieldsForAllObjects {
		if !obj.HasField(field) {
			result.Errors = append(result.Errors, ValidationError{
				Field:   field,
				Message: fmt.Sprintf("Required field '%s' is missing", field),
				Code:    "MISSING_REQUIRED_FIELD",
			})
		}
	}

	// Generate the warnings for the recommended fields
	for _, field := range types.RecommendedFieldsForAllObjects {
		if !obj.HasField(field) {
			result.Warnings = append(result.Warnings, ValidationWarning{
				Field:   field,
				Message: fmt.Sprintf("Recommended field '%s' is missing", field),
				Code:    "MISSING_RECOMMENDED_FIELD",
			})
		}
	}

}

func (obj TMFObjectMap) validateRelatedParty(result *ValidationResult) {

	// Return if the object does not require Seller nor Buyer info
	// It is enough to check Seller info, as it is impossible to have Buyer info without it
	if !obj.RequiresSellerInfo("") {
		return
	}

	// Check that the object has a relatedParty object
	relatedParties := jpath.GetList(obj, "relatedParty")
	if len(relatedParties) == 0 {
		// msg := "Missing " + userRole + " and " + userOperatorRole + " fields"
		result.Errors = append(result.Errors, ValidationError{
			Field:   "relatedParty",
			Message: "Missing 'relatedParty' object",
			Code:    "MISSING_RELATED_PARTY_INFO",
		})

		return
	}

	var sellerDid string
	var sellerOperatorDid string
	var buyerDid string
	var buyerOperatorDid string

	sellerRole := strings.ToLower("Seller")
	sellerOperatorRole := strings.ToLower("SellerOperator")
	buyerRole := strings.ToLower("Buyer")
	buyerOperatorRole := strings.ToLower("BuyerOperator")

	sellerRoleCount := 0
	sellerOperatorRoleCount := 0
	buyerRoleCount := 0
	buyerOperatorRoleCount := 0

	for _, rp := range relatedParties {
		// Cast the entry to a map[string]any
		rpMap, _ := rp.(map[string]any)
		if len(rpMap) == 0 {
			result.Errors = append(result.Errors, ValidationError{
				Field:   "relatedParty",
				Message: "Invalid 'relatedParty' object",
				Code:    "INVALID_RELATED_PARTY_INFO",
			})
			return
		}

		// Extract the "role" field in lowercase for case-insentitive comparison
		rpRole, _ := rpMap["role"].(string)
		rpRole = strings.ToLower(rpRole)

		// If the role of the entry is not one that we are looking for, continue the loop
		if rpRole != sellerRole && rpRole != sellerOperatorRole && rpRole != buyerRole && rpRole != buyerOperatorRole {
			continue
		}

		switch rpRole {
		case sellerRole:
			sellerRoleCount++
			sellerDid, _ = rpMap["name"].(string)
			if sellerDid == "" {
				result.Errors = append(result.Errors, ValidationError{
					Field:   "relatedParty",
					Message: "Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role",
					Code:    "MISSING_RELATED_PARTY_INFO",
				})
			}
		case sellerOperatorRole:
			sellerOperatorRoleCount++
			sellerOperatorDid, _ = rpMap["name"].(string)
			if sellerOperatorDid == "" {
				result.Errors = append(result.Errors, ValidationError{
					Field:   "relatedParty",
					Message: "Missing or invalid 'name' field in 'relatedParty' object for 'SellerOperator' role",
					Code:    "MISSING_RELATED_PARTY_INFO",
				})
			}
		case buyerRole:
			buyerRoleCount++
			buyerDid, _ = rpMap["name"].(string)
			if buyerDid == "" {
				result.Errors = append(result.Errors, ValidationError{
					Field:   "relatedParty",
					Message: "Missing or invalid 'name' field in 'relatedParty' object for 'Buyer' role",
					Code:    "MISSING_RELATED_PARTY_INFO",
				})
			}
		case buyerOperatorRole:
			buyerOperatorRoleCount++
			buyerOperatorDid, _ = rpMap["name"].(string)
			if buyerOperatorDid == "" {
				result.Errors = append(result.Errors, ValidationError{
					Field:   "relatedParty",
					Message: "Missing or invalid 'name' field in 'relatedParty' object for 'BuyerOperator' role",
					Code:    "MISSING_RELATED_PARTY_INFO",
				})
			}
		}

	}

	// Set the error depending on what we have found for Seller and SellerOperator
	if sellerDid == "" && sellerOperatorDid == "" {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "relatedParty",
			Message: "Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role",
			Code:    "MISSING_RELATED_PARTY_INFO",
		})
	} else if sellerDid == "" {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "relatedParty",
			Message: "Missing 'relatedParty' info for 'Seller' role",
			Code:    "MISSING_RELATED_PARTY_INFO",
		})
	} else if sellerOperatorDid == "" {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "relatedParty",
			Message: "Missing 'relatedParty' info for 'SellerOperator' role",
			Code:    "MISSING_RELATED_PARTY_INFO",
		})
	}

	// Check that we have the correct number of Seller and SellerOperator entries
	if sellerRoleCount > 1 {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "relatedParty",
			Message: "Too many 'relatedParty' entries for 'Seller' role",
			Code:    "TOO_MANY_RELATED_PARTY_INFO",
		})
	}
	if sellerOperatorRoleCount > 1 {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "relatedParty",
			Message: "Too many 'relatedParty' entries for 'SellerOperator' role",
			Code:    "TOO_MANY_RELATED_PARTY_INFO",
		})
	}

	if !obj.RequiresBuyerInfo() {
		return
	}

	// Set the error depending on what we have found for Buyer and BuyerOperator
	if buyerDid == "" && buyerOperatorDid == "" {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "relatedParty",
			Message: "Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role",
			Code:    "MISSING_RELATED_PARTY_INFO",
		})
	} else if buyerDid == "" {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "relatedParty",
			Message: "Missing 'relatedParty' info for 'Buyer' role",
			Code:    "MISSING_RELATED_PARTY_INFO",
		})
	} else if buyerOperatorDid == "" {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "relatedParty",
			Message: "Missing 'relatedParty' info for 'BuyerOperator' role",
			Code:    "MISSING_RELATED_PARTY_INFO",
		})
	}

	// Check that we have the correct number of Buyer and BuyerOperator entries
	if buyerRoleCount > 1 {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "relatedParty",
			Message: "Too many 'relatedParty' entries for 'Buyer' role",
			Code:    "TOO_MANY_RELATED_PARTY_INFO",
		})
	}
	if buyerOperatorRoleCount > 1 {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "relatedParty",
			Message: "Too many 'relatedParty' entries for 'BuyerOperator' role",
			Code:    "TOO_MANY_RELATED_PARTY_INFO",
		})
	}

}

// ToTMFRecord converts the object to its storage representation to save it in the local database
// It gets some keys to make efficient SQL queries, and the object is stored as JSON
func (obj TMFObjectMap) ToTMFRecord(resourceName string) *TMFRecord {

	id := obj.ID()
	objectType := obj.Type()
	if objectType == "" {
		objectType = resourceName
	}
	version := obj.Version()
	// TODO: support for v5 API
	apiVersion := "v4"
	lastUpdate := obj.LastUpdate()
	content := obj.ToJSONSimple()

	seller, sellerOperator, _ := obj.GetSellerInfo("v4")
	buyer, buyerOperator, _ := obj.GetBuyerInfo("v4")

	now := time.Now().Unix()

	o := &TMFRecord{
		ID:             id,
		Type:           objectType,
		Version:        version,
		APIVersion:     apiVersion,
		Seller:         seller,
		SellerOperator: sellerOperator,
		Buyer:          buyer,
		BuyerOperator:  buyerOperator,
		LastUpdate:     lastUpdate,
		Content:        content,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	return o
}

// ToJSON converts the TMFObject to JSON bytes
func (obj TMFObjectMap) ToJSON() ([]byte, error) {
	return json.Marshal(obj)
}

// ToJSONSimple converts the TMFObject to JSON bytes without error checking
// This is used when we know the object is valid
func (obj TMFObjectMap) ToJSONSimple() []byte {
	data, _ := obj.ToJSON()
	return data
}

// ToMap converts the TMFObject to a regular map[string]any
func (obj TMFObjectMap) ToMap() map[string]any {
	return maps.Clone(obj)
}

// Utility methods for well-known top-level attributes

func (obj TMFObjectMap) IsPotentiallyPublic() bool {
	resourceDefinition := types.GetResourceDefinition(obj.Type())
	if resourceDefinition == nil {
		return false
	}
	return resourceDefinition.Public
}

// ID returns the object ID
func (obj TMFObjectMap) ID() string {
	if id, ok := obj["id"].(string); ok {
		return id
	}
	return ""
}

func (obj TMFObjectMap) IsIndividual() bool {
	return strings.EqualFold(obj.Type(), "Individual")
}

func (obj TMFObjectMap) IsOrganization() bool {
	return strings.EqualFold(obj.Type(), "Organization")
}

// SetID sets the object ID
func (obj TMFObjectMap) SetID(id string) {
	obj["id"] = id
}

// Href returns the object href
func (obj TMFObjectMap) Href() string {
	if href, ok := obj["href"].(string); ok {
		return href
	}
	return ""
}

// SetHref sets the object href
func (obj TMFObjectMap) SetHref(href string) {
	obj["href"] = href
}

// Version returns the object version
func (obj TMFObjectMap) Version() string {
	if version, ok := obj["version"].(string); ok {
		return version
	}
	return ""
}

// SetVersion sets the object version
func (obj TMFObjectMap) SetVersion(version string) {
	obj["version"] = version
}

// LastUpdate returns the object lastUpdate
func (obj TMFObjectMap) LastUpdate() string {
	if lastUpdate, ok := obj["lastUpdate"].(string); ok {
		return lastUpdate
	}
	return ""
}

// SetLastUpdate sets the object lastUpdate
func (obj TMFObjectMap) SetLastUpdate(lastUpdate string) {
	obj["lastUpdate"] = lastUpdate
}

// SetLastUpdateNow sets the object lastUpdate to current timestamp in RFC3339 format
func (obj TMFObjectMap) SetLastUpdateNow() {
	obj["lastUpdate"] = time.Now().Format(time.RFC3339)
}

// LastModified returns the object lastModified
func (obj TMFObjectMap) LastModified() string {
	if lastModified, ok := obj["lastModified"].(string); ok {
		return lastModified
	}
	return ""
}

// SetLastModified sets the object lastModified
func (obj TMFObjectMap) SetLastModified(lastModified string) {
	obj["lastModified"] = lastModified
}

// SetLastModifiedNow sets the object lastModified to current timestamp in RFC3339 format
func (obj TMFObjectMap) SetLastModifiedNow() {
	obj["lastModified"] = time.Now().Format(time.RFC3339)
}

// Type returns the object @type
func (obj TMFObjectMap) Type() string {
	if objType, ok := obj["@type"].(string); ok {
		return objType
	}
	return ""
}

// SetType sets the object @type
func (obj TMFObjectMap) SetType(objType string) {
	obj["@type"] = objType
}

// SchemaLocation returns the object schemaLocation
func (obj TMFObjectMap) SchemaLocation() string {
	schemaLocation, _ := obj["@schemaLocation"].(string)
	return schemaLocation
}

// SetSchemaLocation sets the object schemaLocation
func (obj TMFObjectMap) SetSchemaLocation(schemaLocation string) {
	obj["@schemaLocation"] = schemaLocation
}

const schemaBothLastUpdateAndRelatedParty = "https://raw.githubusercontent.com/DOME-Marketplace/tmf-api/refs/heads/main/DOME/TrackedShareableEntity.schema.json"
const schemaRelatedParty = "https://raw.githubusercontent.com/DOME-Marketplace/tmf-api/refs/heads/main/DOME/ShareableEntity.schema.json"
const schemaLastUpdate = "https://raw.githubusercontent.com/DOME-Marketplace/tmf-api/refs/heads/main/DOME/TrackedEntity.schema.json"

// SetDefaultSchemaLocation sets the object schemaLocation to its default schema location for the given action
func (obj TMFObjectMap) SetDefaultSchemaLocation(action *types.Action) {

	needsLastUpdate := !action.HasField("lastUpdate")
	needsRelatedParty := !action.HasField("relatedParty")

	if needsLastUpdate && needsRelatedParty {
		obj.SetSchemaLocation(schemaBothLastUpdateAndRelatedParty)
	} else if needsLastUpdate && !needsRelatedParty {
		obj.SetSchemaLocation(schemaLastUpdate)
	} else if !needsLastUpdate && needsRelatedParty {
		obj.SetSchemaLocation(schemaRelatedParty)
	}
}

// LifecycleStatus returns the object lifecycleStatus
func (obj TMFObjectMap) LifecycleStatus() string {
	if lifecycleStatus, ok := obj["lifecycleStatus"].(string); ok {
		return lifecycleStatus
	}
	return ""
}

// SetLifecycleStatus sets the object lifecycleStatus
func (obj TMFObjectMap) SetLifecycleStatus(lifecycleStatus string) {
	obj["lifecycleStatus"] = lifecycleStatus
}

// Name returns the object name
func (obj TMFObjectMap) Name() string {
	if name, ok := obj["name"].(string); ok {
		return name
	}
	return ""
}

// SetName sets the object name
func (obj TMFObjectMap) SetName(name string) {
	obj["name"] = name
}

// Description returns the object description
func (obj TMFObjectMap) Description() string {
	if description, ok := obj["description"].(string); ok {
		return description
	}
	return ""
}

// SetDescription sets the object description
func (obj TMFObjectMap) SetDescription(description string) {
	obj["description"] = description
}

// RelatedParty methods

// RelatedParty returns the relatedParty array
func (obj TMFObjectMap) RelatedParty() []map[string]any {
	if relatedParty, ok := obj["relatedParty"].([]any); ok {
		result := make([]map[string]any, 0, len(relatedParty))
		for _, item := range relatedParty {
			if rp, ok := item.(map[string]any); ok {
				result = append(result, rp)
			}
		}
		return result
	}
	return nil
}

// SetRelatedParty sets the relatedParty array
func (obj TMFObjectMap) SetRelatedParty(relatedParty []map[string]any) {
	// Convert []map[string]any to []any for JSON serialization
	items := make([]any, len(relatedParty))
	for i, rp := range relatedParty {
		items[i] = rp
	}
	obj["relatedParty"] = items
}

// AddRelatedParty adds a related party entry
func (obj TMFObjectMap) AddRelatedParty(relatedParty map[string]any) {
	current := obj.RelatedParty()
	if current == nil {
		current = make([]map[string]any, 0)
	}
	current = append(current, relatedParty)
	obj.SetRelatedParty(current)
}

// HasRelatedParty returns true if the object has related party information
func (obj TMFObjectMap) HasRelatedParty() bool {
	relatedParty := obj.RelatedParty()
	return len(relatedParty) > 0
}

// Validation helper methods

// HasField checks if a specific field exists and is not empty
func (obj TMFObjectMap) HasField(field string) bool {
	value, exists := obj[field]
	if !exists {
		return false
	}

	switch v := value.(type) {
	case string:
		return v != ""
	case []any:
		return len(v) > 0
	case map[string]any:
		return len(v) > 0
	default:
		return true // Other types are considered present if they exist
	}
}

// SetField sets a field value
func (obj TMFObjectMap) SetField(field string, value any) {
	obj[field] = value
}

// GetField returns a field value
func (obj TMFObjectMap) GetField(field string) any {
	return obj[field]
}

// GetStringField returns a field value as string
func (obj TMFObjectMap) GetStringField(field string) string {
	if value, ok := obj[field].(string); ok {
		return value
	}
	return ""
}

// SetStringField sets a field value as string
func (obj TMFObjectMap) SetStringField(field string, value string) {
	obj[field] = value
}

// GetArrayField returns a field value as []any
func (obj TMFObjectMap) GetArrayField(field string) []any {
	if value, ok := obj[field].([]any); ok {
		return value
	}
	return nil
}

// SetArrayField sets a field value as []any
func (obj TMFObjectMap) SetArrayField(field string, value []any) {
	obj[field] = value
}

// GetMapField returns a field value as map[string]any
func (obj TMFObjectMap) GetMapField(field string) map[string]any {
	if value, ok := obj[field].(map[string]any); ok {
		return value
	}
	return nil
}

// SetMapField sets a field value as map[string]any
func (obj TMFObjectMap) SetMapField(field string, value map[string]any) {
	obj[field] = value
}

// IsEmpty returns true if the object is empty
func (obj TMFObjectMap) IsEmpty() bool {
	return len(obj) == 0
}

// Keys returns all the keys in the object
func (obj TMFObjectMap) Keys() []string {
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	return keys
}

// DeleteField removes a field from the object
func (obj TMFObjectMap) DeleteField(field string) {
	delete(obj, field)
}

// String returns a string representation of the object (for debugging)
func (obj TMFObjectMap) String() string {
	data, err := obj.ToJSON()
	if err != nil {
		return fmt.Sprintf("TMFObject{error: %v}", err)
	}
	return string(data)
}

// setSellerAndBuyerInfo adds the required fields to the incoming object argument
// It calls the appropriate version-specific function based on the API version
func (obj TMFObjectMap) SetSellerInfo(serverOperatorDid string, organizationIdentifier string, apiVersion string) (err error) {
	// We do nothing for Individual or Organization objects, which are special and do not have Seller info
	objType := obj.Type()
	objType = strings.ToLower(objType)
	if objType == "individual" || objType == "organization" {
		return nil
	}

	switch apiVersion {
	case "v4":
		return setSellerInfoV4(obj, serverOperatorDid, organizationIdentifier)
	case "v5":
		return setSellerInfoV5(obj, serverOperatorDid, organizationIdentifier)
	default:
		// Default to V5 for backward compatibility
		return setSellerInfoV4(obj, serverOperatorDid, organizationIdentifier)
	}
}

// setSellerInfoV4 adds the required fields to the incoming object argument for V4 API
// Specifically, the Seller and SellerOperator roles are added to the relatedParty list
func setSellerInfoV4(tmfObjectMap map[string]any, serverOperatorDid string, organizationIdentifier string) (err error) {

	// Normalize the organization identifier to the DID format
	if !strings.HasPrefix(organizationIdentifier, "did:elsi:") {
		organizationIdentifier = "did:elsi:" + organizationIdentifier
	}

	if !strings.HasPrefix(serverOperatorDid, "did:elsi:") {
		serverOperatorDid = "did:elsi:" + serverOperatorDid
	}

	// Look for the "Seller", "SellerOperator", "Buyer" and "BuyerOperator" roles
	relatedParties := jpath.GetList(tmfObjectMap, "relatedParty")

	// Build the two entries for V4 format
	sellerEntry := map[string]any{
		"role":          "Seller",
		"id":            "urn:ngsi-ld:organization:" + organizationIdentifier,
		"href":          "urn:ngsi-ld:organization:" + organizationIdentifier,
		"name":          organizationIdentifier,
		"@referredType": "Organization",
	}
	sellerOperator := map[string]any{
		"role":          "SellerOperator",
		"id":            "urn:ngsi-ld:organization:" + serverOperatorDid,
		"href":          "urn:ngsi-ld:organization:" + serverOperatorDid,
		"name":          serverOperatorDid,
		"@referredType": "Organization",
	}

	// If the object does not have a 'relatedParty' array, create one with Seller=caller and SelletOperator=server_operator
	if len(relatedParties) == 0 {
		slog.Debug("setSellerAndBuyerInfoV4: no relatedParty, adding seller and sellerOperator")
		tmfObjectMap["relatedParty"] = []any{sellerEntry, sellerOperator}
		return nil
	}

	// We now search for Seller and SellerOperator entries and overwrite them or add them.
	// The relatedParty array may contain entries which are not these, and we should not touch them.

	foundSeller := false
	foundSellerOperator := false

	newRelatedParties := []any{}

	for _, rp := range relatedParties {

		// Convert entry to a map
		rpMap, _ := rp.(map[string]any)
		if len(rpMap) == 0 {
			return errl.Errorf("invalid relatedParty entry")
		}

		rpRole, _ := rpMap["role"].(string)
		rpRole = strings.ToLower(rpRole)

		if rpRole != "seller" && rpRole != "selleroperator" {
			newRelatedParties = append(newRelatedParties, rp)
			// Go to next entry
			continue
		}

		if rpRole == "seller" {
			// Overwrite the entry, because we can not allow the user to create fake info
			newRelatedParties = append(newRelatedParties, sellerEntry)
			foundSeller = true
			continue
		}
		if rpRole == "selleroperator" {
			// Overwrite the entry, because we can not allow the user to create fake info
			newRelatedParties = append(newRelatedParties, sellerOperator)
			foundSellerOperator = true
			continue
		}

	}

	if !foundSeller {
		// Add the seller if it is not already in the list
		slog.Debug("setSellerAndBuyerInfoV4: adding seller", "organizationIdentifier", organizationIdentifier)
		newRelatedParties = append(newRelatedParties, sellerEntry)
	}

	if !foundSellerOperator {
		// Add the seller operator if it is not already in the list
		slog.Debug("setSellerAndBuyerInfoV4: adding seller operator", "organizationIdentifier", organizationIdentifier)
		newRelatedParties = append(newRelatedParties, sellerOperator)
	}

	tmfObjectMap["relatedParty"] = newRelatedParties

	return nil

}

// setSellerInfoV5 adds the required fields to the incoming object argument
// Specifically, the Seller and SellerOperator roles are added to the relatedParty list
func setSellerInfoV5(tmfObjectMap map[string]any, serverOperatorDid string, organizationIdentifier string) (err error) {

	// Normalize all organization identifiers to the DID format
	if !strings.HasPrefix(organizationIdentifier, "did:elsi:") {
		organizationIdentifier = "did:elsi:" + organizationIdentifier
	}

	// Look for the "Seller", "SellerOperator", "Buyer" and "BuyerOperator" roles
	relatedParties := jpath.GetList(tmfObjectMap, "relatedParty")

	// Build the two entries
	sellerEntry := map[string]any{
		"role":  "Seller",
		"@type": "RelatedPartyRefOrPartyRoleRef",
		"partyOrPartyRole": map[string]any{
			"@type":         "PartyRef",
			"href":          "urn:ngsi-ld:organization:" + organizationIdentifier,
			"id":            "urn:ngsi-ld:organization:" + organizationIdentifier,
			"name":          organizationIdentifier,
			"@referredType": "Organization",
		},
	}
	sellerOperator := map[string]any{
		"role":  "SellerOperator",
		"@type": "RelatedPartyRefOrPartyRoleRef",
		"partyOrPartyRole": map[string]any{
			"@type":         "PartyRef",
			"href":          "urn:ngsi-ld:organization:" + serverOperatorDid,
			"id":            "urn:ngsi-ld:organization:" + serverOperatorDid,
			"name":          serverOperatorDid,
			"@referredType": "Organization",
		},
	}

	if len(relatedParties) == 0 {
		slog.Debug("setSellerAndBuyerInfo: no relatedParty, adding seller and sellerOperator")
		tmfObjectMap["relatedParty"] = []any{sellerEntry, sellerOperator}
		return nil
	}

	foundSeller := false
	foundSellerOperator := false

	newRelatedParties := []any{}

	for _, rp := range relatedParties {

		// Convert entry to a map
		rpMap, _ := rp.(map[string]any)
		if len(rpMap) == 0 {
			return errl.Errorf("invalid relatedParty entry")
		}

		rpRole, _ := rpMap["role"].(string)
		rpRole = strings.ToLower(rpRole)

		if rpRole != "seller" && rpRole != "selleroperator" {
			newRelatedParties = append(newRelatedParties, rp)
			// Go to next entry
			continue
		}

		if rpRole == "seller" {
			// Overwrite the entry, because we can not allow the user to create fake info
			newRelatedParties = append(newRelatedParties, sellerEntry)
			foundSeller = true
			continue
		}
		if rpRole == "selleroperator" {
			// Overwrite the entry, because we can not allow the user to create fake info
			newRelatedParties = append(newRelatedParties, sellerOperator)
			foundSellerOperator = true
			continue
		}

	}

	if !foundSeller {
		// Add the seller if it is not already in the list
		slog.Debug("setSellerAndBuyerInfo: adding seller", "organizationIdentifier", organizationIdentifier)
		newRelatedParties = append(newRelatedParties, sellerEntry)
	}

	if !foundSellerOperator {
		// Add the seller operator if it is not already in the list
		slog.Debug("setSellerAndBuyerInfo: adding seller operator", "organizationIdentifier", organizationIdentifier)
		newRelatedParties = append(newRelatedParties, sellerOperator)
	}

	tmfObjectMap["relatedParty"] = newRelatedParties

	return nil

}

func (obj TMFObjectMap) RequiresSellerInfo(resourceName string) bool {
	var objType string
	if len(resourceName) > 0 {
		objType = resourceName
	} else {
		objType = obj.Type()
	}
	objType = strings.ToLower(objType)
	return !slices.Contains(types.DoNotRequireRelatedParties, objType)
}

func (obj TMFObjectMap) RequiresBuyerInfo() bool {
	objType := obj.Type()
	objType = strings.ToLower(objType)
	return !slices.Contains(types.DoNotRequireBuyerInfo, objType)
}

// GetSellerInfo finds the Seller and SellerOperator identifiers in the relatedParty array of the object,
// using the apiVersion to determine which version of the function to use. If no version is provided, defaults to "v4".
// If some identifier is missing (or both), it returns an error. But it returns what it finds.
// So, even if the returned error is not nil, the caller may check the sellerDid and the sellerOperatorDid.
// This is useful if the caller has logic to handle cases where only one of the values is found.
func (obj TMFObjectMap) GetSellerInfo(apiVersion string) (sellerDid string, sellerOperatorDid string, err error) {
	if !obj.RequiresSellerInfo("") {
		return
	}

	switch apiVersion {
	case "v4":
		return getUserAndUserOperatorInfoV4(obj, "Seller", "SellerOperator")
	case "v5":
		return getUserAndUserOperatorInfoV5(obj, "Seller", "SellerOperator")
	default:
		// Default to V4 for backward compatibility
		return getUserAndUserOperatorInfoV4(obj, "Seller", "SellerOperator")
	}
}

func (obj TMFObjectMap) GetBuyerInfo(apiVersion string) (sellerDid string, sellerOperatorDid string, err error) {
	switch apiVersion {
	case "v4":
		return getUserAndUserOperatorInfoV4(obj, "Buyer", "BuyerOperator")
	case "v5":
		return getUserAndUserOperatorInfoV5(obj, "Buyer", "BuyerOperator")
	default:
		// Default to V4 for backward compatibility
		return getUserAndUserOperatorInfoV4(obj, "Buyer", "BuyerOperator")
	}
}

// getUserAndUserOperatorInfoV4 finds the relatedparty entries with the specified roles.
// It returns what it finds. So, even if the returned error is not nil, the caller may check the sellerDid and the sellerOperatorDid.
// This is useful if the caller has logic to handle cases where only one of the values is found.
func getUserAndUserOperatorInfoV4(tmfObjectMap map[string]any, userRole string, userOperatorRole string) (sellerDid string, sellerOperatorDid string, err error) {
	// In V4, relatedParty is a list of maps with fields like "role", "id", "href", "name", "@referredType"
	// We need to extract the "name" for the "Seller" and "SellerOperator" roles

	// Convert roles to lowercase to facilitate case-insensitive comparison
	userRole = strings.ToLower(userRole)
	userOperatorRole = strings.ToLower(userOperatorRole)

	relatedParties := jpath.GetList(tmfObjectMap, "relatedParty")

	if len(relatedParties) == 0 {
		err = errl.Errorf("no relatedParty")
		return "", "", err
	}

	for _, rp := range relatedParties {
		// Cast the entry to a map[string]any
		rpMap, _ := rp.(map[string]any)
		if len(rpMap) == 0 {
			return "", "", errl.Errorf("invalid relatedParty entry")
		}

		// Extract the "role" field in lowercase for case-insentitive comparison
		rpRole, _ := rpMap["role"].(string)
		rpRole = strings.ToLower(rpRole)

		// If the role of the entry is not one that we are looking for, continue the loop
		if rpRole != userRole && rpRole != userOperatorRole {
			continue
		}

		if rpRole == userRole {
			sellerDid, _ = rpMap["name"].(string)
			continue
		}
		if rpRole == userOperatorRole {
			sellerOperatorDid, _ = rpMap["name"].(string)
			continue
		}
	}

	// Return an error if one or both fields are not set, indicating the condition is not met
	if sellerDid == "" && sellerOperatorDid == "" {
		err = errl.Errorf("no %s or %s", userRole, userOperatorRole)
	} else if sellerDid == "" {
		err = errl.Errorf("no %s", userRole)
	} else if sellerOperatorDid == "" {
		err = errl.Errorf("no %s", userOperatorRole)
	}

	// return what we have foud. So, even if the err is not nil, the caller may check the sellerDid and the sellerOperatorDid.
	// This is useful if the caller has logic to handle cases where only one of the values is found.
	return sellerDid, sellerOperatorDid, err

}

func getUserAndUserOperatorInfoV5(tmfObjectMap map[string]any, userRole string, userOperatorRole string) (sellerDid string, sellerOperatorDid string, err error) {

	userRole = strings.ToLower(userRole)
	userOperatorRole = strings.ToLower(userOperatorRole)

	// Look for the "Seller", "SellerOperator", "Buyer" and "BuyerOperator" roles
	relatedParties := jpath.GetList(tmfObjectMap, "relatedParty")

	if len(relatedParties) == 0 {
		err = errl.Errorf("no relatedParty")
		return
	}

	for _, rp := range relatedParties {

		// Convert entry to a map
		rpMap, _ := rp.(map[string]any)
		if len(rpMap) == 0 {
			return "", "", errl.Errorf("invalid relatedParty entry")
		}

		rpRole, _ := rpMap["role"].(string)
		rpRole = strings.ToLower(rpRole)

		if rpRole != userRole && rpRole != userOperatorRole {
			// Go to next entry
			continue
		}

		if rpRole == userRole {
			party, _ := rpMap["partyOrPartyRole"].(map[string]any)
			sellerDid, _ = party["name"].(string)
			continue
		}
		if rpRole == userOperatorRole {
			party, _ := rpMap["partyOrPartyRole"].(map[string]any)
			sellerOperatorDid, _ = party["name"].(string)
			continue
		}

	}

	if sellerDid == "" && sellerOperatorDid == "" {
		err = errl.Errorf("no %s or %s", userRole, userOperatorRole)
		return
	}

	if sellerDid == "" {
		err = errl.Errorf("no %s", userRole)
		return
	}
	if sellerOperatorDid == "" {
		err = errl.Errorf("no %s", userOperatorRole)
		return
	}

	return

}
