package service

import (
	"log/slog"

	"github.com/hesusruiz/tmforum/internal/errl"
	repo "github.com/hesusruiz/tmforum/tmfserver/repository"
)

// Type aliases for errors so other service files can still use them directly
type ErrObjectExists = repo.ErrObjectExists
type ErrObjectNotFound = repo.ErrObjectNotFound

// CreateObject creates a new TMF object. Returns &ErrObjectExists if the object already existed.
func (svc *Service) CreateObject(obj *repo.TMFRecord) error {
	if obj == nil {
		return errl.Errorf("object is nil")
	}
	slog.Debug("dbLayer: createObject", slog.String("id", obj.ID), slog.String("type", obj.Type), slog.String("version", obj.Version))

	return svc.storage.CreateObject(obj)
}

// GetObject retrieves a TMF object by its ID and type, returning the latest version.
func (svc *Service) GetObject(id, objectType string) (*repo.TMFRecord, error) {
	slog.Debug("dbLayer: getObject", slog.String("id", id), slog.String("type", objectType))
	return svc.storage.GetObject(id, objectType)
}

// UpdateObject updates an existing TMF object. Reports an error if the object is not found.
func (svc *Service) UpdateObject(obj *repo.TMFRecord) error {
	slog.Debug("dbLayer: UpdateObject", slog.String("id", obj.ID), slog.String("type", obj.Type), slog.String("version", obj.Version))
	return svc.storage.UpdateObject(obj)
}

// UpsertObject creates or updates a TMF object.
func (svc *Service) UpsertObject(obj *repo.TMFRecord) error {
	slog.Debug("dbLayer: upsertObject", slog.String("id", obj.ID), slog.String("type", obj.Type), slog.String("version", obj.Version))
	return svc.storage.UpsertObject(obj)
}

// DeleteObject deletes a TMF object by its ID and type.
func (svc *Service) DeleteObject(id, objectType string) error {
	slog.Debug("dbLayer: deleteObject", slog.String("id", id), slog.String("type", objectType))
	return svc.storage.DeleteObject(id, objectType)
}

// ListObjects retrieves TMF objects of a given type, returning only the latest version for each unique ID.
func (svc *Service) ListObjects(req *Request, filter repo.ObjectFilter) ([]repo.TMFRecord, error) {
	if !req.HealthRequest {
		slog.Debug("dbLayer: listObjects", "type", req.ResourceName, "queryParams", req.QueryParams)
	}
	return svc.storage.ListObjects(req.HealthRequest, req.ResourceName, req.QueryParams, filter)
}
