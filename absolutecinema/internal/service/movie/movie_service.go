package movie_service

import (
	"absolutecinema/internal/database/repository/movie"
	"absolutecinema/pkg/fsystem"
	"errors"
	"fmt"
	"path/filepath"

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
	fs   fsystem.FileStorage
}

func NewMovieService(repo movie.Repository, fs fsystem.FileStorage) *service {
	return &service{
		repo: repo,
		fs:   fs,
	}
}

func (s *service) Create(input CreateMovieInput) (*MovieOutput, error) {
	if input.Title == "" {
		return nil, errors.New("title is required")
	}
	if input.Director == "" {
		return nil, errors.New("director is required")
	}

	id := uuid.New()
	storagePath := fmt.Sprintf("resources/movies/%s/", id.String())
	photoFilename := "photo" + filepath.Ext(input.Photo.Filename())

	reader, err := input.Photo.Reader()
	if err != nil {
		return nil, fmt.Errorf("failed to read photo file: %w", err)
	}
	defer reader.Close()

	if err := s.fs.SaveMovieFile(storagePath, photoFilename, reader); err != nil {
		return nil, fmt.Errorf("failed to save photo: %w", err)
	}

	fullPath := filepath.Join(storagePath, photoFilename)

	movieDomain := &movie.Movie{
		ID:          id,
		Title:       input.Title,
		Director:    input.Director,
		Description: input.Description,
		PhotoPath:   fullPath,
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
	if input.ActorIDs != nil {
		movieDomain.ActorIDs = input.ActorIDs
	}

	if input.Photo.FileSize() > 0 {
		storagePath := fmt.Sprintf("resources/movies/%s/", id.String())
		newPhotoFilename := "photo" + filepath.Ext(input.Photo.Filename())

		if movieDomain.PhotoPath != "" {
			oldPhotoFilename := filepath.Base(movieDomain.PhotoPath)
			_ = s.fs.RemoveMovieFile(id.String(), oldPhotoFilename)
		}

		reader, err := input.Photo.Reader()
		if err != nil {
			return nil, fmt.Errorf("read new photo file: %w", err)
		}
		defer reader.Close()

		if err := s.fs.SaveMovieFile(id.String(), newPhotoFilename, reader); err != nil {
			return nil, fmt.Errorf("save new photo: %w", err)
		}

		movieDomain.PhotoPath = filepath.Join(storagePath, newPhotoFilename)
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
