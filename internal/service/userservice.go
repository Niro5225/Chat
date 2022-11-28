package service

import (
	"chat-app/internal/repository"
	"chat-app/models"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user models.User) (*models.User, error) {
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}
	return s.repo.CreateUser(user)
}

func (s *UserService) DeleteUser(id uint64) error {
	return s.repo.DeleteUser(id)
}

func (s *UserService) GetUser(id uint64) (*models.User, error) {
	return s.repo.GetUser(id)
}

func (s *UserService) GetUsers(userFilter *models.UserFilter) ([]models.User, error) {
	return s.repo.GetUsers(userFilter)
}

func (s *UserService) UpdateUser(user models.User) (*models.User, error) {
	return s.repo.UpdateUser(user)
}

func (s *UserService) CreateUserCredential(credential models.UserCredential) (*models.UserCredential, error) {
	if err := credential.ValidateCredential(); err != nil {
		return nil, err
	}
	credential.Encryption_password()
	return s.repo.CreateUserCredential(credential)
}

func (s *UserService) UpdateUserCredential(credential models.UserCredential) (*models.UserCredential, error) {
	if err := credential.ValidateCredential(); err != nil {
		return nil, err
	}
	credential.Encryption_password()
	return s.repo.UpdateUserCredential(credential)
}
