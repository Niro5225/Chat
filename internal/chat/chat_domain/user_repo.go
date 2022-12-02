package chat_domain

import "chat-app/internal/user"

type UserRepository interface {
	GetUser(id uint64) (*user.User, error)
	GetUsers(userFilter *UserFilter) ([]user.User, error)
	// SignIn(email, password string) (*models.User, string, error)                                 //LOGIN
	// SignUp(user models.User, userCredential models.UserCredential) (*models.User, string, error) //REG
	CreateUser(user user.User) (*user.User, error)
	UpdateUser(user user.User) (*user.User, error)
	CreateUserCredential(credential UserCredential) (*UserCredential, error)
	GetUserCredential(email string) (*UserCredential, error)
	UpdateUserCredential(credential UserCredential) (*UserCredential, error)
	DeleteUser(id uint64) error
}
