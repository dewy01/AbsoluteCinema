package handler

import (
	"absolutecinema/internal/openapi/gen/seatgen"
	seat_service "absolutecinema/internal/service/seat"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type SeatHandler struct {
	Service seat_service.Service
}

func NewSeatHandler(svc seat_service.Service) *SeatHandler {
	return &SeatHandler{Service: svc}
}

// POST /seats/
func (h *SeatHandler) PostSeats(w http.ResponseWriter, r *http.Request) {
	var input seatgen.CreateSeatInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	seatOut, err := h.Service.Create(seat_service.CreateSeatInput{
		Number: input.Number,
		Row:    input.Row,
		RoomID: uuid.UUID(input.RoomID),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := seatgen.SeatOutput{
		Id:     &seatOut.ID,
		Number: &seatOut.Number,
		Row:    &seatOut.Row,
		RoomID: &seatOut.RoomID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// GET /seats/{id}
func (h *SeatHandler) GetSeatsId(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	seatOut, err := h.Service.GetByID(uuid.UUID(id))
	if err != nil {
		http.Error(w, "Seat not found", http.StatusNotFound)
		return
	}

	resp := seatgen.SeatOutput{
		Id:     &seatOut.ID,
		Number: &seatOut.Number,
		Row:    &seatOut.Row,
		RoomID: &seatOut.RoomID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// PUT /seats/{id}
func (h *SeatHandler) PutSeatsId(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var input seatgen.UpdateSeatInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.Service.Update(seat_service.UpdateSeatInput{
		ID:     id,
		Number: input.Number,
		Row:    input.Row,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DELETE /seats/{id}
func (h *SeatHandler) DeleteSeatsId(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	if err := h.Service.Delete(uuid.UUID(id)); err != nil {
		http.Error(w, "Failed to delete seat", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GET /seats/room/{roomID}
func (h *SeatHandler) GetSeatsRoomRoomID(w http.ResponseWriter, r *http.Request, roomID openapi_types.UUID) {
	seats, err := h.Service.GetByRoomID(uuid.UUID(roomID))
	if err != nil {
		http.Error(w, "Failed to get seats for room", http.StatusInternalServerError)
		return
	}

	var resp []seatgen.SeatOutput
	for _, s := range seats {
		resp = append(resp, seatgen.SeatOutput{
			Id:     &s.ID,
			Number: &s.Number,
			Row:    &s.Row,
			RoomID: &s.RoomID,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
