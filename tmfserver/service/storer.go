package service

import (
	repo "github.com/hesusruiz/tmforum/tmfserver/repository"
	"github.com/hesusruiz/tmforum/types"
)

// Type aliases for errors so other service files can still use them directly
type ErrObjectExists = repo.ErrObjectExists
type ErrObjectNotFound = repo.ErrObjectNotFound

// TMFStorer abstracts persistence operations for TMF objects.
// It is used for plugging-in different persistence systems
type TMFStorer interface {
	CreateObject(req *types.Request, obj *repo.TMFRecord) error
	GetObject(req *types.Request, id, resourceName string) (*repo.TMFRecord, error)
	UpdateObject(req *types.Request, obj *repo.TMFRecord) error
	UpsertObject(req *types.Request, obj *repo.TMFRecord) error
	DeleteObject(req *types.Request, id, resourceName string) error
	ListObjects(req *types.Request, filter repo.ObjectFilter) ([]repo.TMFRecord, error)
	GetOperationLogs(afterSeq int64, limit int) ([]repo.TMFOpLogRecord, error)
	GetSummaryOperationLogs(page, size int) (totalRecords int, logs []repo.SummaryOpLogRecord, err error)
}
