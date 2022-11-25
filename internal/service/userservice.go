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
	return s.repo.CreateUser(user)
}

func (s *UserService) DeleteUser(id uint64) error {
	return s.repo.DeleteUser(id)
}
