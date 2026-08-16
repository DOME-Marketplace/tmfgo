package fiber

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/gofiber/fiber/v2"
	"github.com/hesusruiz/tmforum/config"
	"github.com/hesusruiz/tmforum/internal/html"
	"github.com/hesusruiz/tmforum/internal/sqlogger"
	"github.com/hesusruiz/tmforum/tmfserver/service"
)

//go:embed templates/*
var templatesFS embed.FS

const adminSessionCookieName = "admin_session"

type AdminHandler struct {
	service      *service.Service
	htmlRenderer *html.Renderer
}

func NewAdminHandler(app *fiber.App, s *service.Service) *AdminHandler {

	htmlRenderer, err := html.NewRenderer(true, &templatesFS, "templates", "tmfserver/handler/fiber/templates", ".html")
	if err != nil {
		panic(fmt.Errorf("failed to create admin templates renderer: %w", err))
	}

	h := &AdminHandler{
		service:      s,
		htmlRenderer: htmlRenderer,
	}

	h.registerRoutes(app)

	return h

}

func (h *AdminHandler) registerRoutes(app *fiber.App) {
	admin := app.Group("/admin")

	admin.Get("/login", h.ShowLogin)
	admin.Post("/login", h.ProcessLogin)
	admin.Get("/logout", h.Logout)

	admin.Use(h.RequireAuth)

	admin.Get("/", h.Dashboard)

	admin.Get("/page/settings", h.Settings)
	admin.Post("/page/settings", h.Settings)

	admin.Get("/page/upstream", h.Upstream)
	admin.Post("/page/upstream", h.Upstream)

	admin.Get("/:resourceName", h.ListObjects)
	admin.Get("/:resourceName/:id", h.ViewObject)
}

func (h *AdminHandler) Dashboard(c *fiber.Ctx) error {
	var level slog.Level
	logger := slog.Default()
	mylogger, ok := logger.Handler().(*sqlogger.SQLogHandler)
	if ok {
		leveler := mylogger.Level()
		level = leveler.Level()
		fmt.Printf("Current log level: %d\n", leveler.Level())
	}

	data := map[string]any{
		"settings": "active",
		"service":  h.service,
		"logLevel": level,
	}
	fmt.Println("Level:", level)
	return h.render(c, "index", data)
}

func (h *AdminHandler) Settings(c *fiber.Ctx) error {

	pageName := "settings"

	pageData := map[string]any{
		pageName:  "active",
		"Service": h.service,
	}

	switch c.Method() {
	case http.MethodGet:
		var level slog.Level
		logger := slog.Default()
		mylogger, ok := logger.Handler().(*sqlogger.SQLogHandler)
		if ok {
			leveler := mylogger.Level()
			level = leveler.Level()
			fmt.Printf("Current log level: %d\n", leveler.Level())
		}

		pageData["LogLevel"] = level
		if c.Query("ok") == "1" {
			pageData["Success"] = true
		}
		return h.render(c, pageName, pageData)

	case http.MethodPost:
		newLogLevelValue := c.FormValue("logLevel")
		if newLogLevelValue == "" {
			slog.Warn("No log level provided")
			pageData["Error"] = "No log level provided"
			return h.render(c, pageName, pageData)
		}

		var level slog.Level
		switch newLogLevelValue {
		case "DEBUG":
			level = slog.LevelDebug
		case "INFO":
			level = slog.LevelInfo
		case "WARN":
			level = slog.LevelWarn
		case "ERROR":
			level = slog.LevelError
		default:
			slog.Warn("Invalid log level provided: " + newLogLevelValue)
			pageData["Error"] = "Invalid log level provided: " + newLogLevelValue
			return h.render(c, pageName, pageData)
		}

		logger := slog.Default()
		mylogger, ok := logger.Handler().(*sqlogger.SQLogHandler)
		if ok {
			leveler := mylogger.Level()
			fmt.Printf("Current log level: %d\n", leveler.Level())
			leveler.Set(level)
			fmt.Printf("New log level: %d\n", level)
		}

		return c.Redirect("/admin/page/" + pageName + "?ok=1")

	default:
		return c.Redirect("/admin")
	}

}

