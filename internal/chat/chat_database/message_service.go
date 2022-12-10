package chat_database

import (
	"chat-app/internal/chat/chat_domain"
	"chat-app/internal/user/user_domain"
)

type MessageService interface {
	GetMessage(id uint64) (*chat_domain.Message, error)
	GetMessages(filter *chat_domain.MessageFilter) ([]chat_domain.Message, error)
	CreateMessage(chat chat_domain.Message) (*chat_domain.Message, error)
	UpdateMessage(chat chat_domain.Message) (*chat_domain.Message, error)
	DeleteMessage(id uint64) error

	CreateUserMessages(userMessage []user_domain.UserMessage) error
	UpdateUserMessage(chat user_domain.UserMessage) (*user_domain.UserMessage, error)
	DeleteUserMessage(userMessage user_domain.UserMessage) error
}

type MessageServiceImp struct {
	repo MessageRepository
}

func NewMessageServiceImp(repo MessageRepository) *MessageServiceImp {
	return &MessageServiceImp{repo: repo}
}

func (s *MessageServiceImp) GetMessage(id uint64) (*chat_domain.Message, error) {
	return s.repo.GetMessage(id)
}

func (s *MessageServiceImp) GetMessages(filter *chat_domain.MessageFilter) ([]chat_domain.Message, error) {
	return s.repo.GetMessages(filter)
}

func (s *MessageServiceImp) CreateMessage(chat chat_domain.Message) (*chat_domain.Message, error) {
	return s.repo.CreateMessage(chat)
}

func (s *MessageServiceImp) UpdateMessage(chat chat_domain.Message) (*chat_domain.Message, error) {
	return s.repo.UpdateMessage(chat)
}

func (s *MessageServiceImp) DeleteMessage(id uint64) error {
	return s.repo.DeleteMessage(id)
}

func (s *MessageServiceImp) CreateUserMessages(userMessage []user_domain.UserMessage) error {
	return s.repo.CreateUserMessages(userMessage)
}

func (s *MessageServiceImp) UpdateUserMessage(userMessage user_domain.UserMessage) (*user_domain.UserMessage, error) {
	return s.repo.UpdateUserMessage(userMessage)
}

func (s *MessageServiceImp) DeleteUserMessage(userMessage user_domain.UserMessage) error {
	return s.repo.DeleteUserMessage(userMessage)
}
