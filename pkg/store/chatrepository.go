package store

import "chat/models"

type ChatRepository struct {
	store *Store
}

func (r *ChatRepository) CreateChat(ch *models.Chat) (*models.Chat, error) {
	if err := ch.ChatValidation(); err != nil {
		return nil, err
	}

	err := r.store.db.QueryRow("INSERT INTO chats (name, description,created_by,created_at) VALUES ($1,$2,$3,$4) RETURNING id",
		ch.Name, ch.Description, ch.CreatedBy, ch.CreatedAt).Scan(&ch.ID)

	if err != nil {
		return nil, err
	}

	return ch, nil
}

func (r *ChatRepository) FindChatById(id uint64) (*models.Chat, error) {
	var ch *models.Chat
	if err := r.store.db.QueryRow(
		"SELECT id,name,description,created_by FROM chats WHERE id = $1", id,
	).Scan(&ch.ID, &ch.Name, &ch.Description, &ch.CreatedBy); err != nil {
		return nil, err
	}
	return ch, nil
}
