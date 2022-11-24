package store

import "chat/models"

type MessageRepository struct {
	store *Store
}

func (r *MessageRepository) CreateMessage(m *models.Message) (*models.Message, error) {

	err := r.store.db.QueryRow("INSERT INTO messages (message_text, chat_id,created_by,created_at) VALUES ($1,$2,$3,$4) RETURNING id",
		m.Text, m.ChatID, m.CreatedBy, m.CreatedAt).Scan(&m.ID)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *MessageRepository) GetChatMessages(id uint64) ([]models.Message, error) {
	var messages []models.Message

	rows, err := r.store.db.Queryx("SELECT * FROM messages WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var m models.Message
		err = rows.StructScan(&m)
		messages = append(messages, m)
	}

	return messages, nil
}
