package handler

import (
	"absolutecinema/internal/openapi/gen/cinemagen"
	cinema_service "absolutecinema/internal/service/cinema"
	"absolutecinema/pkg/ptr"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type CinemaHandler struct {
	Service cinema_service.Service
}

func NewCinemaHandler(svc cinema_service.Service) *CinemaHandler {
	return &CinemaHandler{Service: svc}
}

// GET /cinemas
func (h *CinemaHandler) GetCinemas(w http.ResponseWriter, r *http.Request) {
	cinemas, err := h.Service.GetAll()
	if err != nil {
		http.Error(w, "Failed to get cinemas", http.StatusInternalServerError)
		return
	}

	var resp []cinemagen.CinemaOutput
	for _, c := range cinemas {
		resp = append(resp, cinemagen.CinemaOutput{
			Id:      &c.ID,
			Name:    &c.Name,
			Address: &c.Address,
			RoomIDs: &c.RoomIDs,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// POST /cinemas
func (h *CinemaHandler) PostCinemas(w http.ResponseWriter, r *http.Request) {
	var input cinemagen.CreateCinemaInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	cinemaOut, err := h.Service.Create(cinema_service.CreateCinemaInput{
		Name:    input.Name,
		Address: input.Address,
		RoomIDs: ptr.Deref(input.RoomIDs),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := cinemagen.CinemaOutput{
		Id:      &cinemaOut.ID,
		Name:    &cinemaOut.Name,
		Address: &cinemaOut.Address,
		RoomIDs: &cinemaOut.RoomIDs,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// GET /cinemas/{id}
func (h *CinemaHandler) GetCinemasId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	cinemaOut, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "Cinema not found", http.StatusNotFound)
		return
	}

	resp := cinemagen.CinemaOutput{
		Id:      &cinemaOut.ID,
		Name:    &cinemaOut.Name,
		Address: &cinemaOut.Address,
		RoomIDs: &cinemaOut.RoomIDs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// PUT /cinemas/{id}
func (h *CinemaHandler) PutCinemasId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	var input cinemagen.UpdateCinemaInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	cinemaOut, err := h.Service.Update(id, cinema_service.UpdateCinemaInput{
		Name:    ptr.Deref(input.Name),
		Address: ptr.Deref(input.Address),
		RoomIDs: ptr.Deref(input.RoomIDs),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := cinemagen.CinemaOutput{
		Id:      &cinemaOut.ID,
		Name:    &cinemaOut.Name,
		Address: &cinemaOut.Address,
		RoomIDs: &cinemaOut.RoomIDs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// DELETE /cinemas/{id}
func (h *CinemaHandler) DeleteCinemasId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "Failed to delete cinema", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
