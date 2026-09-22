package cmd

import (
	"os"

	"github.com/DOME-Marketplace/tmfgo/internal/supervisor"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the TM Forum API server",
	Long:  `Starts the TMF API server, initializes the database, PDP rules engine, and serves requests.`,
	RunE: func(c *cobra.Command, args []string) error {
		ourPid := os.Getpid()
		if initFlag || ourPid == 1 {
			supervisor.Run(os.Args[1:])
			return nil
		}

		return StartServer(environment, debugFlag, restartHour, restartMinute, deleteInvalid)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
