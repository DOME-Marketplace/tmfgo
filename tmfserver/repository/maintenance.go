package repository

import (
	"log/slog"
	"time"

	"github.com/cloudflare/tableflip"
	"github.com/hesusruiz/tmforum/config"
	"github.com/hesusruiz/tmforum/sqlitesync"
)

// ScheduleMaintenance schedules periodic database maintenance tasks like VACUUM or backups.
// It runs the maintenance task every 2 hours exactly at odd hours (1, 3, 5, ...).
// Additionally, it schedules a restart every day at 3 o'clock.
// The restart uses the https://github.com/cloudflare/tableflip graceful restart to keep client connections
// alive during the restart.
func ScheduleMaintenance(configuration *config.Config, repo *DBService, upg *tableflip.Upgrader) {

	// Perform an initial maintenance
	PerformMaintenance(repo, configuration.Dbname)

	// interval := 2 * time.Hour
	interval := time.Hour

	go func() {
		// Calculate the next run time: the next odd hour (1, 3, 5, ...)
		now := time.Now()
		// Start with the next full hour
		nextRun := now.Truncate(time.Hour).Add(time.Hour)

		// If that hour is even, add another hour to make it odd
		if nextRun.Hour()%2 == 0 {
			nextRun = nextRun.Add(time.Hour)
		}

		slog.Info("Database maintenance scheduled", "next_run", nextRun, "interval", interval)

		for {
			time.Sleep(time.Until(nextRun))
			PerformMaintenance(repo, configuration.Dbname)

			// If its 3 o'clock, restart the program
			if time.Now().Hour() == 3 {
				slog.Info("Executing scheduled restart...")
				if err := upg.Upgrade(); err != nil {
					slog.Error("Upgrade failed", "error", err)
				}
			}

			nextRun = nextRun.Add(interval)
		}
	}()
}

// PerformMaintenance performs the actual database maintenance tasks (VACUUM and Backup).
func PerformMaintenance(repo *DBService, dbPath string) {
	slog.Info("Executing scheduled database maintenance...")

	// Perform a checkpoint to ensure the WAL file is truncated
	start := time.Now()
	if err := forceWalTruncate(repo); err != nil {
		slog.Error("failed to truncate WAL file", "error", err, "elapsed", time.Since(start))
	} else {
		slog.Info("WAL file truncated successfully", "elapsed", time.Since(start))
	}

	// Perform VACUUM
	start = time.Now()
	if _, err := repo.db.Exec(VacuumSQL); err != nil {
		slog.Error("failed to vacuum database", "error", err, "elapsed", time.Since(start))
	} else {
		slog.Info("Database vacuumed successfully", "elapsed", time.Since(start))
	}

	// Perform Backup
	start = time.Now()
	if err := sqlitesync.Backup(dbPath); err != nil {
		slog.Error("failed to backup database", "error", err, "elapsed", time.Since(start))
	} else {
		slog.Info("Database backup completed successfully", "elapsed", time.Since(start))
	}
}

func forceWalTruncate(repo *DBService) error {
	// TRUNCATE ensures the checkpoint runs to completion, deleting the WAL file.
	// This call can block briefly if a write is in progress.
	_, err := repo.db.Exec("PRAGMA wal_checkpoint(TRUNCATE)")
	return err
}

func walCheckpointPassive(repo *DBService) error {
	// CHECKPOINT ensures the checkpoint runs to completion, deleting the WAL file.
	// This call does not block, even if a write is in progress.
	_, err := repo.db.Exec("PRAGMA wal_checkpoint(PASSIVE)")
	return err
}
