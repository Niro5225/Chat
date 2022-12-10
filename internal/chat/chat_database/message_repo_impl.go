package chat_database

import (
	"chat-app/internal/chat/chat_domain"
	"chat-app/internal/user/user_domain"
	"time"

	"github.com/jmoiron/sqlx"
)

type MessageRepoImpl struct {
	db *sqlx.DB
}

func NewMessageRepoImpl(db *sqlx.DB) *MessageRepoImpl {
	return &MessageRepoImpl{db: db}
}

func (r *MessageRepoImpl) GetMessage(id uint64) (*chat_domain.Message, error) {
	var m chat_domain.Message
	if err := r.db.QueryRow(
		"SELECT id,message_text,chat_id,created_by,created_at FROM messages WHERE id = $1", id,
	).Scan(&m.ID, &m.Text, &m.ChatID, &m.CreatedBy, &m.CreatedAt); err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *MessageRepoImpl) GetMessages(filter *chat_domain.MessageFilter) ([]chat_domain.Message, error) {
	var messages []chat_domain.Message
	if filter != nil {
		if filter.IDs != nil {
			for _, id := range filter.IDs {
				var m chat_domain.Message
				if err := r.db.QueryRow(
					"SELECT id,message_text,chat_id,created_by,created_at FROM messages WHERE id = $1", id,
				).Scan(&m.ID, &m.Text, &m.ChatID, &m.CreatedBy, &m.CreatedAt); err != nil {
					return nil, err
				}
				messages = append(messages, m)
			}
		} else if filter.Search != nil {
			rows, err := r.db.Queryx("SELECT * FROM messages WHERE message_text=$1", filter.Search)
			if err != nil {
				return nil, err
			}
			for rows.Next() {
				var m chat_domain.Message
				err = rows.StructScan(&m)
				messages = append(messages, m)
			}

		} else if filter.ChatIDs != nil {
			for _, id := range filter.ChatIDs {
				var m chat_domain.Message
				if err := r.db.QueryRow(
					"SELECT id,message_text,chat_id,created_by,created_at FROM messages WHERE chat_id = $1", id,
				).Scan(&m.ID, &m.Text, &m.ChatID, &m.CreatedBy, &m.CreatedAt); err != nil {
					return nil, err
				}
				messages = append(messages, m)
			}
		} else if filter.UserIDs != nil {
			for _, id := range filter.UserIDs {
				var m chat_domain.Message
				if err := r.db.QueryRow(
					"SELECT id,message_text,chat_id,created_by,created_at FROM messages WHERE created_by = $1", id,
				).Scan(&m.ID, &m.Text, &m.ChatID, &m.CreatedBy, &m.CreatedAt); err != nil {
					return nil, err
				}
				messages = append(messages, m)
			}
		}
	} else {
		rows, err := r.db.Queryx("SELECT * FROM messages")
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var m chat_domain.Message
			err = rows.StructScan(&m)
			messages = append(messages, m)
		}

	}
	return messages, nil
}

func (r *MessageRepoImpl) CreateMessage(message chat_domain.Message) (*chat_domain.Message, error) {
	err := r.db.QueryRow(
		"INSERT INTO messages (message_text,chat_id,created_by,created_at) VALUES ($1, $2,$3,$4) RETURNING id",
		message.Text, message.ChatID, message.CreatedBy, message.CreatedAt).Scan(&message.ID)

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (r *MessageRepoImpl) UpdateMessage(message chat_domain.Message) (*chat_domain.Message, error) {
	row := r.db.QueryRow("UPDATE messages SET message_text = $2, chat_id = $3, created_by=$4,updated_at=$5 WHERE id = $1",
		message.ChatID, message.Text, message.ChatID, message.CreatedBy, time.Now())
	if row.Err() != nil {
		return nil, row.Err()
	}
	return &message, nil
}

func (r *MessageRepoImpl) DeleteMessage(id uint64) error {
	row := r.db.QueryRow("DELETE FROM messages WHERE id=$1", id)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (r *MessageRepoImpl) CreateUserMessages(userMessages []user_domain.UserMessage) error {
	for _, userMessage := range userMessages {
		row := r.db.QueryRow(
			"INSERT INTO user_message (user_id,message_id,is_read) VALUES ($1, $2,$3)",
			userMessage.UserID, userMessage.MessageID, userMessage.IsRead)

		if row.Err() != nil {
			return row.Err()
		}
	}
	return nil
}

func (r *MessageRepoImpl) UpdateUserMessage(userMessage user_domain.UserMessage) (*user_domain.UserMessage, error) {
	row := r.db.QueryRow("UPDATE user_message SET user_id = $1, message_id = $2, is_read=$3 WHERE user_id = $1",
		userMessage.UserID, userMessage.MessageID, userMessage.IsRead)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return &userMessage, nil
}

func (r *MessageRepoImpl) DeleteUserMessage(userMessage user_domain.UserMessage) error {
	row := r.db.QueryRow("DELETE FROM user_message WHERE user_id=$1 AND message_id=$2", userMessage.UserID, userMessage.MessageID)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}
