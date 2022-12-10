package chat_database

import (
	"chat-app/internal/chat/chat_domain"
	"chat-app/internal/user/user_domain"
)

type MessageRepository interface {
	GetMessage(id uint64) (*chat_domain.Message, error)
	GetMessages(filter *chat_domain.MessageFilter) ([]chat_domain.Message, error)
	CreateMessage(chat chat_domain.Message) (*chat_domain.Message, error)
	UpdateMessage(chat chat_domain.Message) (*chat_domain.Message, error)
	DeleteMessage(id uint64) error

	CreateUserMessages(userMessage []user_domain.UserMessage) error
	UpdateUserMessage(message user_domain.UserMessage) (*user_domain.UserMessage, error)
	DeleteUserMessage(userMessage user_domain.UserMessage) error
}
