package handler

import (
	"absolutecinema/internal/openapi/gen/moviegen"
	movie_service "absolutecinema/internal/service/movie"
	"absolutecinema/pkg/ptr"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type MovieHandler struct {
	Service movie_service.Service
}

func NewMovieHandler(svc movie_service.Service) *MovieHandler {
	return &MovieHandler{Service: svc}
}

// GET /movies
func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Service.GetAll()
	if err != nil {
		http.Error(w, "Failed to get movies", http.StatusInternalServerError)
		return
	}

	var resp []moviegen.MovieOutput
	for _, m := range movies {
		resp = append(resp, ToMovieOutput(m))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// POST /movies
func (h *MovieHandler) PostMovies(w http.ResponseWriter, r *http.Request) {
	var input moviegen.CreateMovieInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	movieOut, err := h.Service.Create(movie_service.CreateMovieInput{
		Title:       input.Title,
		Director:    input.Director,
		Description: ptr.Deref(input.Description),
		Photo:       input.Photo,
		ActorIDs:    ptr.Deref(input.ActorIDs),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := ToMovieOutput(*movieOut)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// GET /movies/{id}
func (h *MovieHandler) GetMoviesId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	movieOut, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	resp := ToMovieOutput(*movieOut)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// PUT /movies/{id}
func (h *MovieHandler) PutMoviesId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	var input moviegen.UpdateMovieInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	movieOut, err := h.Service.Update(id, movie_service.UpdateMovieInput{
		Title:       ptr.Deref(input.Title),
		Director:    ptr.Deref(input.Director),
		Description: ptr.Deref(input.Description),
		Photo:       *input.Photo,
		ActorIDs:    ptr.Deref(input.ActorIDs),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := ToMovieOutput(*movieOut)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// DELETE /movies/{id}
func (h *MovieHandler) DeleteMoviesId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "Failed to delete movie", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
