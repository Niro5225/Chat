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

func TestUserCredential(t *testing.T) *UserCredential {
	return &UserCredential{
		Password: "testpassword",
	}
}

func TestChat(t *testing.T) *Chat {
	return &Chat{
		Name:        "test chat name",
		Description: "test description",
		CreatedBy:   1,
		CreatedAt:   time.Now(),
	}
}

func TestMessage(t *testing.T) *Message {
	return &Message{
		Text:      "test message",
		ChatID:    1,
		CreatedBy: 1,
	}
}
