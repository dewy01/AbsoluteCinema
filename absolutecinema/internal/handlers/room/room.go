package handler

import (
	"absolutecinema/internal/openapi/gen/roomgen"
	room_service "absolutecinema/internal/service/room"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type RoomHandler struct {
	Service room_service.Service
}

func NewRoomHandler(svc room_service.Service) *RoomHandler {
	return &RoomHandler{Service: svc}
}

// POST /rooms/
func (h *RoomHandler) PostRooms(w http.ResponseWriter, r *http.Request) {
	var input roomgen.CreateRoomInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	roomOut, err := h.Service.Create(room_service.CreateRoomInput{
		CinemaID: input.CinemaID,
		Name:     input.Name,
		Seats:    convertSeatInputs(input.Seats),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := roomgen.RoomOutput{
		Id:       &roomOut.ID,
		CinemaID: &roomOut.CinemaID,
		Name:     &roomOut.Name,
		Seats:    convertSeatOutputs(roomOut.Seats),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// GET /rooms/cinema/{cinemaID}
func (h *RoomHandler) GetRoomsCinemaCinemaID(w http.ResponseWriter, r *http.Request, cinemaID uuid.UUID) {
	rooms, err := h.Service.GetByCinemaID(cinemaID)
	if err != nil {
		http.Error(w, "Failed to get rooms", http.StatusInternalServerError)
		return
	}

	var resp []roomgen.RoomOutput
	for _, room := range rooms {
		resp = append(resp, roomgen.RoomOutput{
			Id:       &room.ID,
			CinemaID: &room.CinemaID,
			Name:     &room.Name,
			Seats:    convertSeatOutputs(room.Seats),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// DELETE /rooms/{id}
func (h *RoomHandler) DeleteRoomsId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "Failed to delete room", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GET /rooms/{id}
func (h *RoomHandler) GetRoomsId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	room, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	resp := roomgen.RoomOutput{
		Id:       &room.ID,
		CinemaID: &room.CinemaID,
		Name:     &room.Name,
		Seats:    convertSeatOutputs(room.Seats),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// PUT /rooms/{id}
func (h *RoomHandler) PutRoomsId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	var input roomgen.UpdateRoomInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.Service.Update(room_service.UpdateRoomInput{
		ID:       id,
		CinemaID: input.CinemaID,
		Name:     input.Name,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
