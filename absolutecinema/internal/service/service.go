package service

import (
	"absolutecinema/internal/auth"
	"absolutecinema/internal/database/repository"
	userService "absolutecinema/internal/service/user"
)

type Services struct {
	User userService.Service
}

func NewServices(repos *repository.Repositories, sessionService *auth.Service) *Services {
	return &Services{
		User: userService.NewUserService(repos.User, sessionService),
	}
}
