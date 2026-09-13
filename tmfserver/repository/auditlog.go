package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/hesusruiz/tmforum/internal/errl"
	"github.com/hesusruiz/tmforum/types"
)

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
			l.OldContent = make([]byte, len(oldContentStr.String))
			copy(l.OldContent, oldContentStr.String)
		}
		if newContentStr.Valid {
			l.NewContent = make([]byte, len(newContentStr.String))
			copy(l.NewContent, newContentStr.String)
		}
		logs = append(logs, l)
	}
	if err := rows.Err(); err != nil {
		return nil, errl.Errorf("iterating operation log rows: %w", err)
	}

	return logs, nil
}

// GetOperation retrieves one TMFOpLogRecord by seq.
// It returns a nil object and no error in the case it is not found.
func (repo *DBService) GetOperation(seq int64) (*TMFOpLogRecord, error) {
	var l TMFOpLogRecord
	var oldContentStr, newContentStr sql.NullString
	err := repo.db.QueryRow(`
		SELECT seq, action, object_id, object_type,
		       old_version, new_version, old_last_update, new_last_update,
		       json(old_content), json(new_content),
		       caller_id, server_id, access_token, created_at
		FROM tmf_operation_log
		WHERE seq = ?`,
		seq,
	).Scan(
		&l.Seq, &l.Action, &l.ObjectID, &l.ObjectType,
		&l.OldVersion, &l.NewVersion, &l.OldLastUpdate, &l.NewLastUpdate,
		&oldContentStr, &newContentStr,
		&l.CallerID, &l.ServerID, &l.AccessToken, &l.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil // Operation not found
	} else if err != nil {
		return nil, errl.Errorf("failed to get operation seq=%d: %w", seq, err)
	}
	if oldContentStr.Valid {
		l.OldContent = make([]byte, len(oldContentStr.String))
		copy(l.OldContent, oldContentStr.String)
	}
	if newContentStr.Valid {
		l.NewContent = make([]byte, len(newContentStr.String))
		copy(l.NewContent, newContentStr.String)
	}
	fmt.Println("OldContent: ", string(l.OldContent))

	return &l, nil
}

type SummaryOpLogRecord struct {
	Seq        int64  `json:"seq"`
	Action     string `json:"action"`
	ObjectID   string `json:"object_id"`
	ObjectType string `json:"object_type"`
	CallerID   string `json:"caller_id"`
	ServerID   string `json:"server_id"`
	CreatedAt  int64  `json:"created_at"`
}

// GetSummaryOperationLogs retrieves a page of operation log summaries and the total record count.
// page is 1-based (defaults to 1 if page < 1).
// If size <= 0, a default of 100 is used.
func (repo *DBService) GetSummaryOperationLogs(page, size int) (totalRecords int, logs []SummaryOpLogRecord, err error) {
	if page < 1 {
		page = 1
	}
	if size <= 0 {
		size = 100
	}

	countQuery := `SELECT COUNT(*) FROM tmf_operation_log`
	if err := repo.db.QueryRow(countQuery).Scan(&totalRecords); err != nil {
		return 0, nil, errl.Errorf("failed to count operation logs: %w", err)
	}

	if totalRecords == 0 {
		return 0, []SummaryOpLogRecord{}, nil
	}

	offset := (page - 1) * size

	query := `
		SELECT seq, action, object_id, object_type,
		       caller_id, server_id, created_at
		FROM tmf_operation_log
		ORDER BY seq ASC
		LIMIT ? OFFSET ?`

	rows, err := repo.db.Query(query, size, offset)
	if err != nil {
		return 0, nil, errl.Errorf("failed to query operation logs: %w", err)
	}
	defer rows.Close()

	logs = make([]SummaryOpLogRecord, 0, size)
	for rows.Next() {
		var l SummaryOpLogRecord
		if err := rows.Scan(
			&l.Seq, &l.Action, &l.ObjectID, &l.ObjectType,
			&l.CallerID, &l.ServerID, &l.CreatedAt,
		); err != nil {
			return 0, nil, errl.Errorf("failed to scan operation log row: %w", err)
		}
		logs = append(logs, l)
	}
	if err := rows.Err(); err != nil {
		return 0, nil, errl.Errorf("iterating operation log rows: %w", err)
	}

	return totalRecords, logs, nil
}
