package handler

import (
	"absolutecinema/internal/openapi/gen/screeninggen"
	screening_service "absolutecinema/internal/service/screening"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
)

type ScreeningHandler struct {
	Service screening_service.Service
}

func NewScreeningHandler(svc screening_service.Service) *ScreeningHandler {
	return &ScreeningHandler{Service: svc}
}

// GET /screenings
func (h *ScreeningHandler) GetScreenings(w http.ResponseWriter, r *http.Request, params screeninggen.GetScreeningsParams) {
	time := optionalDateToTime(params.Day)
	screenings, err := h.Service.GetAll(time)
	if err != nil {
		http.Error(w, "Failed to get screenings", http.StatusInternalServerError)
		return
	}

	var resp []screeninggen.ScreeningOutput
	for _, s := range screenings {
		resp = append(resp, toScreeningOutput(s))
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

	writeJSON(w, http.StatusCreated, toScreeningOutput(*screeningOut))
}

// GET /screenings/{id}
func (h *ScreeningHandler) GetScreeningsId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	screeningOut, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "Screening not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusCreated, toScreeningOutput(*screeningOut))
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
func (h *ScreeningHandler) GetScreeningsMovieMovieID(w http.ResponseWriter, r *http.Request, movieID uuid.UUID, params screeninggen.GetScreeningsMovieMovieIDParams) {
	time := optionalDateToTime(params.Day)
	screenings, err := h.Service.GetByMovie(movieID, time)
	if err != nil {
		http.Error(w, "Failed to get screenings by movie ID", http.StatusInternalServerError)
		return
	}

	var resp []screeninggen.ScreeningOutput
	for _, s := range screenings {
		resp = append(resp, toScreeningOutput(s))
	}

	writeJSON(w, http.StatusOK, resp)
}

// GET /screenings/room/{roomID}
func (h *ScreeningHandler) GetScreeningsRoomRoomID(w http.ResponseWriter, r *http.Request, roomID uuid.UUID, params screeninggen.GetScreeningsRoomRoomIDParams) {
	time := optionalDateToTime(params.Day)
	screenings, err := h.Service.GetByRoom(roomID, time)
	if err != nil {
		http.Error(w, "Failed to get screenings by room ID", http.StatusInternalServerError)
		return
	}

	var resp []screeninggen.ScreeningOutput
	for _, s := range screenings {
		resp = append(resp, toScreeningOutput(s))
	}

	writeJSON(w, http.StatusOK, resp)
}

// GET /screenings/cinema/{cinemaID}
func (h *ScreeningHandler) GetScreeningsCinemaCinemaID(w http.ResponseWriter, r *http.Request, cinemaID types.UUID, params screeninggen.GetScreeningsCinemaCinemaIDParams) {
	time := optionalDateToTime(params.Day)
	screenings, err := h.Service.GetByCinema(cinemaID, time)
	if err != nil {
		http.Error(w, "Failed to get screenings by cinema ID", http.StatusInternalServerError)
		return
	}

	var resp []screeninggen.ScreeningOutput
	for _, s := range screenings {
		resp = append(resp, toScreeningOutput(s))
	}

	writeJSON(w, http.StatusOK, resp)
}
