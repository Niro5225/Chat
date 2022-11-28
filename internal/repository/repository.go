package repository

import (
	"chat-app/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetUser(id uint64) (*models.User, error)
	GetUsers(userFilter *models.UserFilter) ([]models.User, error)
	CreateUser(user models.User) (*models.User, error)
	UpdateUser(user models.User) (*models.User, error)
	CreateUserCredential(credential models.UserCredential) (*models.UserCredential, error)
	UpdateUserCredential(credential models.UserCredential) (*models.UserCredential, error)
	DeleteUser(id uint64) error
}

type ChatRepository interface {
	GetChat(id uint64) (*models.Chat, error)
	GetChats(filter *models.ChatFilter) ([]models.Chat, error)
	CreateChat(chat models.Chat) (*models.Chat, error)
	UpdateChat(chat models.Chat) (*models.Chat, error)
	DeleteChat(id uint64) error
}

type MessageRepository interface {
	GetMessage(id uint64) (*models.Message, error)
	GetMessages(filter *models.MessageFilter) ([]models.Message, error)
	CreateMessage(chat models.Message) (*models.Message, error)
	UpdateMessage(chat models.Message) (*models.Message, error)
	DeleteMessage(id uint64) error

	CreateUserMessages(userMessage []models.UserMessage) error
	UpdateUserMessage(message models.UserMessage) (*models.UserMessage, error)
	DeleteUserMessage(userMessage models.UserMessage) error
}

type Repository struct {
	UserRepository
	ChatRepository
	MessageRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository:    NewUserR(db),
		ChatRepository:    NewChatR(db),
		MessageRepository: NewMessageR(db),
	}
}
