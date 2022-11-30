package chat_domain

import "chat-app/internal/models"

type UserRepository interface {
	GetUser(id uint64) (*models.User, error)
	GetUsers(userFilter *models.UserFilter) ([]models.User, error)
	SignIn(email, password string) (*models.User, string, error)                                 //LOGIN
	SignUp(user models.User, userCredential models.UserCredential) (*models.User, string, error) //REG
	CreateUser(user models.User) (*models.User, error)
	UpdateUser(user models.User) (*models.User, error)
	CreateUserCredential(credential models.UserCredential) (*models.UserCredential, error)
	GetUserCredential(email string) (*models.UserCredential, error)
	UpdateUserCredential(credential models.UserCredential) (*models.UserCredential, error)
	DeleteUser(id uint64) error
}
