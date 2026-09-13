package repository

import "database/sql"

func init() {
	RegisterMigration("20260419T000000", upKeepLatestVersion, nil)
}

func upKeepLatestVersion(db *sql.DB) error {
	// Create temporary table with new schema
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS tmf_object_tmp (
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
	);`)
	if err != nil {
		return err
	}

	// Insert only the latest version for each id and type
	_, err = db.Exec(`
		INSERT INTO tmf_object_tmp
		SELECT t.* FROM tmf_object t
		INNER JOIN (
			SELECT id, type, MAX(version) as max_version
			FROM tmf_object
			GROUP BY id, type
		) tm ON t.id = tm.id AND t.type = tm.type AND t.version = tm.max_version;
	`)
	if err != nil && err.Error() != "no such table: tmf_object" {
		// Ignore if table doesn't exist yet
		return err
	}

	// Drop old table if exists
	_, err = db.Exec(`DROP TABLE IF EXISTS tmf_object;`)
	if err != nil {
		return err
	}

	// Rename temporary table
	_, err = db.Exec(`ALTER TABLE tmf_object_tmp RENAME TO tmf_object;`)
	if err != nil {
		return err
	}

	return nil
}
