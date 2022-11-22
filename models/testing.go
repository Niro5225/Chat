package models

import (
	"testing"
	"time"
)

func Test_user(t *testing.T) *User {
	return &User{
		FirstName: "testFirstName",
		LastName:  "testLastName",
		Email:     "test@email.com",
		CreatedAt: time.Now(),
	}
}
