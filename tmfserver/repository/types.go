package repository

import (
	"encoding/json"
	"strings"
	"time"
)

// ValidationResult represents the result of validating a single object
type ValidationResult struct {
	ObjectID      string              `json:"object_id"`
	Href          string              `json:"href"`
	ObjectType    string              `json:"object_type"`
	Valid         bool                `json:"-"`
	Errors        []ValidationError   `json:"errors,omitempty"`
	Warnings      []ValidationWarning `json:"-"`
	ErrorsFixed   []ValidationError   `json:"-"`
	WarningsFixed []ValidationWarning `json:"-"`
	Timestamp     time.Time           `json:"-"`
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// ValidationWarning represents a validation warning
type ValidationWarning struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (vr *ValidationResult) String() string {
	var buf strings.Builder
	buf.WriteString("ValidationError{")
	buf.WriteString("ObjectID=")
	buf.WriteString(vr.ObjectID)

	out, _ := json.Marshal(vr.Errors)

	buf.WriteString(string(out))
	buf.WriteString("}")
	return buf.String()
}
