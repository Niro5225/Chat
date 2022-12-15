package user_domain

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	ID        uint64    `db:"id" json:"id"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func NewUser(FirstName, LastName, Email string) *User {
	createTime := time.Now()
	return &User{FirstName: FirstName, LastName: LastName, Email: Email, CreatedAt: createTime}
}

func (u *User) ValidateUser() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.FirstName, validation.Required, validation.Length(3, 150)),
		validation.Field(&u.LastName, validation.Required, validation.Length(3, 150)),
		//ADD email validate
	)

}
