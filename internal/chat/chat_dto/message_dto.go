package chatdto

import "time"

type MessageDTO struct {
	ID        uint64    `json:"message_id"`
	Text      string    `json:"text"`
	ChatID    uint64    `json:"chat_id"`
	CreatedBy uint64    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
