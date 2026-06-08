package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func New(ctx context.Context, addr string, maxOpenConns, maxIdleConns int, maxIdleTime time.Duration) (*sql.DB, error) {
	if addr == "" {
		return nil, errors.New("database address is required")
	}
	if maxOpenConns <= 0 {
		return nil, errors.New("max open database connections must be greater than zero")
	}
	if maxIdleConns <= 0 {
		return nil, errors.New("max idle database connections must be greater than zero")
	}
	if maxIdleTime <= 0 {
		return nil, errors.New("max idle database time must be greater than zero")
	}

	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	ok := false
	defer func() {
		if !ok {
			_ = db.Close()
		}
	}()

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(maxIdleTime)

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	ok = true
	return db, nil
}
