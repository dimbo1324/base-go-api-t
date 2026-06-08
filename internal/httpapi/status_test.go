package httpapi

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dimbo1324/Base-Go-API-Engine/internal/config"
	"github.com/dimbo1324/Base-Go-API-Engine/internal/store"
)

func TestStatusCheck(t *testing.T) {
	app := newTestApplication()
	req := httptest.NewRequest(http.MethodGet, "/v1/status", nil)
	rec := httptest.NewRecorder()

	app.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status code = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("Content-Type = %q, want application/json", got)
	}
	if got := rec.Body.String(); got != "{\"status\":\"ok\"}\n" {
		t.Fatalf("body = %q, want status response", got)
	}
}

func TestStatusCheckRejectsUnsupportedMethod(t *testing.T) {
	app := newTestApplication()
	req := httptest.NewRequest(http.MethodPost, "/v1/status", nil)
	rec := httptest.NewRecorder()

	app.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status code = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
	}
	if got := rec.Header().Get("Allow"); got != http.MethodGet {
		t.Fatalf("Allow = %q, want GET", got)
	}
}

func TestNotFound(t *testing.T) {
	app := newTestApplication()
	req := httptest.NewRequest(http.MethodGet, "/missing", nil)
	rec := httptest.NewRecorder()

	app.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status code = %d, want %d", rec.Code, http.StatusNotFound)
	}
	if got := rec.Body.String(); got != "{\"error\":\"not found\"}\n" {
		t.Fatalf("body = %q, want not found response", got)
	}
}

func newTestApplication() *Application {
	cfg := config.Config{
		Server: config.ServerConfig{
			Addr:              ":0",
			RequestTimeout:    time.Second,
			ReadHeaderTimeout: time.Second,
			ReadTimeout:       time.Second,
			WriteTimeout:      time.Second,
			IdleTimeout:       time.Second,
			ShutdownTimeout:   time.Second,
		},
	}

	return NewApplication(cfg, store.Storage{}, log.New(io.Discard, "", 0))
}
