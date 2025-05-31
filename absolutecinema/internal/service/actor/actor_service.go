package actor_service

import (
	"absolutecinema/internal/database/repository/actor"
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	Create(input CreateActorInput) (*ActorOutput, error)
	GetByID(id uuid.UUID) (*ActorOutput, error)
	GetAll() ([]ActorOutput, error)
	Update(id uuid.UUID, input UpdateActorInput) (*ActorOutput, error)
	Delete(id uuid.UUID) error
}

type service struct {
	repo actor.Repository
}

func NewActorService(repo actor.Repository) *service {
	return &service{repo: repo}
}

func (s *service) Create(input CreateActorInput) (*ActorOutput, error) {
	if input.Name == "" {
		return nil, errors.New("name is required")
	}

	actorDomain := &actor.Actor{
		ID:       uuid.New(),
		Name:     input.Name,
		MovieIDs: input.MovieIDs,
	}

	err := s.repo.Create(actorDomain)
	if err != nil {
		return nil, err
	}

	return &ActorOutput{
		ID:       actorDomain.ID,
		Name:     actorDomain.Name,
		MovieIDs: actorDomain.MovieIDs,
	}, nil
}

func (s *service) GetByID(id uuid.UUID) (*ActorOutput, error) {
	actorDomain, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &ActorOutput{
		ID:       actorDomain.ID,
		Name:     actorDomain.Name,
		MovieIDs: actorDomain.MovieIDs,
	}, nil
}

func (s *service) GetAll() ([]ActorOutput, error) {
	actors, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	outputs := make([]ActorOutput, len(actors))
	for i, a := range actors {
		outputs[i] = ActorOutput{
			ID:       a.ID,
			Name:     a.Name,
			MovieIDs: a.MovieIDs,
		}
	}
	return outputs, nil
}

func (s *service) Update(id uuid.UUID, input UpdateActorInput) (*ActorOutput, error) {
	actorDomain, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if input.Name != "" {
		actorDomain.Name = input.Name
	}
	if input.MovieIDs != nil {
		actorDomain.MovieIDs = input.MovieIDs
	}

	if err := s.repo.Update(actorDomain); err != nil {
		return nil, err
	}

	return &ActorOutput{
		ID:       actorDomain.ID,
		Name:     actorDomain.Name,
		MovieIDs: actorDomain.MovieIDs,
	}, nil
}

func (s *service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