func (h *AdminHandler) Upstream(c *fiber.Ctx) error {

	pageName := "upstream"

	pageData := map[string]any{
		pageName:  "active",
		"Service": h.service,
	}

	switch c.Method() {
	case http.MethodGet:
		if c.Query("ok") == "1" {
			pageData["Success"] = true
		}
		if c.Query("error") == "1" {
			pageData["Error"] = "Error uploading file"
		}

		proxyConfig := config.GetProxyConfig()
		if proxyConfig == nil {
			pageData["Error"] = "No upstream config found"
			return h.render(c, pageName, pageData)
		}

		upstreamEntries := proxyConfig.GetUpstreamEntries()
		if len(upstreamEntries) == 0 {
			pageData["Error"] = "No upstream entries found"
			return h.render(c, pageName, pageData)
		}

		yamlBytes, err := yaml.Marshal(upstreamEntries)
		if err != nil {
			pageData["Error"] = "Error marshaling upstream config"
			return h.render(c, pageName, pageData)
		}

		pageData["UpstreamEntries"] = string(yamlBytes)

		return h.render(c, pageName, pageData)

	case http.MethodPost:
		// Get first fileHeader from form field "document":
		fileHeader, err := c.FormFile("file")
		if fileHeader == nil {
			fileHeader, err = c.FormFile("file[0]")
		}
		if err != nil {
			slog.Error("Error retrieving the file", "filename", fileHeader.Filename, "error", err)
			return c.Status(fiber.StatusBadRequest).SendString("File not specified in the request")
		}
		slog.Info("Uploading new TMF configuration", "filename", fileHeader.Filename, "size", fileHeader.Size)

		// Do not load files bigger than 5MB
		if fileHeader.Size > 5*1024*1024 {
			slog.Error("File too big", "filename", fileHeader.Filename, "size", fileHeader.Size)
			return c.Status(fiber.StatusBadRequest).SendString("File too big")
		}

		// Read in memory the file
		f, err := fileHeader.Open()
		if err != nil {
			slog.Error("Error opening the file", "filename", fileHeader.Filename, "error", err)
			return c.Status(fiber.StatusBadRequest).SendString("Error opening the file")
		}
		defer func() {
			_ = f.Close()
		}()

		// Read file to buffer
		buffer := make([]byte, fileHeader.Size)
		_, err = f.Read(buffer)
		if err != nil {
			slog.Error("Error reading the file", "filename", fileHeader.Filename, "error", err)
			return c.Status(fiber.StatusBadRequest).SendString("Error reading the file")
		}

		// Parse the contents into the proxy config
		var newUpstreamEntries config.UpstreamEntries
		err = yaml.Unmarshal(buffer, &newUpstreamEntries)
		if err != nil {
			slog.Error("Error parsing the file", "filename", fileHeader.Filename, "error", err)
			return c.Status(fiber.StatusBadRequest).SendString("Error parsing the file: " + err.Error())
		}

		// Make sure that we have at least one routing entry
		if len(newUpstreamEntries) == 0 {
			slog.Error("No routing entries found in the file", "filename", fileHeader.Filename)
			return c.Status(fiber.StatusBadRequest).SendString("No routing entries found in the file")
		}

		// Loop the entries to check that host, port and path are valid
		for _, entry := range newUpstreamEntries {
			if entry.Host == "" {
				slog.Error("Invalid host in the file", "filename", fileHeader.Filename)
				return c.Status(fiber.StatusBadRequest).SendString("Invalid host in the file")
			}
			if entry.Port == 0 {
				slog.Error("Invalid port in the file", "filename", fileHeader.Filename)
				return c.Status(fiber.StatusBadRequest).SendString("Invalid port in the file")
			}
			if entry.Path == "" {
				slog.Error("Invalid path in the file", "filename", fileHeader.Filename)
				return c.Status(fiber.StatusBadRequest).SendString("Invalid path in the file")
			}
		}

		// Update the global proxy config
		config.UpdateProxyConfig(newUpstreamEntries)

		return c.Status(fiber.StatusOK).SendString("New TMF configuration uploaded successfully")

	default:
		return c.Redirect("/admin")
	}

}

func (h *AdminHandler) RequireAuth(c *fiber.Ctx) error {
	path := c.Path()
	if path == "/admin/login" || path == "/admin/logout" {
		return c.Next()
	}

	cookie := c.Cookies(adminSessionCookieName)
	if cookie == "" || !h.isValidSession(cookie) {
		return c.Redirect("/admin/login")
	}

	return c.Next()
}

func (h *AdminHandler) ShowLogin(c *fiber.Ctx) error {
	cookie := c.Cookies(adminSessionCookieName)
	if cookie != "" && h.isValidSession(cookie) {
		return c.Redirect("/admin/")
	}
	return h.render(c, "login", map[string]any{})
}

