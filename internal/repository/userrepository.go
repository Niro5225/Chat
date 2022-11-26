package repository

import (
	"chat-app/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserR struct {
	db *sqlx.DB
}

func NewUserR(db *sqlx.DB) *UserR {
	return &UserR{db: db}
}

func (r *UserR) CreateUser(user models.User) (*models.User, error) {
	err := r.db.QueryRow(
		"INSERT INTO users (first_name,last_name,email,created_at) VALUES ($1, $2,$3,$4) RETURNING id", user.FirstName, user.LastName, user.Email, user.CreatedAt,
	).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserR) DeleteUser(id uint64) error {
	row := r.db.QueryRow("DELETE FROM users WHERE id = $1", id)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (r *UserR) GetUser(id uint64) (*models.User, error) {
	var u models.User
	if err := r.db.QueryRow(
		"SELECT first_name,last_name,email FROM users WHERE id = $1", id,
	).Scan(&u.FirstName, &u.LastName, &u.Email); err != nil {
		return nil, err
	}

	u.ID = id

	return &u, nil
}

func (r *UserR) UpdateUser(user models.User) (*models.User, error) {
	row := r.db.QueryRow("UPDATE users SET first_name = $2, last_name = $3, email=$4,updated_at=$5 WHERE id = $1",
		user.ID, user.FirstName, user.LastName, user.Email, time.Now())
	if row.Err() != nil {
		return nil, row.Err()
	}
	return &user, nil
}

func (r *UserR) CreateUserCredential(credential models.UserCredential) (*models.UserCredential, error) {
	err := r.db.QueryRow("INSERT INTO user_credential (id,email,password) VALUES ($1, $2,$3) RETURNING id",
		credential.ID, credential.Email, credential.Password).Scan(&credential.ID)

	if err != nil {
		return nil, err
	}

	return &credential, nil
}

func (r *UserR) UpdateUserCredential(credential models.UserCredential) (*models.UserCredential, error) {
	row := r.db.QueryRow("UPDATE user_credential SET email=&2,password=$3 WHERE id = $1",
		credential.ID, credential.Email, credential.Password)

	if row.Err() != nil {
		return nil, row.Err()
	}

	return &credential, nil
}
