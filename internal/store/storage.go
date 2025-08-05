package store

import (
	"context"
	"database/sql"
)

type PostRepository interface {
	Create(ctx context.Context, post *Post) error
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetAll() ([]User, error)
}

type Storage struct {
	Posts PostRepository
	Users UserRepository
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UserStore{db},
	}
}