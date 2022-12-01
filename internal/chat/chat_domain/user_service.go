package chat_domain

import "chat-app/internal/models"

type UserService interface {
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

type UserServiceImp struct {
	repo UserRepository
}

func NewUserServiceImp(repo UserRepository) *UserServiceImp {
	return &UserServiceImp{repo: repo}
}

func (s *UserServiceImp) CreateUser(user models.User) (*models.User, error) {
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}
	return s.repo.CreateUser(user)
}

func (s *UserServiceImp) GetUserCredential(email string) (*models.UserCredential, error) {
	return s.repo.GetUserCredential(email)
}

func (s *UserServiceImp) DeleteUser(id uint64) error {
	return s.repo.DeleteUser(id)
}

func (s *UserServiceImp) GetUser(id uint64) (*models.User, error) {
	return s.repo.GetUser(id)
}

func (s *UserServiceImp) GetUsers(userFilter *models.UserFilter) ([]models.User, error) {
	return s.repo.GetUsers(userFilter)
}

func (s *UserServiceImp) SignIn(email, password string) (*models.User, string, error) {
	return s.repo.SignIn(email, password)
}

func (s *UserServiceImp) SignUp(user models.User, userCredential models.UserCredential) (*models.User, string, error) {
	return s.repo.SignUp(user, userCredential)
}

func (s *UserServiceImp) UpdateUser(user models.User) (*models.User, error) {
	return s.repo.UpdateUser(user)
}

func (s *UserServiceImp) CreateUserCredential(credential models.UserCredential) (*models.UserCredential, error) {
	if err := credential.ValidateCredential(); err != nil {
		return nil, err
	}
	return s.repo.CreateUserCredential(credential)
}

func (s *UserServiceImp) UpdateUserCredential(credential models.UserCredential) (*models.UserCredential, error) {
	if err := credential.ValidateCredential(); err != nil {
		return nil, err
	}
	return s.repo.UpdateUserCredential(credential)
}
