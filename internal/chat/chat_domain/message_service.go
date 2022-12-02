package chat_domain

type MessageService interface {
	GetMessage(id uint64) (*Message, error)
	GetMessages(filter *MessageFilter) ([]Message, error)
	CreateMessage(chat Message) (*Message, error)
	UpdateMessage(chat Message) (*Message, error)
	DeleteMessage(id uint64) error

	CreateUserMessages(userMessage []UserMessage) error
	UpdateUserMessage(chat UserMessage) (*UserMessage, error)
	DeleteUserMessage(userMessage UserMessage) error
}

type MessageServiceImp struct {
	repo MessageRepository
}

func NewMessageServiceImp(repo MessageRepository) *MessageServiceImp {
	return &MessageServiceImp{repo: repo}
}

func (s *MessageServiceImp) GetMessage(id uint64) (*Message, error) {
	return s.repo.GetMessage(id)
}

func (s *MessageServiceImp) GetMessages(filter *MessageFilter) ([]Message, error) {
	return s.repo.GetMessages(filter)
}

func (s *MessageServiceImp) CreateMessage(chat Message) (*Message, error) {
	return s.repo.CreateMessage(chat)
}

func (s *MessageServiceImp) UpdateMessage(chat Message) (*Message, error) {
	return s.repo.UpdateMessage(chat)
}

func (s *MessageServiceImp) DeleteMessage(id uint64) error {
	return s.repo.DeleteMessage(id)
}

func (s *MessageServiceImp) CreateUserMessages(userMessage []UserMessage) error {
	return s.repo.CreateUserMessages(userMessage)
}

func (s *MessageServiceImp) UpdateUserMessage(userMessage UserMessage) (*UserMessage, error) {
	return s.repo.UpdateUserMessage(userMessage)
}

func (s *MessageServiceImp) DeleteUserMessage(userMessage UserMessage) error {
	return s.repo.DeleteUserMessage(userMessage)
}
