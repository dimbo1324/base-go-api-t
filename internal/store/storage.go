package store

import "database/sql"

type Storage struct {
	Posts *PostStore
	Users *UserStore
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: NewPostStore(db),
		Users: NewUserStore(db),
	}
}
