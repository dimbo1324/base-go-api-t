package httpapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

func (app *Application) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:              app.config.Server.Addr,
		Handler:           app.Routes(),
		ReadHeaderTimeout: app.config.Server.ReadHeaderTimeout,
		ReadTimeout:       app.config.Server.ReadTimeout,
		WriteTimeout:      app.config.Server.WriteTimeout,
		IdleTimeout:       app.config.Server.IdleTimeout,
	}

	errCh := make(chan error, 1)
	go func() {
		app.logger.Printf("server started on %s", app.config.Server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
			return
		}
		errCh <- nil
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), app.config.Server.ShutdownTimeout)
		defer cancel()

		app.logger.Print("server shutdown started")
		if err := server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("shutdown server: %w", err)
		}

		if err := <-errCh; err != nil {
			return err
		}

		app.logger.Print("server stopped")
		return nil
	}
}
