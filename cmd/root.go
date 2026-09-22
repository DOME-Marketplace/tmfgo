package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/DOME-Marketplace/tmfgo/config"
	"github.com/DOME-Marketplace/tmfgo/internal/sqlogger"
	"github.com/DOME-Marketplace/tmfgo/internal/supervisor"
	"github.com/DOME-Marketplace/tmfgo/tmfserver"
	"github.com/spf13/cobra"
)

var (
	debugFlag      bool
	initFlag       bool
	environment    string
	restartHour    int
	restartMinute  int
	deleteInvalid  bool

	sqlogCloseFn func()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tmfgo",
	Short: "TM Forum Open API server and proxy",
	Long: `tmfgo is a TM Forum (TMF) Open API server and proxy written in Go.
It can operate as a standalone TMF server or as an authenticating and
authorizing proxy in front of an upstream TMF instance.`,
	// When executed with no subcommands (e.g. `tmfgo -run local`), run the server for backward compatibility
	RunE: func(c *cobra.Command, args []string) error {
		// Detect if we are running as PID=1 (an init process in a container)
		ourPid := os.Getpid()
		if initFlag || ourPid == 1 {
			supervisor.Run(os.Args[1:])
			return nil
		}

		return StartServer(environment, debugFlag, restartHour, restartMinute, deleteInvalid)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if sqlogCloseFn != nil {
		sqlogCloseFn()
	}
}

func init() {
	cobra.OnInitialize(initLogging)

	envHelp := fmt.Sprintf("Environment where run: %s, %s, %s, %s, %s, %s, %s",
		config.ISBE_DEV, config.ISBE_PRE, config.ISBE_PRO,
		config.DOME_DEV, config.DOME_PRE, config.DOME_PRO, config.LOCAL)

	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "d", false, "Enable debug logging")
	rootCmd.PersistentFlags().BoolVar(&initFlag, "init", false, "Run as container init process")
	rootCmd.PersistentFlags().StringVar(&environment, "run", "", envHelp)
	rootCmd.PersistentFlags().IntVar(&restartHour, "rh", 3, "Restart program every day at this hour")
	rootCmd.PersistentFlags().IntVar(&restartMinute, "rm", 0, "Restart program every day at this minute")
	rootCmd.PersistentFlags().BoolVar(&deleteInvalid, "purge", false, "Delete invalid objects on startup")
}

func initLogging() {
	var logLevel = new(slog.LevelVar)
	if debugFlag {
		logLevel.Set(slog.LevelDebug)
	} else {
		logLevel.Set(slog.LevelInfo)
	}

	logOptions := &sqlogger.Options{
		Level:  logLevel,
		LogDir: "data/logs",
	}

	ourPid := os.Getpid()
	if ourPid == 1 || os.Getenv("ISBETMF_LOGS_NOCOLOR") == "true" || os.Getenv("TMF_LOGS_NOCOLOR") == "true" {
		logOptions.NoColor = true
	}

	sqlog, err := sqlogger.NewSQLogHandler(logOptions)
	if err != nil {
		slog.Error("failed to initialize SQLogHandler, exiting", slog.Any("error", err))
		os.Exit(1)
	}

	sqlogCloseFn = func() {
		sqlog.Close()
	}

	slog.SetDefault(slog.New(sqlog))
}

// StartServer loads configuration and starts the TMF server
func StartServer(env string, debug bool, rHour, rMinute int, purge bool) error {
	slog.Info("Starting TMF server", "environment", env, "debug", debug, "restartHour", rHour, "restartMinute", rMinute)

	configuration, err := config.LoadConfig(env, debug)
	if err != nil {
		slog.Error("Failed to load configuration", slog.Any("error", err))
		return err
	}
	slog.Info("Configuration loaded", "environment", configuration.Environment, "debug", configuration.Debug, "proxy", configuration.ProxyEnabled)

	configuration.RestartHour = rHour
	configuration.RestartMinute = rMinute

	return tmfserver.Run(configuration, purge)
}
