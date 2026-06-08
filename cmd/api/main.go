package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dimbo1324/Base-Go-API-Engine/internal/config"
	"github.com/dimbo1324/Base-Go-API-Engine/internal/db"
	"github.com/dimbo1324/Base-Go-API-Engine/internal/httpapi"
	"github.com/dimbo1324/Base-Go-API-Engine/internal/store"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.LUTC)

	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("load config: %v", err)
	}

	startupCtx, cancelStartup := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelStartup()

	dbConn, err := db.New(startupCtx, cfg.DB.Addr, cfg.DB.MaxOpenConns, cfg.DB.MaxIdleConns, cfg.DB.MaxIdleTime)
	if err != nil {
		logger.Fatalf("connect database: %v", err)
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			logger.Printf("close database: %v", err)
		}
	}()

	storage := store.NewStorage(dbConn)
	app := httpapi.NewApplication(cfg, storage, logger)

	runCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := app.Run(runCtx); err != nil {
		logger.Fatalf("run server: %v", err)
	}
}
