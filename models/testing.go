package models

import "testing"

func Test_user(t *testing.T) *User {
	return &User{
		Username:           "test_username",
		Uncrypted_Password: "12345678",
	}
}
