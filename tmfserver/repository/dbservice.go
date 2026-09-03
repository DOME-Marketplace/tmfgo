package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hesusruiz/tmforum/internal/errl"
	"github.com/hesusruiz/tmforum/types"
	"github.com/mattn/go-sqlite3"
)

// CreateTMFTableSQL is the SQL statement to create the table 'tmf_object', holding all objects of all types
//
// The most important field is "content", which stores the TM Forum object in SQLite jsonb format.
// Many other fields are just a copy of some fields in the "content" field, to facilitate SQL queries.
// The code is responsible for maintaining these "convenience" fields always in sync with the contents of the JSON object.
//
// The fields in the table are the following:
// "id": copy of the "id" field in the JSON representation.
// "type": copy of the "type" field in the JSON representation.
// "version": copy of the "version" field in the JSON representation.
// "api_version": the version of the TMF API that was used. This is for the future, as now it is always "v4".
// "seller": copy of the "seller" identification found in the "relatedParty" object with role "seller".
// "seller_operator": copy of the "seller_operator" identification found in the "relatedParty" object with role "seller_operator".
// "buyer": copy of the "buyer" identification found in the "relatedParty" object with role "buyer".
// "buyer_operator": copy of the "buyer_operator" identification found in the "relatedParty" object with role "buyer_operator".
// "last_update": copy of the "last_update" field in the JSON representation.
// "content": the content of the object in JSON format.
// "random": a random value to be used in queries for random ordering of lists, to make them fairer when presented to users.
// "created_at": the creation time of the object in Unix format (e.g. 123456789012)
// "updated_at": the update time of the object in Unix format (e.g. 123456789012)
const CreateTMFTableSQL = `CREATE TABLE IF NOT EXISTS tmf_object (
	"id" TEXT NOT NULL,
	"type" TEXT NOT NULL,
	"version" TEXT DEFAULT '',
	"api_version" TEXT DEFAULT '',
	"seller" TEXT DEFAULT '',
	"seller_operator" TEXT DEFAULT '',
	"buyer" TEXT DEFAULT '',
	"buyer_operator" TEXT DEFAULT '',
	"last_update" TEXT DEFAULT '',
	"content" BLOB NOT NULL,
	"random" INTEGER DEFAULT 0,
	"created_at" INTEGER,
	"updated_at" INTEGER,
	PRIMARY KEY ("id", "type")
);`

// TMFRecord represents the storage format of a generic TMForum object, associated to a record in the database.
type TMFRecord struct {
	ID             string           `db:"id"`
	Type           string           `db:"type"`
	Version        string           `db:"version"`
	APIVersion     string           `db:"api_version"`
	Seller         string           `db:"seller"`
	SellerOperator string           `db:"seller_operator"`
	Buyer          string           `db:"buyer"`
	BuyerOperator  string           `db:"buyer_operator"`
	LastUpdate     string           `db:"last_update"`
	Content        []byte           `db:"content"`
	Random         int              `db:"random"`
	CreatedAt      int64            `db:"created_at"`
	UpdatedAt      int64            `db:"updated_at"`
	ContentMap     TMFObjectMap     `db:"-"`
	Validations    ValidationResult `db:"-"`
}

// DeleteTMFTableSQL is the SQL statement to delete the table 'tmf_object'
const DeleteTMFTableSQL = `DROP TABLE IF EXISTS tmf_object;`

// VacuumSQL is the SQL statement to vacuum the database
const VacuumSQL = `VACUUM;`

// CreateTMFOpLogTableSQL is the SQL statement to create the table 'tmf_operation_log' and its indexes.
const CreateTMFOpLogTableSQL = `CREATE TABLE IF NOT EXISTS tmf_operation_log (
	"seq"                     INTEGER PRIMARY KEY,
	"action"                  TEXT NOT NULL,
	"object_id"               TEXT NOT NULL,
	"object_type"             TEXT NOT NULL,
	"old_version"             TEXT DEFAULT '',
	"new_version"             TEXT DEFAULT '',
	"old_last_update"         TEXT DEFAULT '',
	"new_last_update"         TEXT DEFAULT '',
	"old_content"             BLOB,
	"new_content"             BLOB,
	"caller_id"               TEXT DEFAULT '',
	"server_id"               TEXT DEFAULT '',
	"access_token"            TEXT DEFAULT '',
	"created_at"              INTEGER NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_op_log_object ON tmf_operation_log("object_id", "object_type");
CREATE INDEX IF NOT EXISTS idx_op_log_created ON tmf_operation_log("created_at");`

