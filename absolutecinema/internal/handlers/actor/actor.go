package handler

import (
	"absolutecinema/internal/openapi/gen/actorgen"
	actor_service "absolutecinema/internal/service/actor"
	"absolutecinema/pkg/ptr"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type ActorHandler struct {
	Service actor_service.Service
}

func NewActorHandler(svc actor_service.Service) *ActorHandler {
	return &ActorHandler{Service: svc}
}

// GET /actors
func (h *ActorHandler) GetActors(w http.ResponseWriter, r *http.Request) {
	actors, err := h.Service.GetAll()
	if err != nil {
		http.Error(w, "Failed to get actors", http.StatusInternalServerError)
		return
	}

	var resp []actorgen.ActorOutput
	for _, a := range actors {
		resp = append(resp, actorgen.ActorOutput{
			Id:       &a.ID,
			Name:     &a.Name,
			MovieIDs: &a.MovieIDs,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// POST /actors
func (h *ActorHandler) PostActors(w http.ResponseWriter, r *http.Request) {
	var input actorgen.CreateActorInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	actorOut, err := h.Service.Create(actor_service.CreateActorInput{
		Name:     input.Name,
		MovieIDs: ptr.Deref(input.MovieIDs),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := actorgen.ActorOutput{
		Id:       &actorOut.ID,
		Name:     &actorOut.Name,
		MovieIDs: &actorOut.MovieIDs,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// GET /actors/{id}
func (h *ActorHandler) GetActorsId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	actorOut, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "Actor not found", http.StatusNotFound)
		return
	}

	resp := actorgen.ActorOutput{
		Id:       &actorOut.ID,
		Name:     &actorOut.Name,
		MovieIDs: &actorOut.MovieIDs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// PUT /actors/{id}
func (h *ActorHandler) PutActorsId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	var input actorgen.UpdateActorInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	actorOut, err := h.Service.Update(id, actor_service.UpdateActorInput{
		Name:     *input.Name,
		MovieIDs: ptr.Deref(input.MovieIDs),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := actorgen.ActorOutput{
		Id:       &actorOut.ID,
		Name:     &actorOut.Name,
		MovieIDs: &actorOut.MovieIDs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// DELETE /actors/{id}
func (h *ActorHandler) DeleteActorsId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "Failed to delete actor", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
