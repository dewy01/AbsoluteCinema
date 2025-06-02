package service

import (
	"absolutecinema/internal/auth"
	"absolutecinema/internal/database/repository"
	actorService "absolutecinema/internal/service/actor"
	cinemaService "absolutecinema/internal/service/cinema"
	userService "absolutecinema/internal/service/user"
)

type Services struct {
	User   userService.Service
	Actor  actorService.Service
	Cinema cinemaService.Service
}

func NewServices(repos *repository.Repositories, sessionService *auth.Service) *Services {
	return &Services{
		User:   userService.NewUserService(repos.User, sessionService),
		Actor:  actorService.NewActorService(repos.Actor),
		Cinema: cinemaService.NewCinemaService(repos.Cinema),
	}
}
