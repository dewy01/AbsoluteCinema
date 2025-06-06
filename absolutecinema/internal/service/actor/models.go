package actor_service

import "github.com/google/uuid"

type CreateActorInput struct {
	Name     string      `json:"name"`
	MovieIDs []uuid.UUID `json:"movieIDs"`
}

type UpdateActorInput struct {
	Name     string      `json:"name"`
	MovieIDs []uuid.UUID `json:"movieIDs"`
}

type ActorOutput struct {
	ID       uuid.UUID   `json:"id"`
	Name     string      `json:"name"`
	MovieIDs []uuid.UUID `json:"movieIDs"`
}
