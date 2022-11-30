package chat_domain

import "chat-app/internal/models"

type MessageService interface {
	GetMessage(id uint64) (*models.Message, error)
	GetMessages(filter *models.MessageFilter) ([]models.Message, error)
	CreateMessage(chat models.Message) (*models.Message, error)
	UpdateMessage(chat models.Message) (*models.Message, error)
	DeleteMessage(id uint64) error

	CreateUserMessages(userMessage []models.UserMessage) error
	UpdateUserMessage(chat models.UserMessage) (*models.UserMessage, error)
	DeleteUserMessage(userMessage models.UserMessage) error
}

type MessageServiceImp struct {
	repo MessageRepository
}

func NewMessageServiceImp(repo MessageRepository) *MessageServiceImp {
	return &MessageServiceImp{repo: repo}
}

func (s *MessageServiceImp) GetMessage(id uint64) (*models.Message, error) {
	return s.repo.GetMessage(id)
}

func (s *MessageServiceImp) GetMessages(filter *models.MessageFilter) ([]models.Message, error) {
	return s.repo.GetMessages(filter)
}

func (s *MessageServiceImp) CreateMessage(chat models.Message) (*models.Message, error) {
	return s.repo.CreateMessage(chat)
}

func (s *MessageServiceImp) UpdateMessage(chat models.Message) (*models.Message, error) {
	return s.repo.UpdateMessage(chat)
}

func (s *MessageServiceImp) DeleteMessage(id uint64) error {
	return s.repo.DeleteMessage(id)
}

func (s *MessageServiceImp) CreateUserMessages(userMessage []models.UserMessage) error {
	return s.repo.CreateUserMessages(userMessage)
}

func (s *MessageServiceImp) UpdateUserMessage(userMessage models.UserMessage) (*models.UserMessage, error) {
	return s.repo.UpdateUserMessage(userMessage)
}

func (s *MessageServiceImp) DeleteUserMessage(userMessage models.UserMessage) error {
	return s.repo.DeleteUserMessage(userMessage)
}
