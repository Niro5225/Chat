package chatdto

import "time"

type ChatDTO struct {
	Id          uint64    `json:"chat_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedBy   uint64    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
