package models

import (
	"crypto/sha1"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	Id                 int    `json:"-" db:"id"`
	Username           string `json:"username" binding:"required"`
	Encrypted_Password string `json:"password" binding:"required"`
	Uncrypted_Password string
	Token              string
}

func New_ueser(username, password string) *User {
	return &User{Username: username, Uncrypted_Password: password}
}

func (u *User) Before_create() error {
	if len(u.Uncrypted_Password) > 0 {
		enc := encryption_password(u.Uncrypted_Password)
		u.Encrypted_Password = enc
	}

	return nil
}

// func encryption_password(password string) (string, error) {
// 	e, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(e), nil
// }

func encryption_password(password string) string {
	salt := "njsankdmbdlekkgo"
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Uncrypted_Password, validation.By(requiredIf(u.Encrypted_Password == "")), validation.Length(8, 100)),
		validation.Field(&u.Username, validation.Required, validation.Length(3, 150)))
}

func (u *User) Check_password(password string) (bool, error) {
	e_pass := encryption_password(password)
	fmt.Println(e_pass)
	fmt.Println(u.Encrypted_Password)

	if e_pass == u.Encrypted_Password {
		return true, nil
	} else {
		return false, nil
	}
}
