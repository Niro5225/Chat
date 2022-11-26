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

func (r *ChatR) GetChat(id uint64) (*models.Chat, error) {
	var chat *models.Chat
	if err := r.db.QueryRow(
		"SELECT chat_name,chat_description,created_by,created_at,updated_at FROM chats WHERE id = $1", id,
	).Scan(&chat.Name, &chat.Description, &chat.CreatedBy, &chat, chat.CreatedAt, &chat.UpdatedAt); err != nil {
		return nil, err
	}
	return chat, nil
}

func (r *ChatR) UpdateChat(chat models.Chat) (*models.Chat, error) {
	row := r.db.QueryRow("UPDATE chats SET chat_name = $2, chat_description = $3, created_by=$4, updated_at=$5 WHERE id = $1",
		chat.Name, chat.Description, chat.CreatedBy, chat.UpdatedAt)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return &chat, nil
}

func (r *ChatR) DeleteChat(id uint64) error {
	row := r.db.QueryRow("DELETE FROM chats WHERE id=$1", id)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}
