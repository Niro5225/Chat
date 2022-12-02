package chat_domain

type ChatRepository interface {
	GetChat(id uint64) (*Chat, error)
	GetChats(filter *ChatFilter) ([]Chat, error)
	CreateChat(chat Chat) (*Chat, error)
	UpdateChat(chat Chat) (*Chat, error)
	DeleteChat(id uint64) error
}
