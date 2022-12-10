package chat_domain

import "time"

type Message struct {
	ID        uint64    `db:"id"`
	Text      string    `db:"message_text"`
	ChatID    uint64    `db:"chat_id"`
	CreatedBy uint64    `db:"created_by"` // user.ID
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewMessage(Text string, chatId, userId uint64, creatTime time.Time) *Message {
	return &Message{Text: Text, ChatID: chatId, CreatedBy: userId, CreatedAt: creatTime}
}
