package cmd

import (
	"context"
	"log/slog"

	"github.com/DOME-Marketplace/tmfgo/config"
	"github.com/DOME-Marketplace/tmfgo/internal/errl"
	"github.com/DOME-Marketplace/tmfgo/pdp"
	repository "github.com/DOME-Marketplace/tmfgo/tmfserver/repository"
	service "github.com/DOME-Marketplace/tmfgo/tmfserver/service"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Delete invalid objects from the database",
	Long:  `Connects to the database and purges invalid or corrupted TMF objects.`,
	RunE: func(c *cobra.Command, args []string) error {
		configuration, err := config.LoadConfig(environment, debugFlag)
		if err != nil {
			return errl.Errorf("failed to load configuration: %w", err)
		}

		dbService, err := repository.NewDBService(configuration.Dbname, configuration.ServerOperatorOrganizationIdentifier)
		if err != nil {
			return errl.Errorf("failed to connect to database: %w", err)
		}
		defer func() {
			_ = dbService.Close()
		}()

		rulesEngine, err := pdp.NewPDPService(&pdp.Config{
			PolicyFileName: configuration.PolicyFileName,
			Debug:          configuration.Debug,
		})
		if err != nil {
			return errl.Errorf("failed to create rules engine: %w", err)
		}

		tmfService, err := service.NewTMFService(configuration, dbService, rulesEngine)
		if err != nil {
			return errl.Errorf("failed to create service: %w", err)
		}

		slog.Info("Executing object purge...")
		if err := tmfService.RetrieveAll(context.Background(), true); err != nil {
			return errl.Errorf("purge failed: %w", err)
		}

		slog.Info("Purge completed successfully")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)
}
