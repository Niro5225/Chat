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
