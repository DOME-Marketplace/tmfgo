package main

import (
	"flag" // Added
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"log/slog"

	"github.com/cloudflare/tableflip"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hesusruiz/tmforum/config"
	"github.com/hesusruiz/tmforum/internal/errl"
	"github.com/hesusruiz/tmforum/internal/sqlogger"
	_ "github.com/hesusruiz/tmforum/migrations"
	"github.com/hesusruiz/tmforum/pdp"
	fiberhandler "github.com/hesusruiz/tmforum/tmfserver/handler/fiber"
	repository "github.com/hesusruiz/tmforum/tmfserver/repository"
	service "github.com/hesusruiz/tmforum/tmfserver/service"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var debugFlag bool
	var init bool
	var environment string
	var restartHour, restartMinute int

	envHelp := fmt.Sprintf("Environment where run: %s, %s, %s, %s, %s, %s, %s", config.ISBE_DEV, config.ISBE_PRE, config.ISBE_PRO, config.DOME_DEV, config.DOME_PRE, config.DOME_PRO, config.LOCAL)

	// Parse command-line flags
	flag.BoolVar(&debugFlag, "d", false, "Enable debug logging")
	flag.BoolVar(&init, "init", false, "Run as init process")
	flag.StringVar(&environment, "run", "", envHelp)
	flag.IntVar(&restartHour, "rh", 3, "Restart program every day at this hour")
	flag.IntVar(&restartMinute, "rm", 0, "Restart program every day at this minute")
	flag.Parse()

	// Configure the slog logger
	var logLevel = new(slog.LevelVar)
	if debugFlag {
		logLevel.Set(slog.LevelDebug)
	} else {
		logLevel.Set(slog.LevelInfo)
	}

	// Initialize the custom SQLogHandler
	logOptions := &sqlogger.Options{
		Level:  logLevel,
		LogDir: "data/logs",
	}

	// Check if the logs should be colored:
	// - If the process is running in a container (pid=1) then do not color the logs
	// - If the environment variable ISBETMF_LOGS_NOCOLOR is set to "true" then do not color the logs
	ourPid := os.Getpid()
	if ourPid == 1 || os.Getenv("ISBETMF_LOGS_NOCOLOR") == "true" || os.Getenv("TMF_LOGS_NOCOLOR") == "true" {
		logOptions.NoColor = true
	}

	// Initialize the logging system
	sqlog, err := sqlogger.NewSQLogHandler(logOptions)
	if err != nil {
		slog.Error("failed to initialize SQLogHandler, exiting", slog.Any("error", err))
		os.Exit(1)
	}
	defer sqlog.Close()

	// And set the default logging system for all components
	slog.SetDefault(slog.New(sqlog))

	// Detect if we are running as PID=1 (an init process in a container),
	// and act accordingly.
	runAsInit := init || ourPid == 1

	if runAsInit {
		// Exclude the name of the program from the list of arguments
		args := os.Args[1:]
		runAsInitProcess(args)
	} else {
		slog.Info("We are the NORMAL process!", "environment", environment, "debug", debugFlag, "restartHour", restartHour, "restartMinute", restartMinute)

		// Generate a default configuration suitable for the environment
		configuration, err := config.LoadConfig(environment, debugFlag)
		if err != nil {
			slog.Error("Failed to load configuration", slog.Any("error", err))
			os.Exit(1)
		}
		slog.Info("Configuration loaded", "environment", configuration.Environment, "debug", configuration.Debug, "proxy", configuration.ProxyEnabled)

		// Set restart schedule
		configuration.RestartHour = restartHour
		configuration.RestartMinute = restartMinute

		err = runNormalProcess(configuration)
		if err != nil {
			slog.Error("failed to run normal process", slog.Any("error", err))
			os.Exit(1)
		}
	}

}

func cleanup(db *repository.DBService) {
	// This deferred function will run!
	fmt.Println("Running deferred cleanup functions...")

	// Close database connection (triggers WAL cleanup)
	fmt.Println("Closing database connection and exiting...")
	_ = db.Close()

	fmt.Println("Database connections closed.")
}

