package chat_domain

import "chat-app/internal/models"

type ChatRepository interface {
	GetChat(id uint64) (*models.Chat, error)
	GetChats(filter *models.ChatFilter) ([]models.Chat, error)
	CreateChat(chat models.Chat) (*models.Chat, error)
	UpdateChat(chat models.Chat) (*models.Chat, error)
	DeleteChat(id uint64) error
}