// DeleteTMFOpLogTableSQL is the SQL statement to delete the table 'tmf_operation_log'
const DeleteTMFOpLogTableSQL = `DROP TABLE IF EXISTS tmf_operation_log;`

// TMFOpLogRecord represents an entry in the tmf_operation_log table.
type TMFOpLogRecord struct {
	Seq           int64  `db:"seq"`
	Action        string `db:"action"`
	ObjectID      string `db:"object_id"`
	ObjectType    string `db:"object_type"`
	OldVersion    string `db:"old_version"`
	NewVersion    string `db:"new_version"`
	OldLastUpdate string `db:"old_last_update"`
	NewLastUpdate string `db:"new_last_update"`
	OldContent    []byte `db:"old_content"`
	NewContent    []byte `db:"new_content"`
	CallerID      string `db:"caller_id"`
	ServerID      string `db:"server_id"`
	AccessToken   string `db:"access_token"`
	CreatedAt     int64  `db:"created_at"`
}

// DBService is the database layer for TMF objects.
type DBService struct {
	db                 *sql.DB
	server_operator_id string
	stopCheckpoint     chan struct{}
	closeOnce          sync.Once
}

// NewDBService creates a new database service.
func NewDBService(dbName string, serverOperatorID string) (*DBService, error) {

	// Build the connection string with the parameters we want to use.
	// We specify the parameters even if they are the default ones, to make it explicit.

	// _journal_mode=WAL Enables Write-Ahead Logging for high concurrency (many readers, one writer).
	dbName = "file:" + dbName + "?_journal_mode=WAL"

	// _cache_size=-100000 Sets the cache size to 100000 kilobytes (100MB) instead of the default 2MB,
	// which is a good balance between performance and memory usage.
	// The negative sign indicates that the cache size is in kilobytes, not pages.
	dbName = dbName + "&_cache_size=-100000"

	// _busy_timeout=5000 Sets the 5 seconds timeout that the connection will wait and retry when the database is locked,
	// to mitigate the SQLITE_BUSY errors.
	dbName = dbName + "&_busy_timeout=5000"

	// _foreign_keys=on Enforces foreign key constraints.
	dbName = dbName + "&_foreign_keys=on"

	// _synchronous=NORMAL Offers a good balance between data safety and performance.
	dbName = dbName + "&_synchronous=NORMAL"

	// _txlock=immediate To prevent potential deadlocks or missed busy timeouts that can occur when a read transaction
	// is implicitly upgraded to a write transaction.
	dbName = dbName + "&_txlock=immediate"

	// _cache=shared Enables shared cache mode, which allows multiple connections to share the same cache.
	dbName = dbName + "&_cache=shared"

	// _defer_foreign_keys=on Delays the enforcement of foreign key constraints until the transaction is committed.
	dbName = dbName + "&_defer_foreign_keys=on"

	// Connect to the database.
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, errl.Errorf("failed to connect to database: %w", err)
	}
	slog.Info("db: opened", slog.String("dbName", dbName))

	// Create tables if they do not exist, and run migrations
	slog.Info("db: About to create tables if they do not exist")
	err = CreateTables(db)
	if err != nil {
		return nil, errl.Error(err)
	}

	repo := &DBService{
		db:                 db,
		server_operator_id: serverOperatorID,
		stopCheckpoint:     make(chan struct{}),
	}

	// Start a background timer to perform a passive WAL checkpoint every minute
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-repo.stopCheckpoint:
				return
			case <-ticker.C:
				if err := walCheckpointPassive(repo); err != nil {
					slog.Debug("passive WAL checkpoint failed", slog.Any("error", err))
				}
			}
		}
	}()

	return repo, nil
}

// CreateTables creates the tables in the database if they do not exist.
// It also handles automatic schema/data migration when possible.
func CreateTables(db *sql.DB) error {

	if _, err := db.Exec(CreateTMFTableSQL); err != nil {
		return errl.Errorf("failed to create tmf_object table: %w", err)
	}

	if _, err := db.Exec(CreateTMFOpLogTableSQL); err != nil {
		return errl.Errorf("failed to create tmf_operation_log table: %w", err)
	}

	if err := RunMigrationsUp(db); err != nil {
		return errl.Error(err)
	}

	return nil
}

// ErrObjectExists is returned when trying to create an object that already exists.
type ErrObjectExists struct {
	ID   string
	Type string
}

func (e *ErrObjectExists) Error() string {
	return fmt.Sprintf("object with id %s and type %s already exists", e.ID, e.Type)
}

