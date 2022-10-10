package store

import (
	"chat/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type User_repository struct {
	store *Store
}

const (
	token_ttl = int64(12 * time.Hour)
	token_key = "sdafAF#fasd#fdfsdcvbbthrhj#dsasda"
)

func (r *User_repository) Create(u *models.User) (*models.User, error) {

	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.Before_create(); err != nil {
		return nil, err
	}

	err := r.store.db.QueryRow(
		"INSERT INTO users (username,password_hash) VALUES ($1, $2) RETURNING id", u.Username, u.Encrypted_Password,
	).Scan(&u.Id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *User_repository) Login(u *models.User) (*models.User, error) {
	t, err := generate_token(u)
	if err != nil {
		return nil, err
	}

	u.Token = t

	return u, nil
}

type Token_claims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func generate_token(u *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Token_claims{
		jwt.StandardClaims{
			ExpiresAt: token_ttl,
		}, u.Id,
	})

	return token.SignedString([]byte(token_key))
}

func (r *User_repository) Find_by_username(username string) (*models.User, error) {
	u := &models.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, username, password_hash FROM users WHERE username = $1", username,
	).Scan(&u.Id, &u.Username, &u.Encrypted_Password); err != nil {
		return nil, err
	}

	return u, nil
}
