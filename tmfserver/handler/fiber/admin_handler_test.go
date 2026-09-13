package fiber

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/DOME-Marketplace/tmfgo/config"
	"github.com/DOME-Marketplace/tmfgo/tmfserver/service"
	"github.com/gofiber/fiber/v2"
)

func createTestAdminApp(adminToken string) (*fiber.App, *AdminHandler) {
	app := fiber.New()
	cnf := &config.Config{
		AdminToken:     adminToken,
		VerifierServer: "https://verifier.dome-marketplace-sbx.org",
	}
	s, _ := service.NewTMFService(cnf, nil, nil)
	adminHandler := NewAdminHandler(app, s)
	return app, adminHandler
}

func TestAdminRoutesProtection(t *testing.T) {
	adminToken := "secretAdminPass123"
	app, _ := createTestAdminApp(adminToken)

	// 1. Unauthenticated access to /admin/ should redirect to /admin/login
	req := httptest.NewRequest("GET", "/admin/", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("failed test request: %v", err)
	}
	if resp.StatusCode != http.StatusFound {
		t.Errorf("expected 302 Found, got %d", resp.StatusCode)
	}
	if loc := resp.Header.Get("Location"); loc != "/admin/login" {
		t.Errorf("expected redirect location /admin/login, got %s", loc)
	}

	// 2. GET /admin/login should serve login page with 200 OK
	req = httptest.NewRequest("GET", "/admin/login", nil)
	resp, err = app.Test(req)
	if err != nil {
		t.Fatalf("failed test request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 OK for /admin/login, got %d", resp.StatusCode)
	}

	// 3. POST /admin/login with incorrect password should fail and return 200 with error
	form := url.Values{}
	form.Set("email", "admin@example.com")
	form.Set("password", "wrongpassword")
	req = httptest.NewRequest("POST", "/admin/login", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err = app.Test(req)
	if err != nil {
		t.Fatalf("failed test request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", resp.StatusCode)
	}

	// 4. POST /admin/login with correct password should succeed and return redirect to /admin/ with session cookie
	form = url.Values{}
	form.Set("email", "admin@example.com")
	form.Set("password", adminToken)
	req = httptest.NewRequest("POST", "/admin/login", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err = app.Test(req)
	if err != nil {
		t.Fatalf("failed test request: %v", err)
	}
	if resp.StatusCode != http.StatusFound {
		t.Errorf("expected 302 Found, got %d", resp.StatusCode)
	}
	if loc := resp.Header.Get("Location"); loc != "/admin/" {
		t.Errorf("expected redirect to /admin/, got %s", loc)
	}

	cookies := resp.Cookies()
	var sessionCookie *http.Cookie
	for _, c := range cookies {
		if c.Name == adminSessionCookieName {
			sessionCookie = c
			break
		}
	}
	if sessionCookie == nil || sessionCookie.Value == "" {
		t.Fatalf("expected admin_session cookie to be set")
	}

	// 5. Accessing /admin/ with valid session cookie should return 200 OK
	req = httptest.NewRequest("GET", "/admin/", nil)
	req.AddCookie(sessionCookie)
	resp, err = app.Test(req)
	if err != nil {
		t.Fatalf("failed test request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 OK with valid session, got %d", resp.StatusCode)
	}

	// 6. GET /admin/logout should clear cookie and redirect to /admin/login
	req = httptest.NewRequest("GET", "/admin/logout", nil)
	req.AddCookie(sessionCookie)
	resp, err = app.Test(req)
	if err != nil {
		t.Fatalf("failed test request: %v", err)
	}
	if resp.StatusCode != http.StatusFound {
		t.Errorf("expected 302 Found, got %d", resp.StatusCode)
	}
	if loc := resp.Header.Get("Location"); loc != "/admin/login" {
		t.Errorf("expected redirect location /admin/login, got %s", loc)
	}
}
