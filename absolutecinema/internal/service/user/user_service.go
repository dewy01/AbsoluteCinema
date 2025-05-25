package user_service

import (
	"absolutecinema/internal/auth"
	"absolutecinema/internal/database/repository/user"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Register(input CreateUserInput) (*UserOutput, error)
	Login(input LoginInput) (*UserOutput, error)
}

type service struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) *service {
	return &service{repo: repo}
}

func (s *service) Register(input CreateUserInput) (*UserOutput, error) {
	if input.Password == "" || input.ConfirmPassword == "" {
		return nil, errors.New("password and confirmation are required")
	}

	if input.Password != input.ConfirmPassword {
		return nil, errors.New("passwords do not match")
	}

	existingUser, err := s.repo.GetByEmail(input.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	hashedPw, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &user.User{
		ID:       uuid.New(),
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPw),
		Role:     auth.User,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return &UserOutput{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (s *service) Login(input LoginInput) (*UserOutput, error) {
	user, err := s.repo.GetByEmail(input.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return &UserOutput{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}
