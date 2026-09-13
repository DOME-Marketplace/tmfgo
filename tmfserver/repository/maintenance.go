package repository

import (
	"log/slog"
	"time"

	"github.com/DOME-Marketplace/tmfgo/config"
	"github.com/DOME-Marketplace/tmfgo/internal/errl"
	"github.com/DOME-Marketplace/tmfgo/sqlitesync"
	"github.com/cloudflare/tableflip"
)

// ScheduleMaintenance schedules periodic database maintenance tasks like VACUUM or backups.
// It runs the maintenance task every 2 hours exactly at odd hours (1, 3, 5, ...).
// Additionally, it schedules a restart every day at 3 o'clock.
// The restart uses the https://github.com/cloudflare/tableflip graceful restart to keep client connections
// alive during the restart.
func ScheduleMaintenance(configuration *config.Config, repo *DBService, upg *tableflip.Upgrader) {

	// Perform an initial maintenance
	err := PerformMaintenance(repo, configuration.Dbname)
	if err != nil {
		slog.Error("failed to perform initial maintenance", "error", err)
	}

	go func() {
		// Every hour, at minute 0, perform maintenance and restart the process each night at 3am
		interval := time.Hour

		now := time.Now()
		nextRun := now.Truncate(time.Hour).Add(time.Hour)

		for {
			d := time.Until(nextRun)
			// Safety measure, if we try to schedule close to the hour
			if d <= 0 {
				nextRun = nextRun.Add(interval)
				continue
			}

			slog.Info("Database maintenance scheduled", "next_run", nextRun, "interval", interval)

			select {
			case <-repo.stopMaintenance:
				return
			case <-time.After(d):

				now := time.Now()

				if err := PerformMaintenance(repo, configuration.Dbname); err != nil {
					slog.Error("failed to perform maintenance", "error", err)
				}

				// If its 3 o'clock, restart the program
				if now.Hour() == 3 {
					slog.Info("Executing scheduled restart...")
					if err := upg.Upgrade(); err != nil {
						slog.Error("Upgrade failed", "error", err)
					}
				}

				nextRun = nextRun.Add(interval)
			}

		}
	}()

}

// PerformMaintenance performs the actual database maintenance tasks (WAL checkpoint,VACUUM and Backup).
func PerformMaintenance(repo *DBService, dbPath string) error {
	slog.Info("Executing scheduled database maintenance...")

	// Perform a checkpoint to ensure the WAL file is truncated
	start := time.Now()
	if err := forceWalTruncate(repo); err != nil {
		return errl.Errorf("failed to truncate WAL file: %w", err)
	}
	slog.Info("WAL file truncated successfully", "elapsed", time.Since(start))

	// Perform VACUUM
	start = time.Now()
	if _, err := repo.db.Exec(VacuumSQL); err != nil {
		return errl.Errorf("failed to vacuum database: %w", err)
	}
	slog.Info("Database vacuumed successfully", "elapsed", time.Since(start))

	// Perform Backup
	start = time.Now()
	if err := sqlitesync.Backup(dbPath); err != nil {
		return errl.Errorf("failed to backup database: %w", err)
	}
	slog.Info("Database backup completed successfully", "elapsed", time.Since(start))

	return nil
}

// forceWalTruncate forces a WAL checkpoint to truncate the WAL file
// TRUNCATE ensures the checkpoint runs to completion, deleting the WAL file.
func forceWalTruncate(repo *DBService) error {
	_, err := repo.db.Exec("PRAGMA wal_checkpoint(TRUNCATE)")
	return err
}
