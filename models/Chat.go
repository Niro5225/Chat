package models

import "time"

type Chat struct {
	ID          uint64
	Name        string
	Description string
	CreatedBy   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewChat(name, descr string, id int, createTime time.Time) *Chat {
	return &Chat{Name: name, Description: descr, CreatedBy: id, CreatedAt: createTime}
}

func (ch *Chat) ChatValidation() error {
	return nil
}

//TODO:Validation
