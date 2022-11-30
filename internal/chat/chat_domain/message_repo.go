package chat_domain

import "chat-app/internal/models"

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
