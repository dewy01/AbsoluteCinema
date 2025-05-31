package handlers

import (
	actorHandler "absolutecinema/internal/handlers/actor"
	cinemaHandler "absolutecinema/internal/handlers/cinema"
	userHandler "absolutecinema/internal/handlers/user"
	"absolutecinema/internal/openapi/gen/actorgen"
	"absolutecinema/internal/openapi/gen/cinemagen"
	"absolutecinema/internal/openapi/gen/usergen"
	"absolutecinema/internal/service"
)

type Handlers struct {
	User   usergen.ServerInterface
	Actor  actorgen.ServerInterface
	Cinema cinemagen.ServerInterface
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		User:   userHandler.NewUserHandler(services.User),
		Actor:  actorHandler.NewActorHandler(services.Actor),
		Cinema: cinemaHandler.NewCinemaHandler(services.Cinema),
	}
}
