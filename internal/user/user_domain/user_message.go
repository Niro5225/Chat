package user_domain

type UserMessage struct {
	UserID    uint64 `db:"user_id"`    // user.ID
	MessageID uint64 `db:"message_id"` // message.ID
	IsRead    bool   `db:"is_read"`    // Set true after read
}
