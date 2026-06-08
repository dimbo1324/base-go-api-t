package httpapi

import (
	"log"
	"os"

	"github.com/dimbo1324/Base-Go-API-Engine/internal/config"
	"github.com/dimbo1324/Base-Go-API-Engine/internal/store"
)

type Application struct {
	config config.Config
	store  store.Storage
	logger *log.Logger
}

func NewApplication(cfg config.Config, storage store.Storage, logger *log.Logger) *Application {
	if logger == nil {
		logger = log.New(os.Stdout, "", log.LstdFlags|log.LUTC)
	}

	return &Application{
		config: cfg,
		store:  storage,
		logger: logger,
	}
}
