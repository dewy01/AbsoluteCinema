package service

import (
	"absolutecinema/internal/database/repository"
	userService "absolutecinema/internal/service/user"
)

type Services struct {
	User userService.Service
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		User: userService.NewUserService(repos.User),
	}
}
