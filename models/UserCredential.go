package models

import (
	"crypto/sha1"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UserCredential struct {
	ID       uint64 // user.ID One to One
	Email    string
	Password string
}

func NewUserCredential(password string) *UserCredential {
	return &UserCredential{Password: password}
}

func Encryption_password(password string) string {
	salt := "njsankdmbdlekkgo"
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (u *UserCredential) Validate_password() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Password, validation.Required, validation.Length(8, 150)),
		//ADD email validate
	)

}