func (e *ErrObjectExists) Is(target error) bool {
	switch target.(type) {
	case *ErrObjectExists:
		return true
	default:
		return false
	}
}

// ErrObjectNotFound is returned when trying to perform an operation on an object that does not exist.
type ErrObjectNotFound struct {
	ID   string
	Type string
}

func (e *ErrObjectNotFound) Error() string {
	return fmt.Sprintf("object with id %s and type %s not found", e.ID, e.Type)
}

func (e *ErrObjectNotFound) Is(target error) bool {
	switch target.(type) {
	case *ErrObjectNotFound:
		return true
	default:
		return false
	}
}

// Close closes the database connection.
func (repo *DBService) Close() error {
	repo.closeOnce.Do(func() {
		if repo.stopCheckpoint != nil {
			close(repo.stopCheckpoint)
		}
	})
	return repo.db.Close()
}

// getObjectWithTx retrieves a TMF object within an existing transaction.
func (repo *DBService) getObjectWithTx(tx *sql.Tx, id, objectType string) (*TMFRecord, error) {
	var obj TMFRecord
	err := tx.QueryRow(`
		SELECT id, type, version, api_version, seller, seller_operator, buyer, buyer_operator, last_update, json(content), random, created_at, updated_at
		FROM tmf_object
		WHERE id = :id AND type = :type`,
		sql.Named("id", id),
		sql.Named("type", objectType),
	).Scan(
		&obj.ID, &obj.Type, &obj.Version, &obj.APIVersion,
		&obj.Seller, &obj.SellerOperator, &obj.Buyer, &obj.BuyerOperator,
		&obj.LastUpdate, &obj.Content, &obj.Random, &obj.CreatedAt, &obj.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil // Object not found
	} else if err != nil {
		return nil, errl.Errorf("failed to get object id=%s type=%s: %w", id, objectType, err)
	}
	return &obj, nil
}

// recordOperation writes an audit/replication log entry to tmf_operation_log within an existing transaction.
func (repo *DBService) recordOperation(tx *sql.Tx, req *types.Request, action string, oldRec, newRec *TMFRecord) error {
	var objID, objType string
	var oldVersion, oldLastUpdate string
	var newVersion, newLastUpdate string
	var oldContent, newContent any

	if oldRec != nil {
		objID = oldRec.ID
		objType = oldRec.Type
		oldVersion = oldRec.Version
		oldLastUpdate = oldRec.LastUpdate
		if len(oldRec.Content) > 0 {
			oldContent = oldRec.Content
		}
	}

	if newRec != nil {
		objID = newRec.ID
		objType = newRec.Type
		newVersion = newRec.Version
		newLastUpdate = newRec.LastUpdate
		if len(newRec.Content) > 0 {
			newContent = newRec.Content
		}
	}

	var callerID, accessToken string
	if req != nil {
		callerID = req.AuthUser.OrganizationIdentifier
		accessToken = req.AuthUser.AccessToken
	}

	now := time.Now().Unix()

	query := `INSERT INTO tmf_operation_log (
		action, object_id, object_type,
		old_version, new_version, old_last_update, new_last_update,
		old_content, new_content,
		caller_id, server_id, access_token, created_at
	) VALUES (
		:action, :object_id, :object_type,
		:old_version, :new_version, :old_last_update, :new_last_update,
		CASE WHEN :old_content IS NULL THEN NULL ELSE jsonb(:old_content) END,
		CASE WHEN :new_content IS NULL THEN NULL ELSE jsonb(:new_content) END,
		:caller_id, :server_id, :access_token, :created_at
	)`

	_, err := tx.Exec(query,
		sql.Named("action", action),
		sql.Named("object_id", objID),
		sql.Named("object_type", objType),
		sql.Named("old_version", oldVersion),
		sql.Named("new_version", newVersion),
		sql.Named("old_last_update", oldLastUpdate),
		sql.Named("new_last_update", newLastUpdate),
		sql.Named("old_content", oldContent),
		sql.Named("new_content", newContent),
		sql.Named("caller_id", callerID),
		sql.Named("server_id", repo.server_operator_id),
		sql.Named("access_token", accessToken),
		sql.Named("created_at", now),
	)
	if err != nil {
		return errl.Errorf("failed to record operation in log: %w", err)
	}

	return nil
}

