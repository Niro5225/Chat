package chat_domain

import "chat-app/internal/models"

type ChatService interface {
	GetChat(id uint64) (*models.Chat, error)
	GetChats(filter *models.ChatFilter) ([]models.Chat, error)
	CreateChat(chat models.Chat) (*models.Chat, error)
	UpdateChat(chat models.Chat) (*models.Chat, error)
	DeleteChat(id uint64) error
}

type ChatServiceImp struct {
	repo ChatRepository
}

func NewChatServiceImp(repo ChatRepository) *ChatServiceImp {
	return &ChatServiceImp{repo: repo}
}

func (s *ChatServiceImp) CreateChat(chat models.Chat) (*models.Chat, error) {
	return s.repo.CreateChat(chat)
}

func (s *ChatServiceImp) GetChat(id uint64) (*models.Chat, error) {
	return s.repo.GetChat(id)
}

func (s *ChatServiceImp) GetChats(filter *models.ChatFilter) ([]models.Chat, error) {
	return s.repo.GetChats(filter)
}

func (s *ChatServiceImp) UpdateChat(chat models.Chat) (*models.Chat, error) {
	return s.repo.UpdateChat(chat)
}

func (s *ChatServiceImp) DeleteChat(id uint64) error {
	return s.repo.DeleteChat(id)
}
