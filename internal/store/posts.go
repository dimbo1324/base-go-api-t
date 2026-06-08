package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

const createPostQuery = `
	INSERT INTO posts (user_id, title, content, tags)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at
`

type PostStore struct {
	db *sql.DB
}

func NewPostStore(db *sql.DB) *PostStore {
	return &PostStore{db: db}
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	if post == nil {
		return errors.New("store: post is nil")
	}

	tags := post.Tags
	if tags == nil {
		tags = []string{}
	}

	return s.db.QueryRowContext(
		ctx,
		createPostQuery,
		post.UserID,
		post.Title,
		post.Content,
		pq.Array(tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
}
