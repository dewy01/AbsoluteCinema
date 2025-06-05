package handler

import (
	"absolutecinema/internal/openapi/gen/screeninggen"
	screening_service "absolutecinema/internal/service/screening"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type ScreeningHandler struct {
	Service screening_service.Service
}

func NewScreeningHandler(svc screening_service.Service) *ScreeningHandler {
	return &ScreeningHandler{Service: svc}
}

// GET /screenings
func (h *ScreeningHandler) GetScreenings(w http.ResponseWriter, r *http.Request) {
	screenings, err := h.Service.GetAll()
	if err != nil {
		http.Error(w, "Failed to get screenings", http.StatusInternalServerError)
		return
	}

	var resp []screeninggen.ScreeningOutput
	for _, s := range screenings {
		resp = append(resp, screeninggen.ScreeningOutput{
			Id:        &s.ID,
			MovieID:   &s.MovieID,
			RoomID:    &s.RoomID,
			StartTime: &s.StartTime,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// POST /screenings
func (h *ScreeningHandler) PostScreenings(w http.ResponseWriter, r *http.Request) {
	var input screeninggen.CreateScreeningInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	screeningOut, err := h.Service.Create(screening_service.CreateScreeningInput{
		MovieID:   uuid.UUID(input.MovieID),
		RoomID:    uuid.UUID(input.RoomID),
		StartTime: input.StartTime,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := screeninggen.ScreeningOutput{
		Id:        &screeningOut.ID,
		MovieID:   &screeningOut.MovieID,
		RoomID:    &screeningOut.RoomID,
		StartTime: &screeningOut.StartTime,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// GET /screenings/{id}
func (h *ScreeningHandler) GetScreeningsId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	screeningOut, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "Screening not found", http.StatusNotFound)
		return
	}

	resp := screeninggen.ScreeningOutput{
		Id:        &screeningOut.ID,
		MovieID:   &screeningOut.MovieID,
		RoomID:    &screeningOut.RoomID,
		StartTime: &screeningOut.StartTime,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// PUT /screenings/{id}
func (h *ScreeningHandler) PutScreeningsId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	var input screeninggen.PutScreeningsIdJSONBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.Service.UpdateStartTime(id, input.StartTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DELETE /screenings/{id}
func (h *ScreeningHandler) DeleteScreeningsId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "Failed to delete screening", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GET /screenings/movie/{movieID}
func (h *ScreeningHandler) GetScreeningsMovieMovieID(w http.ResponseWriter, r *http.Request, movieID uuid.UUID) {
	screenings, err := h.Service.GetByMovie(movieID)
	if err != nil {
		http.Error(w, "Failed to get screenings by movie ID", http.StatusInternalServerError)
		return
	}

	var resp []screeninggen.ScreeningOutput
	for _, s := range screenings {
		resp = append(resp, screeninggen.ScreeningOutput{
			Id:        &s.ID,
			MovieID:   &s.MovieID,
			RoomID:    &s.RoomID,
			StartTime: &s.StartTime,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GET /screenings/room/{roomID}
func (h *ScreeningHandler) GetScreeningsRoomRoomID(w http.ResponseWriter, r *http.Request, roomID uuid.UUID) {
	screenings, err := h.Service.GetByRoom(roomID)
	if err != nil {
		http.Error(w, "Failed to get screenings by room ID", http.StatusInternalServerError)
		return
	}

	var resp []screeninggen.ScreeningOutput
	for _, s := range screenings {
		resp = append(resp, screeninggen.ScreeningOutput{
			Id:        &s.ID,
			MovieID:   &s.MovieID,
			RoomID:    &s.RoomID,
			StartTime: &s.StartTime,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
