package config

import (
	"testing"
	"time"
)

func TestLoadUsesDefaults(t *testing.T) {
	t.Setenv(EnvAddr, "")
	t.Setenv(EnvDBAddr, "")
	t.Setenv(EnvDBMaxOpenConns, "")
	t.Setenv(EnvDBMaxIdleConns, "")
	t.Setenv(EnvDBMaxIdleTime, "")
	t.Setenv(legacyEnvDBMaxIdleTime, "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.Server.Addr != defaultAddr {
		t.Fatalf("Server.Addr = %q, want %q", cfg.Server.Addr, defaultAddr)
	}
	if cfg.DB.MaxOpenConns != defaultDBMaxOpenConns {
		t.Fatalf("DB.MaxOpenConns = %d, want %d", cfg.DB.MaxOpenConns, defaultDBMaxOpenConns)
	}
	if cfg.DB.MaxIdleTime != defaultDBMaxIdleTime {
		t.Fatalf("DB.MaxIdleTime = %s, want %s", cfg.DB.MaxIdleTime, defaultDBMaxIdleTime)
	}
}

func TestLoadReadsEnvironment(t *testing.T) {
	t.Setenv(EnvAddr, ":9090")
	t.Setenv(EnvDBAddr, "postgres://example")
	t.Setenv(EnvDBMaxOpenConns, "11")
	t.Setenv(EnvDBMaxIdleConns, "7")
	t.Setenv(EnvDBMaxIdleTime, "20m")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.Server.Addr != ":9090" {
		t.Fatalf("Server.Addr = %q, want :9090", cfg.Server.Addr)
	}
	if cfg.DB.Addr != "postgres://example" {
		t.Fatalf("DB.Addr = %q, want postgres://example", cfg.DB.Addr)
	}
	if cfg.DB.MaxOpenConns != 11 {
		t.Fatalf("DB.MaxOpenConns = %d, want 11", cfg.DB.MaxOpenConns)
	}
	if cfg.DB.MaxIdleConns != 7 {
		t.Fatalf("DB.MaxIdleConns = %d, want 7", cfg.DB.MaxIdleConns)
	}
	if cfg.DB.MaxIdleTime != 20*time.Minute {
		t.Fatalf("DB.MaxIdleTime = %s, want 20m", cfg.DB.MaxIdleTime)
	}
}

func TestLoadRejectsInvalidInteger(t *testing.T) {
	t.Setenv(EnvDBMaxOpenConns, "bad")

	if _, err := Load(); err == nil {
		t.Fatal("Load() error = nil, want error")
	}
}

func TestLoadRejectsInvalidDuration(t *testing.T) {
	t.Setenv(EnvDBMaxIdleTime, "bad")

	if _, err := Load(); err == nil {
		t.Fatal("Load() error = nil, want error")
	}
}
