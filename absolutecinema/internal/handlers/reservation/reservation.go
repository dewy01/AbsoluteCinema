package reservationhandler

import (
	"encoding/json"
	"net/http"

	"absolutecinema/internal/openapi/gen/reservationgen"
	reservation_service "absolutecinema/internal/service/reservation"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Handler struct {
	Service reservation_service.Service
}

func NewReservationHandler(s reservation_service.Service) *Handler {
	return &Handler{
		Service: s,
	}
}

// POST /reservations/
func (h *Handler) PostReservations(w http.ResponseWriter, r *http.Request) {
	var input reservationgen.CreateReservationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	serviceInput := reservation_service.CreateReservationInput{
		ScreeningID:   input.ScreeningID,
		UserID:        input.UserID,
		GuestName:     *input.GuestName,
		GuestEmail:    string(*input.GuestEmail),
		ReservedSeats: make([]reservation_service.ReservedSeat, len(input.ReservedSeats)),
	}

	for i, seat := range input.ReservedSeats {
		serviceInput.ReservedSeats[i] = reservation_service.ReservedSeat{
			SeatID: uuid.UUID(seat.SeatID),
		}
	}

	out, err := h.Service.Create(serviceInput)
	if err != nil {
		http.Error(w, "failed to create reservation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(out)
}

// GET /reservations/user/{userID}
func (h *Handler) GetReservationsUserUserID(w http.ResponseWriter, r *http.Request, userID openapi_types.UUID) {
	reservations, err := h.Service.GetByUserID(userID)
	if err != nil {
		http.Error(w, "failed to get reservations: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reservations)
}

// DELETE /reservations/{id}
func (h *Handler) DeleteReservationsId(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	err := h.Service.Delete(id)
	if err != nil {
		http.Error(w, "failed to delete reservation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /reservations/{id}
func (h *Handler) GetReservationsId(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	res, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "failed to get reservation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// PUT /reservations/{id}
func (h *Handler) PutReservationsId(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var body reservationgen.PutReservationsIdJSONBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid JSON body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := h.Service.UpdatePDF(id, body.PdfPath)
	if err != nil {
		http.Error(w, "failed to update PDF path: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
