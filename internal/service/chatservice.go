package service

import (
	"chat-app/internal/repository"
	"chat-app/models"
)

type ChatService struct {
	repo repository.ChatRepository
}

func NewChatService(repo repository.ChatRepository) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) CreateChat(chat models.Chat) (*models.Chat, error) {
	return s.repo.CreateChat(chat)
}

func (s *ChatService) GetChat(id uint64) (*models.Chat, error) {
	return s.repo.GetChat(id)
}

func (s *ChatService) GetChats(filter *models.ChatFilter) ([]models.Chat, error) {
	return s.repo.GetChats(filter)
}

func (s *ChatService) UpdateChat(chat models.Chat) (*models.Chat, error) {
	return s.repo.UpdateChat(chat)
}

func (s *ChatService) DeleteChat(id uint64) error {
	return s.repo.DeleteChat(id)
}
