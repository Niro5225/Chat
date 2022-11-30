package chat_domain

type Services struct {
	UserService
	ChatService
	MessageService
}

func NewServices(r *Repository) *Services {
	return &Services{
		UserService:    NewUserServiceImp(r.UserRepository),
		ChatService:    NewChatServiceImp(r.ChatRepository),
		MessageService: NewMessageServiceImp(r.MessageRepository),
	}
}
