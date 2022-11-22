package models

import (
	"crypto/sha1"
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	ID        uint64
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New_ueser(FirstName, LastName, Email string) *User {
	createTime := time.Now()
	return &User{FirstName: FirstName, LastName: LastName, Email: Email, CreatedAt: createTime}
}

// func (u *User) Before_create() error {
// 	if len(u.Uncrypted_Password) > 0 {
// 		enc := encryption_password(u.Uncrypted_Password)
// 		u.Encrypted_Password = enc
// 	}

// 	return nil
// }

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
		validation.Field(&u.FirstName, validation.Required, validation.Length(3, 150)),
		validation.Field(&u.LastName, validation.Required, validation.Length(3, 150)),
		//ADD email validate
	)

}
