package chat_domain

import "chat-app/internal/user/user_domain"

type MessageRepository interface {
	GetMessage(id uint64) (*Message, error)
	GetMessages(filter *MessageFilter) ([]Message, error)
	CreateMessage(chat Message) (*Message, error)
	UpdateMessage(chat Message) (*Message, error)
	DeleteMessage(id uint64) error

	CreateUserMessages(userMessage []user_domain.UserMessage) error
	UpdateUserMessage(message user_domain.UserMessage) (*user_domain.UserMessage, error)
	DeleteUserMessage(userMessage user_domain.UserMessage) error
}
