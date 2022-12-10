package user_domain

type UserMessage struct {
	UserID    uint64 // user.ID
	MessageID uint64 // message.ID
	IsRead    bool   // Set true after read
}
