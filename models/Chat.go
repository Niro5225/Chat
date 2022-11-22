package models

import "time"

type Chat struct {
	ID          uint64
	Name        string
	Description string
	CreatedBy   uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