// runNormalProcess starts the TMF API server and handles its lifecycle,
// including database connection, rules engine initialization, and graceful shutdown.
func runNormalProcess(configuration *config.Config) error {

	// Set TABLEFLIP for seamless restarts and upgrades
	upg, err := tableflip.New(tableflip.Options{
		PIDFile: "isbetmf.pid",
	})
	if err != nil {
		return errl.Errorf("failed to create tableflip upgrader: %w", err)
	}
	defer upg.Stop()

	// Connect to the database and create tables if they do not exist
	dbService, err := repository.NewDBService(configuration.Dbname, configuration.ServerOperatorOrganizationIdentifier)
	if err != nil {
		return errl.Errorf("failed to connect to database: %w", err)
	}
	defer cleanup(dbService)

	// Create the PDP (aka Policy Decision Point or rules engine)
	rulesEngine, err := pdp.NewPDPService(&pdp.Config{
		PolicyFileName: configuration.PolicyFileName,
		Debug:          configuration.Debug,
	})
	if err != nil {
		return errl.Errorf("failed to create rules engine: %w", err)
	}

	// Create the service, which will use the database and the rules engine
	tmfService, err := service.NewTMFService(configuration, dbService, rulesEngine)
	if err != nil {
		return errl.Errorf("failed to create service: %w", err)
	}

	// Create Fiber web server with custom configuration
	webServer := fiber.New(fiber.Config{
		AppName:        "TMForum API Server",
		ServerHeader:   "TMForum",
		ProxyHeader:    "X-Forwarded-For",
		ReadBufferSize: 64 * 1024, // 64 KB — allows large Authorization headers (e.g. JWTs with many claims)
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			meth := fmt.Sprintf("<= %s %s", c.Method(), c.Path())
			slog.Error(meth, slog.Any("error", err), slog.Int("status", code), slog.String("ip", c.IP()))

			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Add middleware in proper order

	// Recovery middleware - should be first to catch panics
	webServer.Use(recover.New(recover.Config{
		EnableStackTrace: configuration.Debug,
	}))

	// Favicon middleware - serve favicon.ico
	webServer.Use(fiberhandler.Favicon)

	// Request ID middleware - for tracing requests
	webServer.Use(fiberhandler.RequestID)

	// CORS middleware - enable cross-origin requests
	webServer.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,X-Request-Id",
		AllowCredentials: false,
		ExposeHeaders:    "X-Request-Id",
		MaxAge:           86400,
	}))

	// Compression middleware - compress responses
	webServer.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// Serve the OpenAPI UI. We support V4 and V5
	webServer.Static("/oapiv5", "./www/oapiv5")
	webServer.Static("/oapiv4", "./www/oapiv4")
	webServer.Static("/assets", "./www/assets")

	// Logger middleware - log requests and replies
	webServer.Use(fiberhandler.FiberRequestLogger)

	// Create handler and set the routes for the APIs
	fiberhandler.NewHandler(webServer, tmfService)

	// Create and register admin handler
	fiberhandler.NewAdminHandler(webServer, tmfService)

	// Schedule periodic maintenance tasks
	repository.ScheduleMaintenance(configuration, dbService, upg)

	// For tableflip to work, Listen must be called before signaling we are ready
	ln, err := upg.Listen("tcp", "0.0.0.0:9991")
	if err != nil {
		slog.Error("failed to listen on port 9991, exiting", slog.Any("error", err))
		panic(err)
	}
	defer func() {
		_ = ln.Close()
	}()

	// Start the server in a separate goroutine
	go func() {
		slog.Info("TMF API server starting: http://localhost:9991/oapiv4/index.html")
		err := webServer.Listener(ln)
		if err != nil {
			slog.Error("Error starting TMF API server", "error", errl.Error(err))
			panic(err)
		}
	}()

	// Capture the termination signals to be able to perform a clean shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// Start the signal handler in a separate goroutine
	go func() {
		for sig := range sigs {
			// Process each type of signal
			switch sig {

			case syscall.SIGINT, syscall.SIGTERM:
				fmt.Println("CHILD: Received SIGINT or SIGTERM, exiting...")

				// Close listeners and inherited FDs
				// Call upg.Stop() to shut down listeners immediately
				upg.Stop()

			case syscall.SIGHUP:
				// Perform a Tableflip upgrade
				fmt.Println("CHILD: Received SIGHUP, upgrading...")
				_ = upg.Upgrade()
			}
		}
	}()

	// Signal that we are ready so Tableflip can stop the parent process
	slog.Info("CHILD: Server is ready")
	if err := upg.Ready(); err != nil {
		panic(errl.Error(err))
	}

	// Wait until we are told to exit by the Tableflip mechanism.
	// This happens when the child process has signalled that it is ready.
	fmt.Println("CHILD: Waiting for Tableflip to exit...")
	<-upg.Exit()
	fmt.Println("CHILD: Tableflip exit received")

	// Wait for connections to drain for a maximum of 30 seconds
	fmt.Println("CHILD: Waiting 30 seconds for connections to drain...")
	err = webServer.ShutdownWithTimeout(30 * time.Second)
	if err != nil {
		return errl.Errorf("failed to shutdown web server: %w", err)
	}
	fmt.Println("CHILD: Exiting without error")
	return nil

}

