package movie_service

import (
	"absolutecinema/internal/database/repository/movie"
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	Create(input CreateMovieInput) (*MovieOutput, error)
	GetByID(id uuid.UUID) (*MovieOutput, error)
	GetAll() ([]MovieOutput, error)
	Update(id uuid.UUID, input UpdateMovieInput) (*MovieOutput, error)
	Delete(id uuid.UUID) error
}

type service struct {
	repo movie.Repository
}

func NewMovieService(repo movie.Repository) *service {
	return &service{repo: repo}
}

func (s *service) Create(input CreateMovieInput) (*MovieOutput, error) {
	if input.Title == "" {
		return nil, errors.New("title is required")
	}
	if input.Director == "" {
		return nil, errors.New("director is required")
	}

	movieDomain := &movie.Movie{
		ID:          uuid.New(),
		Title:       input.Title,
		Director:    input.Director,
		Description: input.Description,
		PhotoPath:   input.PhotoPath,
		ActorIDs:    input.ActorIDs,
	}

	if err := s.repo.Create(movieDomain); err != nil {
		return nil, err
	}

	return &MovieOutput{
		ID:          movieDomain.ID,
		Title:       movieDomain.Title,
		Director:    movieDomain.Director,
		Description: movieDomain.Description,
		PhotoPath:   movieDomain.PhotoPath,
		ActorIDs:    movieDomain.ActorIDs,
	}, nil
}

func (s *service) GetByID(id uuid.UUID) (*MovieOutput, error) {
	movieDomain, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &MovieOutput{
		ID:          movieDomain.ID,
		Title:       movieDomain.Title,
		Director:    movieDomain.Director,
		Description: movieDomain.Description,
		PhotoPath:   movieDomain.PhotoPath,
		ActorIDs:    movieDomain.ActorIDs,
	}, nil
}

func (s *service) GetAll() ([]MovieOutput, error) {
	movies, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	outputs := make([]MovieOutput, len(movies))
	for i, m := range movies {
		outputs[i] = MovieOutput{
			ID:          m.ID,
			Title:       m.Title,
			Director:    m.Director,
			Description: m.Description,
			PhotoPath:   m.PhotoPath,
			ActorIDs:    m.ActorIDs,
		}
	}

	return outputs, nil
}

func (s *service) Update(id uuid.UUID, input UpdateMovieInput) (*MovieOutput, error) {
	movieDomain, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if input.Title != "" {
		movieDomain.Title = input.Title
	}
	if input.Director != "" {
		movieDomain.Director = input.Director
	}
	if input.Description != "" {
		movieDomain.Description = input.Description
	}
	if input.PhotoPath != "" {
		movieDomain.PhotoPath = input.PhotoPath
	}
	if input.ActorIDs != nil {
		movieDomain.ActorIDs = input.ActorIDs
	}

	if err := s.repo.Update(movieDomain); err != nil {
		return nil, err
	}

	return &MovieOutput{
		ID:          movieDomain.ID,
		Title:       movieDomain.Title,
		Director:    movieDomain.Director,
		Description: movieDomain.Description,
		PhotoPath:   movieDomain.PhotoPath,
		ActorIDs:    movieDomain.ActorIDs,
	}, nil
}

func (s *service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
