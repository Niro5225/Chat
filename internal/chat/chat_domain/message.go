package chat_domain

import "time"

type Message struct {
	ID        uint64
	Text      string
	ChatID    uint64
	CreatedBy uint64 // user.ID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewMessage(Text string, chatId, userId uint64, creatTime time.Time) *Message {
	return &Message{Text: Text, ChatID: chatId, CreatedBy: userId, CreatedAt: creatTime}
}
