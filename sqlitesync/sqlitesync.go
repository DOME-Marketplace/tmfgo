package sqlitesync

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/DOME-Marketplace/tmfgo/internal/errl"
)

// Sync executes the sqlite3_rsync binary to synchronize an origin database to a destination database.
// It assumes the sqlite3_rsync binary is present in the current working directory.
// This function blocks until the command completes.
func Sync(origin, destination string) error {

	binaryPath := "sqlite3_rsync"
	cmd := exec.Command(binaryPath, origin, destination)

	// Capture stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Execute the command
	err := cmd.Run()
	if err != nil {
		return errl.Errorf("sqlite3_rsync failed: %w, stderr: %s, stdout: %s", err, stderr.String(), stdout.String())
	}

	return nil
}

// Backup creates a backup of the origin database in a "backups" subdirectory.
// The backup filename includes the day of the week (1=Monday, ..., 7=Sunday).
func Backup(origin string) error {
	// Get absolute path of origin to safely manipulate directories
	absOrigin, err := filepath.Abs(origin)
	if err != nil {
		return errl.Errorf("failed to get absolute path of origin: %w", err)
	}

	dir := filepath.Dir(absOrigin)
	filename := filepath.Base(absOrigin)
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)

	// Create backups directory
	backupDir := filepath.Join(dir, "backups")
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return errl.Errorf("failed to create backup directory: %w", err)
	}

	// Determine day of week (1=Monday, ..., 7=Sunday)
	weekday := time.Now().Weekday()
	dayNum := int(weekday)
	if dayNum == 0 {
		dayNum = 7
	}

	// Construct destination filename
	backupFilename := fmt.Sprintf("%s_%d%s", name, dayNum, ext)
	destination := filepath.Join(backupDir, backupFilename)

	// Perform sync
	slog.Info("Performing backup", slog.String("origin", origin), slog.String("destination", destination))
	if err := Sync(origin, destination); err != nil {
		return errl.Errorf("failed to perform sync: %w", err)
	}

	return nil
}
