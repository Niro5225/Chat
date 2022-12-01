package chat_domain

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	ID        uint64    `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewUser(FirstName, LastName, Email string) *User {
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

func (u *User) ValidateUser() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.FirstName, validation.Required, validation.Length(3, 150)),
		validation.Field(&u.LastName, validation.Required, validation.Length(3, 150)),
		//ADD email validate
	)

}
