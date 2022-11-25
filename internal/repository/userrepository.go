package repository

import (
	"chat-app/models"

	"github.com/jmoiron/sqlx"
)

type UserR struct {
	db *sqlx.DB
}

func NewUserR(db *sqlx.DB) *UserR {
	return &UserR{db: db}
}

func (r *UserR) CreateUser(user models.User) (*models.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	err := r.db.QueryRow(
		"INSERT INTO users (first_name,last_name,email,created_at) VALUES ($1, $2,$3,$4) RETURNING id", user.FirstName, user.LastName, user.Email, user.CreatedAt,
	).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserR) DeleteUser(id uint64) error {
	row := r.db.QueryRow("DELETE FROM users WHERE id=$1", id)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}
