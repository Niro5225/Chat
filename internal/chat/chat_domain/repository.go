package chat_domain

import (
	"chat-app/internal/chat/chat_database"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	UserRepository
	ChatRepository
	MessageRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository:    chat_database.NewUserRepoImpl(db),
		ChatRepository:    chat_database.NewChatRepoImpl(db),
		MessageRepository: chat_database.NewMessageRepoImpl(db),
	}
}
