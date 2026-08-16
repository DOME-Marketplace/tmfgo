package migrations

import (
	"database/sql"
	"log/slog"

	"github.com/hesusruiz/tmforum/tmfserver/repository"
)

func init() {
	repository.RegisterMigration("20251102T203201", migration_up_20251102T203201, nil)
}

func migration_up_20251102T203201(db *sql.DB) error {
	slog.Info("Hello from migration_up_20251102T203201")

	// According to the SQLite manual, this is a general procedure for making arbitrary changes to the schema design of some table X.
	// https://www.sqlite.org/lang_altertable.html
	//
	// The only schema altering commands directly supported by SQLite are the "rename table", "rename column",
	// "add column", "drop column" commands.
	// However, applications can make other arbitrary changes to the format of a table using a simple sequence of operations.
	// The steps to make arbitrary changes to the schema design of some table X are as follows:
	//
	// 1. If foreign key constraints are enabled, disable them using PRAGMA foreign_keys=OFF.
	// 2. Start a transaction.
	// 3. Remember the format of all indexes, triggers, and views associated with table X.
	//    This information will be needed in step 8 below.
	//    One way to do this is to run a query like the following: SELECT type, sql FROM sqlite_schema WHERE tbl_name='X'.
	// 4. Use CREATE TABLE to construct a new table "new_X" that is in the desired revised format of table X.
	//    Make sure that the name "new_X" does not collide with any existing table name, of course.
	// 5. Transfer content from X into new_X using a statement like: INSERT INTO new_X SELECT ... FROM X.
	// 6. Drop the old table X: DROP TABLE X.
	// 7. Change the name of new_X to X using: ALTER TABLE new_X RENAME TO X.
	// 8. Use CREATE INDEX, CREATE TRIGGER, and CREATE VIEW to reconstruct indexes, triggers, and views associated with table X.
	//    Perhaps use the old format of the triggers, indexes, and views saved from step 3 above as a guide,
	//    making changes as appropriate for the alteration.
	// 9. If any views refer to table X in a way that is affected by the schema change, then drop those views
	//    using DROP VIEW and recreate them with whatever changes are necessary to accommodate the schema change using CREATE VIEW.
	// 10. If foreign key constraints were originally enabled then run PRAGMA foreign_key_check
	//     to verify that the schema change did not break any foreign key constraints.
	// 11. Commit the transaction started in step 2.
	// 12. If foreign keys constraints were originally enabled, reenable them now.

	return nil
}
