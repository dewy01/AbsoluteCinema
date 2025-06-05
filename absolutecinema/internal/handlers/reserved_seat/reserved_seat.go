package reservedseathandler

import (
	"absolutecinema/internal/openapi/gen/reserved_seatgen"
	reservedseat_service "absolutecinema/internal/service/reserved_seat"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type ReservedSeatHandler struct {
	Service reservedseat_service.Service
}

func NewReservedSeatHandler(svc reservedseat_service.Service) *ReservedSeatHandler {
	return &ReservedSeatHandler{Service: svc}
}

// POST /reserved-seats/
func (h *ReservedSeatHandler) PostReservedSeats(w http.ResponseWriter, r *http.Request) {
	var input reserved_seatgen.CreateReservedSeatInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result, err := h.Service.Create(reservedseat_service.CreateReservedSeatInput{
		ReservationID: input.ReservationID,
		SeatID:        input.SeatID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := reserved_seatgen.ReservedSeatOutput{
		Id:            &result.ID,
		ReservationID: &result.ReservationID,
		SeatID:        &result.SeatID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// DELETE /reserved-seats/reservation/{reservationID}
func (h *ReservedSeatHandler) DeleteReservedSeatsReservationReservationID(w http.ResponseWriter, r *http.Request, reservationID uuid.UUID) {
	if err := h.Service.DeleteByReservationID(reservationID); err != nil {
		http.Error(w, "Failed to delete reserved seats", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GET /reserved-seats/reservation/{reservationID}
func (h *ReservedSeatHandler) GetReservedSeatsReservationReservationID(w http.ResponseWriter, r *http.Request, reservationID uuid.UUID) {
	seats, err := h.Service.GetByReservationID(reservationID)
	if err != nil {
		http.Error(w, "Failed to get reserved seats", http.StatusInternalServerError)
		return
	}

	var resp []reserved_seatgen.ReservedSeatOutput
	for _, s := range seats {
		resp = append(resp, reserved_seatgen.ReservedSeatOutput{
			Id:            &s.ID,
			ReservationID: &s.ReservationID,
			SeatID:        &s.SeatID,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// DELETE /reserved-seats/{id}
func (h *ReservedSeatHandler) DeleteReservedSeatsId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "Failed to delete reserved seat", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GET /reserved-seats/{id}
func (h *ReservedSeatHandler) GetReservedSeatsId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	seat, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "Reserved seat not found", http.StatusNotFound)
		return
	}

	resp := reserved_seatgen.ReservedSeatOutput{
		Id:            &seat.ID,
		ReservationID: &seat.ReservationID,
		SeatID:        &seat.SeatID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
