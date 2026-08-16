package migrations

import (
	"database/sql"
	"log/slog"

	"github.com/hesusruiz/tmforum/internal/errl"
	"github.com/hesusruiz/tmforum/tmfserver/repository"
)

func init() {
	repository.RegisterMigration("20260425T192500", migration_up_20260425T192500, nil)
}

func migration_up_20260425T192500(db *sql.DB) error {
	slog.Info("Migrating all records to JSONB for the content column")

	// Perform an update for all records, updating the content column to store a jsonb(column)
	_, err := db.Exec(`UPDATE tmf_object SET content = jsonb(content)`)
	if err != nil {
		return errl.Errorf("migrating to JSONBcontent: %w", err)
	}

	return nil
}
