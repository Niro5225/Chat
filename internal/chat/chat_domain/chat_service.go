package chat_domain

type ChatService interface {
	GetChat(id uint64) (*Chat, error)
	GetChats(filter *ChatFilter) ([]Chat, error)
	CreateChat(chat Chat) (*Chat, error)
	UpdateChat(chat Chat) *Chat
	DeleteChat(id uint64) error
}

type ChatServiceImp struct {
	repo ChatRepository
}

func NewChatServiceImp(repo ChatRepository) *ChatServiceImp {
	return &ChatServiceImp{repo: repo}
}

func (s *ChatServiceImp) CreateChat(chat Chat) (*Chat, error) {
	return s.repo.CreateChat(chat)
}

func (s *ChatServiceImp) GetChat(id uint64) (*Chat, error) {
	return s.repo.GetChat(id)
}

func (s *ChatServiceImp) GetChats(filter *ChatFilter) ([]Chat, error) {
	return s.repo.GetChats(filter)
}

func (s *ChatServiceImp) UpdateChat(chat Chat) (*Chat, error) {
	return s.repo.UpdateChat(chat)
}

func (s *ChatServiceImp) DeleteChat(id uint64) error {
	return s.repo.DeleteChat(id)
}
