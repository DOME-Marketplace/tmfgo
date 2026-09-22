package tmfserver

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DOME-Marketplace/tmfgo/config"
	"github.com/DOME-Marketplace/tmfgo/internal/errl"
	_ "github.com/DOME-Marketplace/tmfgo/migrations"
	"github.com/DOME-Marketplace/tmfgo/pdp"
	"github.com/DOME-Marketplace/tmfgo/tmfserver/admin"
	fiberhandler "github.com/DOME-Marketplace/tmfgo/tmfserver/handler/fiber"
	repository "github.com/DOME-Marketplace/tmfgo/tmfserver/repository"
	service "github.com/DOME-Marketplace/tmfgo/tmfserver/service"
	"github.com/cloudflare/tableflip"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/mattn/go-sqlite3"
)

func cleanup(db *repository.DBService) {
	fmt.Println("Running deferred cleanup functions...")
	fmt.Println("Closing database connection and exiting...")
	_ = db.Close()
	fmt.Println("Database connections closed.")
}

// Run starts the TMF API server and handles its lifecycle,
// including database connection, rules engine initialization, Fiber HTTP engine,
// tableflip upgrades, and graceful shutdown.
func Run(configuration *config.Config, deleteInvalid bool) error {

	// Set TABLEFLIP for seamless restarts and upgrades
	upg, err := tableflip.New(tableflip.Options{
		PIDFile: "tmfgo.pid",
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

	if deleteInvalid {
		err = tmfService.RetrieveAll(context.Background(), deleteInvalid)
		if err != nil {
			return errl.Errorf("failed to retrieve all objects: %w", err)
		}
		return nil
	}

	// Schedule retrieve all when proxy mode is enabled
	tmfService.ScheduleRetrieveAll()

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
	admin.NewHandler(webServer, tmfService)

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
