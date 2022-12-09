package user_domain

import (
	"chat-app/internal/chat/chat_domain"
	"chat-app/internal/user"
)

type UserRepository interface {
	GetUser(id uint64) (*user.User, error)
	GetUsers(userFilter *chat_domain.UserFilter) ([]user.User, error)
	// SignIn(email, password string) (*models.User, string, error)                                 //LOGIN
	// SignUp(user models.User, userCredential models.UserCredential) (*models.User, string, error) //REG
	CreateUser(user user.User) (*user.User, error)
	UpdateUser(user user.User) (*user.User, error)
	CreateUserCredential(credential chat_domain.UserCredential) (*chat_domain.UserCredential, error)
	GetUserCredential(email string) (*chat_domain.UserCredential, error)
	UpdateUserCredential(credential chat_domain.UserCredential) (*chat_domain.UserCredential, error)
	DeleteUser(id uint64) error
}
