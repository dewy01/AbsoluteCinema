package user_service

import (
	"absolutecinema/internal/auth"
	"absolutecinema/internal/database/repository/user"
	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Register(input CreateUserInput) (*UserOutput, error)
	Login(ctx context.Context, input LoginInput) (*UserOutput, *auth.Session, error)
	Logout(ctx context.Context, sessionID uuid.UUID) error
	GetMe(ctx context.Context, sessionID uuid.UUID) (auth.SessionData, error)
	Update(id uuid.UUID, input UpdateUserInput) (*UserOutput, error)
	GetByID(id uuid.UUID) (*UserOutput, error)
	Delete(id uuid.UUID) error
}

type service struct {
	repo    user.Repository
	session *auth.Service
}

func NewUserService(repo user.Repository, sessionService *auth.Service) *service {
	return &service{repo: repo, session: sessionService}
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

func (s *service) Login(ctx context.Context, input LoginInput) (*UserOutput, *auth.Session, error) {
	user, err := s.repo.GetByEmail(input.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		return nil, nil, errors.New("invalid email or password")
	}

	session, err := s.session.New(ctx, auth.SessionData{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return nil, nil, err
	}

	return &UserOutput{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, &session, nil
}

func (s *service) Logout(ctx context.Context, sessionID uuid.UUID) error {
	session, err := s.session.Get(ctx, sessionID)
	if err != nil {
		return err
	}

	return s.session.Delete(ctx, session)
}

func (s *service) GetMe(ctx context.Context, sessionID uuid.UUID) (auth.SessionData, error) {
	session, err := s.session.Get(ctx, sessionID)
	if err != nil {
		return auth.SessionData{}, err
	}

	return session.Data, nil
}

func (s *service) Update(id uuid.UUID, input UpdateUserInput) (*UserOutput, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if input.Password != "" || input.ConfirmPassword != "" {
		if input.Password != input.ConfirmPassword {
			return nil, errors.New("passwords do not match")
		}
		hashedPw, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPw)
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return &UserOutput{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (s *service) GetByID(id uuid.UUID) (*UserOutput, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &UserOutput{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (s *service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
