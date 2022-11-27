package repository

import (
	"chat-app/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type MessageR struct {
	db *sqlx.DB
}

func NewMessageR(db *sqlx.DB) *MessageR {
	return &MessageR{db: db}
}

func (r *MessageR) GetMessage(id uint64) (*models.Message, error) {
	var m models.Message
	if err := r.db.QueryRow(
		"SELECT id,message_text,chat_id,created_by,created_at FROM messages WHERE id = $1", id,
	).Scan(&m.ID, &m.Text, &m.ChatID, &m.CreatedBy, &m.CreatedAt); err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *MessageR) CreateMessage(message models.Message) (*models.Message, error) {
	err := r.db.QueryRow(
		"INSERT INTO messages (message_text,chat_id,created_by,created_at) VALUES ($1, $2,$3,$4) RETURNING id",
		message.Text, message.ChatID, message.CreatedBy, message.CreatedAt).Scan(&message.ID)

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (r *MessageR) UpdateMessage(message models.Message) (*models.Message, error) {
	row := r.db.QueryRow("UPDATE messages SET message_text = $2, chat_id = $3, created_by=$4,updated_at=$5 WHERE id = $1",
		message.ChatID, message.Text, message.ChatID, message.CreatedBy, time.Now())
	if row.Err() != nil {
		return nil, row.Err()
	}
	return &message, nil
}

func (r *MessageR) DeleteMessage(id uint64) error {
	row := r.db.QueryRow("DELETE FROM messages WHERE id=$1", id)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

// func (r *MessageR) CreateUserMessages(userMessage []models.UserMessage) (*models.UserMessage, error) {
// 	err := r.db.QueryRow(
// 		"INSERT INTO user_message (user_id,message_id,is_read) VALUES ($1, $2,$3) RETURNING id")

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &message, nil
// }

func (r *MessageR) UpdateUserMessage(userMessage models.UserMessage) (*models.UserMessage, error) {
	return nil, nil
}

func (r *MessageR) DeleteUserMessage(userMessage models.UserMessage) error {
	row := r.db.QueryRow("DELETE FROM user_message WHERE user_id=$1 AND message_id=$2", userMessage.UserID, userMessage.MessageID)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}
