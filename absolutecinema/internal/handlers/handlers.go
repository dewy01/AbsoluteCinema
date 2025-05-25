package handlers

import (
	userHandler "absolutecinema/internal/handlers/user"
	"absolutecinema/internal/openapi/gen/usergen"
	"absolutecinema/internal/service"
)

type Handlers struct {
	User usergen.ServerInterface
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		User: userHandler.NewUserHandler(services.User),
	}
}
