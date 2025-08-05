package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID			int64  `json:"id"`
	Username	string `json:"username"`
	Email		string `json:"email"`
	Password	string `json:"-"`
	CreatedAt	string `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {

	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) GetAll() ([]User, error) {
	query := `
		SELECT id, username, email, created_at
		FROM users
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}