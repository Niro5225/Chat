package service

import (
	"chat-app/internal/repository"
	"chat-app/models"
)

type MessageService struct {
	repo repository.MessageRepository
}

func (s *MessageService) GetMessage(id uint64) (*models.Message, error) {
	return s.repo.GetMessage(id)
}

func (s *MessageService) CreateMessage(chat models.Message) (*models.Message, error) {
	return s.repo.CreateMessage(chat)
}

func (s *MessageService) UpdateMessage(chat models.Message) (*models.Message, error) {
	return s.repo.UpdateMessage(chat)
}

func (s *MessageService) DeleteMessage(id uint64) error {
	return s.repo.DeleteMessage(id)
}

// func (s *MessageService) CreateUserMessages(userMessage []models.UserMessage) (*models.UserMessage, error) {
// 	return s.repo.CreateUserMessages(userMessage)
// }

func (s *MessageService) UpdateUserMessage(chat models.UserMessage) (*models.UserMessage, error) {
	return s.repo.UpdateUserMessage(chat)
}

func (s *MessageService) DeleteUserMessage(userMessage models.UserMessage) error {
	return s.repo.DeleteUserMessage(userMessage)
}
