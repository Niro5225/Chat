package chat_database

import (
	"chat-app/internal/chat/chat_domain"

	"github.com/jmoiron/sqlx"
)

type ChatRepoImpl struct {
	db *sqlx.DB
}

func NewChatRepoImpl(db *sqlx.DB) *ChatRepoImpl {
	return &ChatRepoImpl{db: db}
}

func (r *ChatRepoImpl) CreateChat(chat chat_domain.Chat) (*chat_domain.Chat, error) {
	if err := chat.ChatValidation(); err != nil {
		return nil, err
	}

	err := r.db.QueryRow("INSERT INTO chats (chat_name, chat_description,created_by,created_at) VALUES ($1,$2,$3,$4) RETURNING id", chat.Name, chat.Description, chat.CreatedBy, chat.CreatedAt).Scan(&chat.ID)

	if err != nil {
		return nil, err
	}

	return &chat, nil
}

func (r *ChatRepoImpl) GetChat(id uint64) (*chat_domain.Chat, error) {
	var chat chat_domain.Chat
	if err := r.db.QueryRow(
		"SELECT chat_name,chat_description,created_by,created_at FROM chats WHERE id = $1", id,
	).Scan(&chat.Name, &chat.Description, &chat.CreatedBy, &chat.CreatedAt); err != nil {
		return nil, err
	}
	return &chat, nil
}

func (r *ChatRepoImpl) GetChats(filter *chat_domain.ChatFilter) ([]chat_domain.Chat, error) {
	var chats []chat_domain.Chat
	if filter != nil {
		if filter.IDs != nil {
			for _, id := range filter.IDs {
				var chat chat_domain.Chat
				if err := r.db.QueryRow(
					"SELECT id, chat_name,chat_description,created_by,created_at FROM chats WHERE id = $1", id,
				).Scan(&chat.ID, &chat.Name, &chat.Description, &chat.CreatedBy, &chat.CreatedAt); err != nil {
					return nil, err
				}
				chats = append(chats, chat)
			}
		} else if filter.Search != nil {
			rows, err := r.db.Queryx("SELECT * FROM chats WHERE chat_name=$1 OR chat_description=$1", filter.Search)
			if err != nil {
				return nil, err
			}
			for rows.Next() {
				var chat chat_domain.Chat
				err = rows.StructScan(&chat)
				chats = append(chats, chat)
			}
		} else if filter.UserIDs != nil {
			for _, id := range filter.UserIDs {
				var chat chat_domain.Chat
				if err := r.db.QueryRow(
					"SELECT id, chat_name,chat_description,created_by,created_at FROM chats WHERE created_by = $1", id,
				).Scan(&chat.ID, &chat.Name, &chat.Description, &chat.CreatedBy, &chat.CreatedAt); err != nil {
					return nil, err
				}
				chats = append(chats, chat)
			}
		}
	} else {
		rows, err := r.db.Queryx("SELECT * FROM chats")
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var chat chat_domain.Chat
			err = rows.StructScan(&chat)
			chats = append(chats, chat)
		}
	}
	return chats, nil
}

func (r *ChatRepoImpl) UpdateChat(chat chat_domain.Chat) (*chat_domain.Chat, error) {
	row := r.db.QueryRow("UPDATE chats SET chat_name = $2, chat_description = $3, created_by=$4, updated_at=$5 WHERE id = $1",
		chat.ID, chat.Name, chat.Description, chat.CreatedBy, chat.UpdatedAt)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return &chat, nil
}

func (r *ChatRepoImpl) DeleteChat(id uint64) error {
	row := r.db.QueryRow("DELETE FROM chats WHERE id=$1", id)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}
