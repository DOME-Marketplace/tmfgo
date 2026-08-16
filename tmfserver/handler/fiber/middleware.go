package fiber

import (
	"context"
	"fmt"
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/gofiber/fiber/v2"
)

const HeaderXRequestID = "X-Request-ID"
const ContextKeyRequestID = "requestid"

var requestCounter atomic.Uint64

// RequestID is a middleware that generates a request ID.
// The generated request ID is simply a prefixed request counter.
func RequestID(c *fiber.Ctx) error {

	// Get id from request, else we generate one
	rid := c.Get(HeaderXRequestID)
	if rid == "" {
		rid = fmt.Sprintf("TMFGo-%d", requestCounter.Add(1))
	}

	// Set new id to response header
	c.Set(HeaderXRequestID, rid)

	// Add the request ID to locals
	c.Locals(ContextKeyRequestID, rid)

	// Next handler will take care of the request
	return c.Next()
}

var noLoggingFor = map[string]bool{
	"/health":      true,
	"/favicon.ico": true,
}

func isPathLoggable(path string) bool {
	_, found := noLoggingFor[path]
	return !found
}

// FiberRequestLogger logs HTTP requests on entry and exit
func FiberRequestLogger(c *fiber.Ctx) error {

	reqId := c.Locals("requestid").(string)

	// Log entry, except the /health request, to keep logs clean
	if isPathLoggable(c.Path()) {
		slog.Info("=> "+c.Method()+" "+c.Path(), slog.String("request_id", reqId), slog.String("ip", c.IP()))
	}

	// Go to next middleware, measuring elapsed time
	start := time.Now()
	err := c.Next()
	end := time.Now()
	latency := end.Sub(start)

	// Log the requests that received an error
	if err != nil {
		slog.Error("<= "+c.Method()+" "+c.Path()+" UNKNOWN ERROR",
			slog.Any("error", err),
			slog.String("request_id", reqId),
			slog.String("ip", c.IP()),
			slog.Duration("latency", latency))
		return err
	}

	// Log with different logging levels depending on the status code
	code := c.Response().StatusCode()

	logLevel := slog.LevelInfo
	if code >= 400 {
		logLevel = slog.LevelError
	}

	// Log all replies, except those in the noLoggingFor map
	if isPathLoggable(c.Path()) {
		methodAndPath := fmt.Sprintf("<= %s %d %s", c.Method(), code, c.Path())
		slog.LogAttrs(context.Background(),
			logLevel,
			methodAndPath,
			slog.Int("status", code),
			slog.String("request_id", reqId),
			slog.String("ip", c.IP()),
			slog.Duration("latency", latency))
	}

	// Add a specific warning log entry if the elapsed time is bigger than 10 secs, which is symptom of a problem
	const slowThreshold = 10 * time.Second
	if latency > slowThreshold {
		slog.Warn("=> Slow request detected", slog.String("request_id", reqId), slog.Duration("latency", latency))
	}

	return nil

}
