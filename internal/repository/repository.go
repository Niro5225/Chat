package repository

import (
	"chat-app/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	// GetUser(id uint64) (*User, error)
	// GetUsers(userFilter *UserFilter) (*User, error)
	CreateUser(user models.User) (*models.User, error)
	// UpdateUser(user User) (*User, error)
	// CreateUserCredential(credential UserCredential) (*UserCredential, error)
	// UpdateUserCredential(credential UserCredential) (*UserCredential, error)
	DeleteUser(id uint64) error
}

type ChatRepository interface {
	// GetChat(id uint64) (*Chat, error)
	// GetChats(filter *ChatFilter) ([]Chat, error)
	CreateChat(chat models.Chat) (*models.Chat, error)
	// UpdateChat(chat Chat) (*Chat, error)
	// DeleteChat(id uint64) error
}

type MessageRepository interface {
	// GetMessage(id uint64) (*Message, error)
	// GetMessages(filter *MessageFilter) ([]Message, error)
	// CreateMessage(chat Message) (*Message, error)
	// UpdateMessage(chat Message) (*Message, error)
	// DeleteMessage(id uint64) error

	// CreateUserMessages(userMessage []UserMessage) (*UserMessage, error)
	// UpdateUserMessage(chat UserMessage) (*UserMessage, error)
	// DeleteUserMessage(userMessage UserMessage) error
}

type Repository struct {
	UserRepository
	ChatRepository
	MessageRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository: NewUserR(db),
		ChatRepository: NewChatR(db),
	}
}
