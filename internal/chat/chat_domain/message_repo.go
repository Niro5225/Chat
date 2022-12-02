package chat_domain

type MessageRepository interface {
	GetMessage(id uint64) (*Message, error)
	GetMessages(filter *MessageFilter) ([]Message, error)
	CreateMessage(chat Message) (*Message, error)
	UpdateMessage(chat Message) (*Message, error)
	DeleteMessage(id uint64) error

	CreateUserMessages(userMessage []UserMessage) error
	UpdateUserMessage(message UserMessage) (*UserMessage, error)
	DeleteUserMessage(userMessage UserMessage) error
}
