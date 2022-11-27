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

func NewChat(name, descr string, id uint64, createTime time.Time) *Chat {
	return &Chat{Name: name, Description: descr, CreatedBy: id, CreatedAt: createTime}
}

func (ch *Chat) ChatValidation() error {
	return nil
}

//TODO:Validation
