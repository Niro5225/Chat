package models

import "time"

type Message struct {
	ID        uint64
	Text      string
	ChatID    uint64
	CreatedBy uint64 // user.ID
	CreatedAt time.Time
	UpdatedAt time.Time
}
