package chat_domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserService interface {
	GetUser(id uint64) (*User, error)
	GetUsers(userFilter *UserFilter) ([]User, error)
	SignIn(email, password string) (*User, string, error)                   //LOGIN
	SignUp(user User, userCredential UserCredential) (*User, string, error) //REG
	CreateUser(user User) (*User, error)
	UpdateUser(user User) (*User, error)
	CreateUserCredential(credential UserCredential) (*UserCredential, error)
	GetUserCredential(email string) (*UserCredential, error)
	UpdateUserCredential(credential UserCredential) (*UserCredential, error)
	DeleteUser(id uint64) error
}

const tokenKey = "ndkasd#nasjnda#kndkj"

type UserServiceImp struct {
	repo UserRepository
}

func NewUserServiceImp(repo UserRepository) *UserServiceImp {
	return &UserServiceImp{repo: repo}
}

func (s *UserServiceImp) CreateUser(user User) (*User, error) {
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}
	return s.repo.CreateUser(user)
}

func (s *UserServiceImp) GetUserCredential(email string) (*UserCredential, error) {
	return s.repo.GetUserCredential(email)
}

func (s *UserServiceImp) DeleteUser(id uint64) error {
	return s.repo.DeleteUser(id)
}

func (s *UserServiceImp) GetUser(id uint64) (*User, error) {
	return s.repo.GetUser(id)
}

func (s *UserServiceImp) GetUsers(userFilter *UserFilter) ([]User, error) {
	return s.repo.GetUsers(userFilter)
}

type NewTokenClaims struct {
	jwt.StandardClaims
	UserId uint64
}

func GenerateToken(userId uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix()},
		userId})

	return token.SignedString([]byte(tokenKey))
}

func (s *UserServiceImp) SignIn(email, password string) (*User, string, error) {
	uc, err := s.GetUserCredential(email)
	if err != nil {
		return nil, "", err
	}
	err = uc.CheckPasswords(password)
	if err != nil {
		return nil, "", err
	}
	user, err := s.GetUser(uc.ID)
	if err != nil {
		return nil, "", err
	}
	token, err := GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *UserServiceImp) SignUp(user User, userCredential UserCredential) (*User, string, error) {
	token, err := GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}
	return &user, token, nil
}

func (s *UserServiceImp) UpdateUser(user User) (*User, error) {
	return s.repo.UpdateUser(user)
}

func (s *UserServiceImp) CreateUserCredential(credential UserCredential) (*UserCredential, error) {
	if err := credential.ValidateCredential(); err != nil {
		return nil, err
	}
	return s.repo.CreateUserCredential(credential)
}

func (s *UserServiceImp) UpdateUserCredential(credential UserCredential) (*UserCredential, error) {
	if err := credential.ValidateCredential(); err != nil {
		return nil, err
	}
	return s.repo.UpdateUserCredential(credential)
}