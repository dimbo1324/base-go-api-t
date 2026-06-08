package store

import (
	"context"
	"database/sql"
	"errors"
)

const createUserQuery = `
	INSERT INTO users (username, password_hash, email)
	VALUES ($1, $2, $3)
	RETURNING id, created_at
`

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	if user == nil {
		return errors.New("store: user is nil")
	}

	return s.db.QueryRowContext(ctx, createUserQuery, user.Username, user.PasswordHash, user.Email).
		Scan(&user.ID, &user.CreatedAt)
}
