package repository

import (
	"chat-app/models"

	"github.com/jmoiron/sqlx"
)

type ChatR struct {
	db *sqlx.DB
}

func NewChatR(db *sqlx.DB) *ChatR {
	return &ChatR{db: db}
}

func (r *ChatR) CreateChat(chat models.Chat) (*models.Chat, error) {
	if err := chat.ChatValidation(); err != nil {
		return nil, err
	}

	err := r.db.QueryRow("INSERT INTO chats (chat_name, chat_description,created_by,created_at) VALUES ($1,$2,$3,$4) RETURNING id", chat.Name, chat.Description, chat.CreatedBy, chat.CreatedAt).Scan(&chat.ID)

	if err != nil {
		return nil, err
	}

	return &chat, nil
}