func (h *AdminHandler) ProcessLogin(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	adminToken := h.service.AdminToken()
	if adminToken == "" || password != adminToken {
		data := map[string]any{
			"Error": "Invalid password",
			"Email": email,
		}
		return h.render(c, "login", data)
	}

	sessionToken := h.createSessionToken()
	c.Cookie(&fiber.Cookie{
		Name:     adminSessionCookieName,
		Value:    sessionToken,
		Path:     "/admin",
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		SameSite: "Lax",
	})

	return c.Redirect("/admin/")
}

func (h *AdminHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     adminSessionCookieName,
		Value:    "",
		Path:     "/admin",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		SameSite: "Lax",
	})
	return c.Redirect("/admin/login")
}

func (h *AdminHandler) isValidSession(tokenStr string) bool {
	adminToken := h.service.AdminToken()
	if adminToken == "" {
		return false
	}

	parts := strings.Split(tokenStr, ".")
	if len(parts) != 2 {
		return false
	}

	expStr, sig := parts[0], parts[1]
	expUnix, err := strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		return false
	}

	if time.Now().Unix() > expUnix {
		return false
	}

	expectedSig := h.computeSessionHMAC(adminToken, expStr)
	return hmac.Equal([]byte(sig), []byte(expectedSig))
}

func (h *AdminHandler) createSessionToken() string {
	adminToken := h.service.AdminToken()
	exp := time.Now().Add(24 * time.Hour).Unix()
	expStr := strconv.FormatInt(exp, 10)
	sig := h.computeSessionHMAC(adminToken, expStr)
	return fmt.Sprintf("%s.%s", expStr, sig)
}

func (h *AdminHandler) computeSessionHMAC(adminToken string, expStr string) string {
	mac := hmac.New(sha256.New, []byte(adminToken))
	mac.Write([]byte("admin_session:" + expStr))
	return hex.EncodeToString(mac.Sum(nil))
}

func (h *AdminHandler) ListObjects(c *fiber.Ctx) error {
	resourceName := c.Params("resourceName")

	// Default to v4 for now, could be made configurable or inferred
	apiFamily := "productCatalogManagement"
	switch resourceName {
	case "agreement":
		apiFamily = "agreementManagement"
	case "individual", "organization":
		apiFamily = "party"
	}

	req := &service.Request{
		Method:       "GET",
		Action:       service.ActionLIST,
		APIfamily:    apiFamily,
		APIVersion:   "v4",
		ResourceName: resourceName,
		QueryParams:  url.Values{"limit": []string{"20"}}, // Default limit
	}

	// Create a context with a timeout of 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp := h.service.ListTMFObjects(ctx, req)

	if resp.StatusCode >= 400 {
		return c.Status(resp.StatusCode).SendString(fmt.Sprintf("Error fetching objects: %v", resp.Body))
	}

	data := map[string]any{
		"ResourceName": resourceName,
		"Objects":      resp.Body,
	}

	return h.render(c, "list", data)
}

func (h *AdminHandler) ViewObject(c *fiber.Ctx) error {
	resourceName := c.Params("resourceName")
	id := c.Params("id")

	apiFamily := "productCatalogManagement"
	switch resourceName {
	case "agreement":
		apiFamily = "agreementManagement"
	case "individual", "organization":
		apiFamily = "party"
	}

	req := &service.Request{
		Method:       "GET",
		Action:       service.ActionREAD,
		APIfamily:    apiFamily,
		APIVersion:   "v4",
		ResourceName: resourceName,
		ID:           id,
	}

	// Create a context with a timeout of 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp := h.service.GetTMFObject(ctx, req)

	if resp.StatusCode >= 400 {
		return c.Status(resp.StatusCode).SendString(fmt.Sprintf("Error fetching object: %v", resp.Body))
	}

	// Pretty print JSON
	jsonBytes, _ := json.MarshalIndent(resp.Body, "", "  ")

	data := map[string]any{
		"ResourceName": resourceName,
		"ID":           id,
		"Object":       resp.Body,
		"JSON":         string(jsonBytes),
	}

	return h.render(c, "detail", data)
}

func (h *AdminHandler) render(c *fiber.Ctx, name string, data map[string]any) error {
	c.Set("Content-Type", "text/html")
	return h.htmlRenderer.RenderFiber(c, name, data)
}