// CreateObject creates a new TMF object. Returns &ErrObjectExists if the object already existed.
func (repo *DBService) CreateObject(req *types.Request, obj *TMFRecord) error {
	if obj == nil {
		return errl.Errorf("object is nil")
	}
	slog.Debug("dbLayer: createObject", slog.String("id", obj.ID), slog.String("type", obj.Type), slog.String("version", obj.Version))

	// Make sure timestamps are correct
	now := time.Now()
	obj.CreatedAt = now.Unix()
	obj.UpdatedAt = now.Unix()

	tx, err := repo.db.Begin()
	if err != nil {
		return errl.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Execute the SQL
	_, err = tx.Exec(`INSERT INTO tmf_object
		(id, type, version, api_version, seller, seller_operator, buyer, buyer_operator, last_update, content, created_at, updated_at)
		VALUES (:id, :type, :version, :api_version, :seller, :seller_operator, :buyer, :buyer_operator, :last_update, jsonb(:content), :created_at, :updated_at)`,
		sql.Named("id", obj.ID),
		sql.Named("type", obj.Type),
		sql.Named("version", obj.Version),
		sql.Named("api_version", obj.APIVersion),
		sql.Named("seller", obj.Seller),
		sql.Named("seller_operator", obj.SellerOperator),
		sql.Named("buyer", obj.Buyer),
		sql.Named("buyer_operator", obj.BuyerOperator),
		sql.Named("last_update", obj.LastUpdate),
		sql.Named("content", obj.Content),
		sql.Named("created_at", obj.CreatedAt),
		sql.Named("updated_at", obj.UpdatedAt),
	)
	if err != nil {
		if sqliteErr, ok := errors.AsType[sqlite3.Error](err); ok {
			if sqliteErr.Code == sqlite3.ErrConstraint && sqliteErr.ExtendedCode == sqlite3.ErrConstraintPrimaryKey {
				return &ErrObjectExists{ID: obj.ID, Type: obj.Type}
			}
		}
		return errl.Errorf("failed to create object id=%s type=%s: %w", obj.ID, obj.Type, err)
	}

	if err := repo.recordOperation(tx, req, "CREATE", nil, obj); err != nil {
		return err
	}

	return tx.Commit()
}

// GetObject retrieves a TMF object by its ID and type.
// If the object is not found anywhere, it returns a nil object and no error.
func (repo *DBService) GetObject(req *types.Request, id, objectType string) (*TMFRecord, error) {
	slog.Debug("dbLayer: getObject", slog.String("id", id), slog.String("type", objectType))

	var obj TMFRecord
	err := repo.db.QueryRow(`
		SELECT id, type, version, api_version, seller, seller_operator, buyer, buyer_operator, last_update, json(content), random, created_at, updated_at
		FROM tmf_object
		WHERE id = :id AND type = :type`,
		sql.Named("id", id),
		sql.Named("type", objectType),
	).Scan(
		&obj.ID, &obj.Type, &obj.Version, &obj.APIVersion,
		&obj.Seller, &obj.SellerOperator, &obj.Buyer, &obj.BuyerOperator,
		&obj.LastUpdate, &obj.Content, &obj.Random, &obj.CreatedAt, &obj.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		slog.Info("DBLayer: Object not found", slog.String("id", id), slog.String("type", objectType))
		return nil, nil // Object not found
	} else if err != nil {
		err = errl.Errorf("failed to get object id=%s type=%s: %w", id, objectType, err)
	}
	return &obj, err
}

// UpdateObject updates an existing TMF object row.
//
// Returns:
//   - ErrObjectNotFound  – no row exists for the given (id, type).
func (repo *DBService) UpdateObject(req *types.Request, obj *TMFRecord) error {
	if obj == nil {
		return errl.Errorf("object is nil")
	}
	slog.Debug("dbLayer: UpdateObject", slog.String("id", obj.ID), slog.String("type", obj.Type), slog.String("version", obj.Version))

	tx, err := repo.db.Begin()
	if err != nil {
		return errl.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	oldObj, err := repo.getObjectWithTx(tx, obj.ID, obj.Type)
	if err != nil {
		return err
	}
	if oldObj == nil {
		return &ErrObjectNotFound{ID: obj.ID, Type: obj.Type}
	}

	// Make sure timestamps are correct
	obj.UpdatedAt = time.Now().Unix()

	// Update the row for this object, storing the latest version and content.
	// Note: seller and buyer are intentionally excluded from the SET clause – they cannot
	// be changed after the object is created.
	res, err := tx.Exec(`UPDATE tmf_object
		SET   version     = :version,
		      last_update = :last_update,
		      content     = jsonb(:content),
		      updated_at  = :updated_at
		WHERE id      = :id
		  AND type    = :type`,
		sql.Named("version", obj.Version),
		sql.Named("last_update", obj.LastUpdate),
		sql.Named("content", obj.Content),
		sql.Named("updated_at", obj.UpdatedAt),
		sql.Named("id", obj.ID),
		sql.Named("type", obj.Type),
	)
	if err != nil {
		return errl.Errorf("failed to update object id=%s type=%s: %w", obj.ID, obj.Type, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errl.Errorf("failed to get rows affected object id=%s type=%s: %w", obj.ID, obj.Type, err)
	}

	if rowsAffected == 0 {
		return &ErrObjectNotFound{ID: obj.ID, Type: obj.Type}
	}

	if err := repo.recordOperation(tx, req, "UPDATE", oldObj, obj); err != nil {
		return err
	}

	return tx.Commit()
}

// UpsertObject creates or updates a TMF object.
//
// Semantics:
//   - If no row exists for the given (id, type): the object is inserted.
//   - If the same (id, type) already exists: the row is updated in-place
//     (version, content, last_update, updated_at). Seller and buyer are NOT changed.
func (repo *DBService) UpsertObject(req *types.Request, obj *TMFRecord) error {
	if obj == nil {
		return errl.Errorf("object is nil")
	}
	slog.Debug("dbLayer: UpsertObject", slog.String("id", obj.ID), slog.String("type", obj.Type), slog.String("version", obj.Version))

	tx, err := repo.db.Begin()
	if err != nil {
		return errl.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	oldObj, err := repo.getObjectWithTx(tx, obj.ID, obj.Type)
	if err != nil {
		return err
	}

	now := time.Now()
	obj.UpdatedAt = now.Unix()

	var action string
	if oldObj == nil {
		action = "CREATE"
		obj.CreatedAt = now.Unix()

		_, err = tx.Exec(`INSERT INTO tmf_object
			(id, type, version, api_version, seller, seller_operator, buyer, buyer_operator, last_update, content, created_at, updated_at)
			VALUES (:id, :type, :version, :api_version, :seller, :seller_operator, :buyer, :buyer_operator, :last_update, jsonb(:content), :created_at, :updated_at)`,
			sql.Named("id", obj.ID),
			sql.Named("type", obj.Type),
			sql.Named("version", obj.Version),
			sql.Named("api_version", obj.APIVersion),
			sql.Named("seller", obj.Seller),
			sql.Named("seller_operator", obj.SellerOperator),
			sql.Named("buyer", obj.Buyer),
			sql.Named("buyer_operator", obj.BuyerOperator),
			sql.Named("last_update", obj.LastUpdate),
			sql.Named("content", obj.Content),
			sql.Named("created_at", obj.CreatedAt),
			sql.Named("updated_at", obj.UpdatedAt),
		)
		if err != nil {
			return errl.Errorf("failed to insert object id=%s type=%s: %w", obj.ID, obj.Type, err)
		}
	} else {
		action = "UPDATE"
		obj.CreatedAt = oldObj.CreatedAt

		_, err = tx.Exec(`UPDATE tmf_object
			SET   version     = :version,
			      last_update = :last_update,
			      content     = jsonb(:content),
			      updated_at  = :updated_at
			WHERE id      = :id
			  AND type    = :type`,
			sql.Named("version", obj.Version),
			sql.Named("last_update", obj.LastUpdate),
			sql.Named("content", obj.Content),
			sql.Named("updated_at", obj.UpdatedAt),
			sql.Named("id", obj.ID),
			sql.Named("type", obj.Type),
		)
		if err != nil {
			return errl.Errorf("failed to update object id=%s type=%s: %w", obj.ID, obj.Type, err)
		}
	}

	if err := repo.recordOperation(tx, req, action, oldObj, obj); err != nil {
		return err
	}

	return tx.Commit()
}

// DeleteObject deletes a TMF object by its ID and type.
func (repo *DBService) DeleteObject(req *types.Request, id, resourceName string) error {
	slog.Debug("dbLayer: deleteObject", slog.String("id", id), slog.String("type", resourceName))

	tx, err := repo.db.Begin()
	if err != nil {
		return errl.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	oldObj, err := repo.getObjectWithTx(tx, id, resourceName)
	if err != nil {
		return err
	}

	// Execute the SQL
	res, err := tx.Exec("DELETE FROM tmf_object WHERE id = ? AND type = ?", id, resourceName)
	if err != nil {
		return errl.Errorf("failed to delete object id=%s type=%s: %w", id, resourceName, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errl.Errorf("failed to get rows affected object id=%s type=%s: %w", id, resourceName, err)
	}

	if rowsAffected > 0 {
		if err := repo.recordOperation(tx, req, "DELETE", oldObj, nil); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// GetOperationLogs retrieves operation log entries with seq > afterSeq, ordered by seq ASC.
// If limit <= 0, a default of 100 is used.
func (repo *DBService) GetOperationLogs(afterSeq int64, limit int) ([]TMFOpLogRecord, error) {
	if limit <= 0 {
		limit = 100
	}

	query := `
		SELECT seq, action, object_id, object_type,
		       old_version, new_version, old_last_update, new_last_update,
		       json(old_content), json(new_content),
		       caller_id, server_id, access_token, created_at
		FROM tmf_operation_log
		WHERE seq > ?
		ORDER BY seq ASC
		LIMIT ?`

	rows, err := repo.db.Query(query, afterSeq, limit)
	if err != nil {
		return nil, errl.Errorf("failed to query operation logs: %w", err)
	}
	defer rows.Close()

	var logs []TMFOpLogRecord
	for rows.Next() {
		var l TMFOpLogRecord
		var oldContentStr, newContentStr sql.NullString
		if err := rows.Scan(
			&l.Seq, &l.Action, &l.ObjectID, &l.ObjectType,
			&l.OldVersion, &l.NewVersion, &l.OldLastUpdate, &l.NewLastUpdate,
			&oldContentStr, &newContentStr,
			&l.CallerID, &l.ServerID, &l.AccessToken, &l.CreatedAt,
		); err != nil {
			return nil, errl.Errorf("failed to scan operation log row: %w", err)
		}
		if oldContentStr.Valid {
			l.OldContent = []byte(oldContentStr.String)
		}
		if newContentStr.Valid {
			l.NewContent = []byte(newContentStr.String)
		}
		logs = append(logs, l)
	}
	if err := rows.Err(); err != nil {
		return nil, errl.Errorf("iterating operation log rows: %w", err)
	}

	return logs, nil
}

// ObjectFilter is a function that filters TMF objects. If it returns false, the object is excluded from the result.
type ObjectFilter func(obj *TMFRecord) bool

// ListObjects retrieves TMF objects of a given type, returning only the latest version for each unique ID.
// It supports pagination, filtering, and sorting according to TMF630 guidelines.
// filter is a callback function that is called for each object. If it returns false, the object is excluded from the result.
func (repo *DBService) ListObjects(req *types.Request, filter ObjectFilter) ([]TMFRecord, error) {
	healthRequest := false
	resourceName := ""
	var queryParams url.Values
	if req != nil {
		healthRequest = req.HealthRequest
		resourceName = req.ResourceName
		queryParams = req.QueryParams
	}
	if !healthRequest {
		slog.Debug("dbLayer: listObjects", "type", resourceName, "queryParams", queryParams)
	}

	// Parse the parameters according to TM Forum specs and build the SELECT
	baseQuery, args, limit, offset, err := BuildSelectFromParms(resourceName, queryParams)
	if err != nil {
		return nil, errl.Errorf("failed to build select query: %w", err)
	}
	if !healthRequest {
		fmt.Printf("SQL: %s\nARGS: %v\n", baseQuery, args)
	}

	var objs []TMFRecord
	var offsetCounter int

	// Run the SQL
	rows, err := repo.db.Query(baseQuery, args...)
	if err != nil {
		return nil, errl.Errorf("performing query %s with args %v: %w", baseQuery, args, err)
	}
	defer func() {
		_ = rows.Close()
	}()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var obj TMFRecord
		if err := rows.Scan(&obj.ID, &obj.Type, &obj.Version, &obj.APIVersion,
			&obj.Seller, &obj.Buyer, &obj.LastUpdate, &obj.Content, &obj.CreatedAt, &obj.UpdatedAt); err != nil {

			return nil, errl.Errorf("iterating over rows in query %s with args %v: %w", baseQuery, args, err)

		}

		// Callback to the caller for filtering/ammendment of the object
		if filter != nil && !filter(&obj) {
			// The callback said that we should not include this object
			continue
		}

		// If the object has passed all checks, we still have to discard the first 'offset' objects, as specified by the user
		if offsetCounter < offset {
			offsetCounter++
			continue
		}

		// Now we add the object to the result array and check if we reached the limit as specified by the user
		objs = append(objs, obj)

		if len(objs) >= limit {
			break
		}

	}
	if err = rows.Err(); err != nil {
		return nil, errl.Errorf("performing query %s with args %v: %w", baseQuery, args, err)
	}

	return objs, err
}

// BuildSelectFromParms creates a SELECT statement based on the query values.
// Some keys are columns in the database row, but most of them are in the JSON object in the 'content' column
// For TMF objects with same id, selects the one with the latest version.
func BuildSelectFromParms(resourceName string, queryValues url.Values) (query string, arguments []any, qlimit int, qoffset int, theerr error) {

	// Default values if the user did not specify them. -1 is equivalent to no values provided.
	// Offset and limit will not be included in the SQL query, but will be used by the caller to limit the number of objects returned,
	// after filtering is applied to the set of objects returned by the database SELECT statement.
	var limit = -1
	var offset = -1

	var queryBuilder StringRenderer
	var args []any

	// The main SELECT statement
	queryBuilder.Render(
		`SELECT id, type, version, api_version, seller, buyer, last_update, json(content), created_at, updated_at FROM tmf_object`,
	)

	// The main WHERE clause: normally we expect the resource name of object to be specified, but we support a query for all object types
	// There will be additionally inner SELECTs when we use the content column
	if len(resourceName) > 0 {
		queryBuilder.Render(" WHERE type = ?")
		args = append(args, resourceName)
	}

	// Build the WHERE by processing the query values specified by the user
	for key, values := range queryValues {

		// Create additional parts of the SELECT, with some special processing
		switch key {
		case "sort", "fields":
			// TODO: implement processing for these parameters. For the moment, they must be implemented by the caller.
			continue

		case "limit":
			// Just extract the value for later, it will not be used in the SELECT
			limitStr := queryValues.Get("limit")
			if limitStr != "" {
				if l, err := strconv.Atoi(limitStr); err == nil {
					limit = l
				}
			}

		case "offset":
			// Just extract the value for later, it will not be used in the SELECT
			offsetStr := queryValues.Get("offset")
			if offsetStr != "" {
				if l, err := strconv.Atoi(offsetStr); err == nil {
					offset = l
				}
			}

		case "seller", "buyer":
			// A shortcut for DOME and ISBE, to simplify life to applications (but can be also done in a TMF-compliant way).
			// Special processing to allow specifying multiple values in the form 'seller=id1,id2,id3'.
			// We also support the standard HTTP query strings like 'seller=id1,id2&seller=id3'
			vals := processValues(values)

			// Use either an equality (when one element) or an inclusion expression (when several)
			if len(vals) == 1 {
				queryBuilder.Render(" AND ", key, " = ?")
			} else if len(vals) > 1 {
				queryBuilder.Render(" AND ", key, " IN ").RenderSQLList(vals)
			}
			for _, v := range vals {
				args = append(args, v)
			}

		case "category.id", "productSpecification.id":
			// Simplification of the query for common array fields at the first level of the JSON object

			object := strings.TrimSuffix(key, ".id")

			// Special processing because TMForum allows to specify multiple values
			// in the form 'lifecycleStatus=Launched,Active'
			vals := processValues(values)

			if len(vals) == 1 {
				queryBuilder.Render(
					" AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.", object, "') WHERE json_extract(value, '$.id') = ?)",
				)
			} else if len(vals) > 1 {
				queryBuilder.Render(
					" AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.", object, "') WHERE json_extract(value, '$.id') IN ").RenderSQLList(vals).Render(")")
			}
			for _, v := range vals {
				args = append(args, v)
			}

		case "organizationIdentification[*].identificationId", "organizationIdentification.identificationId",
			"individualIdentification[*].identificationId", "individualIdentification.identificationId", "individualIdentification.id":
			// Simplification of the query in the identification arrays for the special case of organization identification data
			arrayName := "organizationIdentification"
			keyName := "identificationId"

			if strings.HasPrefix("individualIdentification", key) {
				arrayName = "individualIdentification"
				keyName = "identificationId"
			}

			// Special processing because TMForum allows to specify multiple values
			// in the form 'lifecycleStatus=Launched,Active'
			vals := processValues(values)

			if len(vals) == 1 {
				queryBuilder.Render(
					" AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.", arrayName, "') WHERE json_extract(value, '$.", keyName, "') = ?)",
				)
			} else if len(vals) > 1 {
				queryBuilder.Render(
					" AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.", arrayName, "') WHERE json_extract(value, '$.", keyName, "') IN ").RenderSQLList(vals).Render(")")
			}
			for _, v := range vals {
				args = append(args, v)
			}

		default:

			// Special processing because TMForum allows to specify multiple values
			// in the form 'lifecycleStatus=Launched,Active'
			vals := processValues(values)

			// We perform special processing when the key is simple (no dots), to use a simple and more efficient SQL expression.
			pathParts := strings.Split(key, ".")
			if len(pathParts) == 1 {

				if len(vals) == 1 {
					queryBuilder.Render(" AND content->>'$.", key, "' = ?")
				} else {
					queryBuilder.Render(" AND content->>'$.", key, "' IN ").RenderSQLList(vals)
				}
				for _, v := range vals {
					args = append(args, v)
				}
			} else {
				subSql, _, err := GenerateRecursiveJSONQuery("tmf_object", key, vals)
				if err != nil {
					return "", nil, 0, 0, err
				}
				queryBuilder.Render(" AND ", subSql)
				for _, v := range vals {
					args = append(args, v)
				}
			}

		}
	}

	// Build the query, with the statement and the arguments to be used
	sql := queryBuilder.String()

	return sql, args, limit, offset, nil
}

// GenerateRecursiveJSONQuery generates a SQL query to search for a value in a JSON object, given a JSON path where some elements may be arrays.
func GenerateRecursiveJSONQuery(tableName string, pathInput string, values []string) (string, string, error) {

	if len(values) == 0 {
		return "", "", fmt.Errorf("invalid values: no values provided for recursive JSON queries")
	}

	pathParts := strings.Split(pathInput, ".")

	var like StringRenderer
	like.WriteString("$")

	for _, p := range pathParts {
		field, index := extractFromBrackets(p)
		if index == "*" {
			index = "%"
		}
		if index == "" {
			like.Render('.', field, '%')
		} else {
			like.Render('.', field, '[', index, ']')
		}
	}

	if len(values) == 1 {
		sql := "EXISTS (SELECT 1 FROM json_tree(" + tableName + ".content) WHERE json_tree.fullkey LIKE '" + like.String() + "' AND json_tree.value = ?)"
		return strings.TrimSpace(sql), values[0], nil
	} else {
		var buf StringRenderer
		buf.Render("EXISTS (SELECT 1 FROM json_tree(" + tableName + ".content) WHERE json_tree.fullkey LIKE '" + like.String() + "' AND json_tree.value IN ")
		buf.RenderSQLList(values)
		buf.Render(")")
		return strings.TrimSpace(buf.String()), "", nil
	}
}

// StringRenderer is a utility for efficiently building strings by rendering values to a buffer.
type StringRenderer struct {
	strings.Builder
}

// Render renders the given inputs to the buffer.
func (r *StringRenderer) Render(inputs ...any) *StringRenderer {
	for _, s := range inputs {
		switch v := s.(type) {
		case string:
			r.WriteString(v)
		case []byte:
			r.Write(v)
		case int:
			r.WriteString(strconv.FormatInt(int64(v), 10))
		case byte:
			r.WriteByte(v)
		case rune:
			r.WriteRune(v)
		default:
			slog.Error("attemping to write something not a string, int, rune, []byte or byte: %T", s)
		}
	}
	return r
}

// Renderln renders the given inputs to the buffer, followed by a newline.
func (r *StringRenderer) Renderln(inputs ...any) *StringRenderer {
	r.Render(inputs...)
	r.Render('\n')
	return r
}

// RenderSQLList renders an SQL argument list.
// The actual values are not used, just the lengh of the list.
func (r *StringRenderer) RenderSQLList(inputs []string) *StringRenderer {
	r.Render("(")

	for i := range inputs {
		if i > 0 {
			r.Render(",")
		}
		r.Render("?")
	}
	r.Render(")")
	return r
}

// processValues converts from a slice of strings, each possibly a comma separeted set of values, to a slice of strings
func processValues(values []string) []string {
	var vals []string
	for _, v := range values {
		parts := strings.Split(v, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
			vals = append(vals, parts[i])
		}
	}
	return vals
}

func extractFromBrackets(s string) (string, string) {
	// Find the position of the opening bracket
	start := strings.Index(s, "[")
	if start == -1 {
		return s, ""
	}

	// Find the position of the closing bracket
	// We search from 'start' onwards to be safe
	end := strings.Index(s[start:], "]")
	if end == -1 {
		return s, ""
	}

	// The prefix is everything before the opening bracket
	prefix := s[:start]

	// The content is everything between '[' and ']'
	// 'end' is relative to s[start:], so the absolute position of ']' is start + end
	content := s[start+1 : start+end]

	return prefix, content
}
