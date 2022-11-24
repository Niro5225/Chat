package store

import (
	"chat/models"
	"errors"
	"fmt"
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
func (r *User_repository) Create(u *models.User, uc *models.UserCredential) (*models.User, error) {

	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := uc.Validate_password(); err != nil {
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

	hash := models.Encryption_password(uc.Password)

	fmt.Println(uc)

	err = r.store.db.QueryRow(
		"INSERT INTO user_credential (id,email,password) VALUES ($1, $2,$3) RETURNING id", u.ID, u.Email, hash).Scan(&u.ID)

	if err != nil {
		return nil, err
	}

	return u, nil
}

//Метод логина
func (r *User_repository) Login(u *models.User, password string) (*models.User, error) {
	//генерация токена
	_, err := generate_token(u)
	if err != nil {
		return nil, err
	}

	uc, err := r.store.User().FindUserCredential(u)

	if err != nil {
		return nil, err
	}

	fmt.Println(uc.Password, " ", password)

	if uc.Password != models.Encryption_password(password) {
		return nil, errors.New("Wrong password")
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

func (r *User_repository) FindUserCredential(u *models.User) (*models.UserCredential, error) {
	uc := &models.UserCredential{}
	if err := r.store.db.QueryRow(
		"SELECT email,password FROM user_credential WHERE id = $1", u.ID,
	).Scan(&uc.Email, &uc.Password); err != nil {
		return nil, err
	}
	return uc, nil
}

func (r *User_repository) FindById(id int) (*models.User, error) {
	var u *models.User
	if err := r.store.db.QueryRow(
		"SELECT id,first_name,last_name,email FROM users WHERE id = $1", id,
	).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *User_repository) Get_all_users() ([]models.User, error) {
	var users []models.User

	rows, err := r.store.db.Queryx("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var u models.User
		err = rows.StructScan(&u)
		users = append(users, u)
	}

	return users, nil

}
