package actor_service

import "github.com/google/uuid"

type CreateActorInput struct {
	Name     string
	MovieIDs []uuid.UUID
}

type UpdateActorInput struct {
	Name     string
	MovieIDs []uuid.UUID
}

type ActorOutput struct {
	ID       uuid.UUID
	Name     string
	MovieIDs []uuid.UUID
}
