package chat_domain

import "time"

type Chat struct {
	ID          uint64    `db:"id"`
	Name        string    `db:"chat_name"`
	Description string    `db:"chat_description"`
	CreatedBy   uint64    `db:"created_by"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func NewChat(name, descr string, id uint64, createTime time.Time) *Chat {
	return &Chat{Name: name, Description: descr, CreatedBy: id, CreatedAt: createTime}
}

func (ch *Chat) ChatValidation() error {
	return nil
}

//TODO:Validation
