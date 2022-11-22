package models

type UserCredential struct {
	ID       uint64 // user.ID One to One
	Email    string
	Password string
}