// runAsInitProcess acts as an init process, launching a child process and forwarding
// signals to it. This is typically used in container environments where
// `main` might be PID 1.
//
// It sets up the child process to share stdout/stderr, places it in the same
// process group, and captures system signals (SIGINT, SIGTERM, SIGHUP) to
// gracefully relay them to the child.
//
//   - args: Command-line arguments to pass to the child process.
func runAsInitProcess(args []string) {

	ourPid := os.Getpid()

	// Get the name of our executable, to be able to restart it automatically
	ourExecPath, err := os.Executable()
	if err != nil {
		slog.Error("Failed to get executable path", slog.Any("error", err))
		panic(err)
	}

	slog.Info("We are the INIT process!", "PID", ourPid, "executable", ourExecPath, "args", args)

	// Pass to child all arguments except the "-init" flag, so the child runs as a normal process.
	childArgs := make([]string, 0, len(args))
	for _, a := range args {
		if a != "-init" && a != "--init" {
			childArgs = append(childArgs, a)
		}
	}
	cmd := exec.Command(ourExecPath, childArgs...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Set ProcessGroupID for child process as init process. Both will be under same process group
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// Start the child (a fork of ourselves) without waiting for termination
	slog.Info("INIT: starting child process")
	if err := cmd.Start(); err != nil {
		slog.Error("INIT: failed to start child process", "error", err)
		return
	}

	// We need notification of all relevant signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	done := make(chan bool, 1)

	// Forward all signals to the child in a goroutine
	go func() {
		for sig := range sigs {

			// Forward the signal to the child process
			slog.Info("INIT: forwarding signal to child process", "signal", sig, "PID", cmd.Process.Pid)
			err := cmd.Process.Signal(sig)
			if err != nil {
				slog.Error("INIT: failed to forward signal to child process", "signal", sig, "PID", cmd.Process.Pid, "error", err)
			}

			// If the signal was SIGTERM or SIGINT, wait 10 seconds for the child to terminate and send a KILL signal
			if sig == syscall.SIGTERM || sig == syscall.SIGINT {

				go func() {
					// Wait 10 seconds for the child process to finish
					time.Sleep(10 * time.Second)
					// Kill the child immediately
					_ = cmd.Process.Kill()
				}()

				slog.Info("INIT: using DONE channel to terminate init process")
				done <- true
			}
		}
	}()

	// Enter in a goroutine an infinite loop reaping periodically the zombie children
	go func() {
		for {
			var ws syscall.WaitStatus
			pid, err := syscall.Wait4(-1, &ws, syscall.WNOHANG, nil)
			if pid <= 0 || err != nil {
				time.Sleep(5 * time.Second)
			} else {
				slog.Info("INIT: reaped zombie child with PID", "PID", pid)
			}
		}

	}()

	slog.Info("INIT: awaiting signal to terminate init process")
	<-done

	// Wait for the child process to finish and release its resources
	slog.Info("INIT: waiting for child process to finish")
	_, _ = cmd.Process.Wait()

	slog.Info("INIT: exiting init process")
}
