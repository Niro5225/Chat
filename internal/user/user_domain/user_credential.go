package user_domain

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type UserCredential struct {
	ID       uint64 // user.ID One to One
	Email    string
	Password string
}

func NewUserCredential(user_id uint64, password, email string) *UserCredential {
	HashPassword := encryption_password(password)
	return &UserCredential{ID: user_id, Password: HashPassword, Email: email}
}

func encryption_password(password string) string {
	cost := 10
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	hashPassword := fmt.Sprintf("%s", hash)
	return hashPassword

}

func (uc *UserCredential) CheckPasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(uc.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func (u *UserCredential) ValidateCredential() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Password, validation.Required, validation.Length(8, 150)),
		//ADD email validate
	)

}
