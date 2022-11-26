package models

import (
	"crypto/sha1"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UserCredential struct {
	ID       uint64 // user.ID One to One
	Email    string
	Password string
}

func NewUserCredential(user_id uint64, password, email string) *UserCredential {
	return &UserCredential{ID: user_id, Password: password, Email: email}
}

func (uc *UserCredential) Encryption_password() {
	salt := "njsankdmbdlekkgo"
	hash := sha1.New()
	hash.Write([]byte(uc.Password))

	uc.Password = string(hash.Sum([]byte(salt)))

}

func (u *UserCredential) ValidateCredential() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Password, validation.Required, validation.Length(8, 150)),
		//ADD email validate
	)

}
