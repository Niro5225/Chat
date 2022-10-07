package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                 int    `json:"-" db:"id"`
	Username           string `json:"username" binding:"required"`
	Encrypted_Password string `json:"password" binding:"required"`
	Uncrypted_Password string
}

func (u *User) Before_create() error {
	if len(u.Uncrypted_Password) > 0 {
		enc, err := encryption_password(u.Uncrypted_Password)
		if err != nil {
			return err
		}
		u.Encrypted_Password = enc
	}

	return nil
}

func encryption_password(password string) (string, error) {
	e, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(e), nil
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Uncrypted_Password, validation.By(requiredIf(u.Encrypted_Password == "")), validation.Length(8, 100)),
		validation.Field(&u.Username, validation.Required, validation.Length(3, 150)))
}
