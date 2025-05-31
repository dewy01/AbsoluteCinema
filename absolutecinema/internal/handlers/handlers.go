package handlers

import (
	actorHandler "absolutecinema/internal/handlers/actor"
	userHandler "absolutecinema/internal/handlers/user"
	"absolutecinema/internal/openapi/gen/actorgen"
	"absolutecinema/internal/openapi/gen/usergen"
	"absolutecinema/internal/service"
)

type Handlers struct {
	User  usergen.ServerInterface
	Actor actorgen.ServerInterface
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		User:  userHandler.NewUserHandler(services.User),
		Actor: actorHandler.NewActorHandler(services.Actor),
	}
}
