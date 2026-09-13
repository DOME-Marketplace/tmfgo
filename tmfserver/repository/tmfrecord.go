package repository

import (
	"log/slog"
	"strings"
	"time"

	"github.com/DOME-Marketplace/tmfgo/internal/errl"
)

// NewTMFRecord creates a new TMFObject in storage format
func NewTMFRecord(id, objectType, version, apiVersion, lastUpdate string, content []byte) *TMFRecord {
	now := time.Now()

	o := &TMFRecord{
		ID:         id,
		Type:       objectType,
		Version:    version,
		APIVersion: apiVersion,
		LastUpdate: lastUpdate,
		Content:    content,
		CreatedAt:  now.Unix(),
		UpdatedAt:  now.Unix(),
	}

	return o

}

func (o *TMFRecord) SetBuyerID(buyerID string) {
	if !strings.HasPrefix(buyerID, "did:elsi:") {
		buyerID = "did:elsi:" + buyerID
	}
	o.Buyer = buyerID
}

func (o *TMFRecord) SetSellerID(sellerID string) {
	if !strings.HasPrefix(sellerID, "did:elsi:") {
		sellerID = "did:elsi:" + sellerID
	}
	o.Seller = sellerID
}

// GetCreatedAt returns the time.Time representation of the Unix timestamp of the internal field CreatedAt
func (o *TMFRecord) GetCreatedAt() time.Time {
	return time.Unix(o.CreatedAt, 0)
}

// GetUpdatedAt returns the time.Time representation of the Unix timestamp of the internal field UpdatedAt
func (o *TMFRecord) GetUpdatedAt() time.Time {
	return time.Unix(o.UpdatedAt, 0)
}

// Validate validates the TMFObject.
// It returns an error if the object is invalid, and detailed validations are in o.Validations
func (o *TMFRecord) Validate() error {
	// We need to unmarshal the JSON to a Map
	// Doing this we get the validations performed and the result in o.Validations
	_, err := o.ToTMFObjectMap()
	return err
}

// ToTMFObjectMap returns a TMFObjectMap reusing any previous marshalling and validation.
// To force to always marshal of o.Content and running validations, use o.MustToTMFObjectMap
// After the call, the validations results are in o.Validations
func (o *TMFRecord) ToTMFObjectMap() (TMFObjectMap, error) {
	if o == nil {
		return nil, errl.Errorf("object is nil")
	}
	if o.ContentMap != nil {
		return o.ContentMap, nil
	}

	return o.MustToTMFObjectMap()
}

// MustToTMFObjectMap returns a TMFObjectMap after running validations
func (o *TMFRecord) MustToTMFObjectMap() (TMFObjectMap, error) {

	omap, err := NewTMFObjectMap(o.Content)
	if err != nil {
		o.Validations.ObjectID = o.ID
		o.Validations.ObjectType = o.Type
		o.Validations.Valid = false
		o.Validations.Timestamp = time.Now()
		o.Validations.Errors = append(o.Validations.Errors, ValidationError{
			Field:   "object",
			Message: errl.Error(err).Error(),
			Code:    "INVALID_OBJECT",
		})

		err = errl.Errorf("failed to unmarshal object content: %w", err)
		return nil, err
	}

	o.Validations = omap.Validate(o.Type)
	if len(o.Validations.Errors) > 0 {
		slog.Error("invalid object", "id", o.ID, "type", o.Type, "errors", o.Validations.String())
		return nil, errl.Errorf("invalid object: %s", o.Validations.String())
	}

	o.ContentMap = omap

	return o.ContentMap, nil
}

func (o *TMFRecord) ToTMFObjectMapCreate() (TMFObjectMap, error) {
	if o == nil {
		return nil, errl.Errorf("object is nil")
	}
	if o.ContentMap != nil {
		return o.ContentMap, nil
	}

	return o.MustToTMFObjectMapCreate()
}

func (o *TMFRecord) MustToTMFObjectMapCreate() (TMFObjectMap, error) {

	omap, err := NewTMFObjectMap(o.Content)
	if err != nil {
		o.Validations.ObjectID = o.ID
		o.Validations.ObjectType = o.Type
		o.Validations.Valid = false
		o.Validations.Timestamp = time.Now()
		o.Validations.Errors = append(o.Validations.Errors, ValidationError{
			Field:   "object",
			Message: errl.Error(err).Error(),
			Code:    "INVALID_OBJECT",
		})

		err = errl.Errorf("failed to unmarshal object content: %w", err)
		return nil, err
	}

	o.Validations = omap.ValidateCreate(o.Type)
	if len(o.Validations.Errors) > 0 {
		slog.Error("invalid object", "id", o.ID, "type", o.Type, "errors", o.Validations.String())
		return nil, errl.Errorf("invalid object: %s", o.Validations.String())
	}

	o.ContentMap = omap

	return o.ContentMap, nil
}
