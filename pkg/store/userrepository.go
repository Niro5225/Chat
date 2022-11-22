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

//Метод создания пользователя
func (r *User_repository) Create(u *models.User) (*models.User, error) {

	if err := u.Validate(); err != nil {
		return nil, err
	}

	// if err := u.Before_create(); err != nil {
	// 	return nil, err
	// }

	err := r.store.db.QueryRow(
		"INSERT INTO users (first_name,last_name,email,created_at) VALUES ($1, $2,$3,$4) RETURNING id", u.FirstName, u.LastName, u.Email, u.CreatedAt,
	).Scan(&u.ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

//Метод логина
func (r *User_repository) Login(u *models.User) (*models.User, error) {
	//генерация токена
	_, err := generate_token(u)
	if err != nil {
		return nil, err
	}

	// u.Token = t

	return u, nil
}

//Дополнение стандартных полей для создания токена
type Token_claims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

//Метод генерации токена
func generate_token(u *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Token_claims{
		jwt.StandardClaims{
			ExpiresAt: token_ttl,
		}, int(u.ID),
	})

	return token.SignedString([]byte(token_key))
}

//Метод поиска пользователя по имени пользователя
func (r *User_repository) Find_by_email(email string) (*models.User, error) {
	u := &models.User{}
	if err := r.store.db.QueryRow(
		"SELECT id,first_name,last_name FROM users WHERE email = $1", email,
	).Scan(&u.ID, &u.FirstName, &u.LastName); err != nil {
		return nil, err
	}
	return u, nil
}
