package reservedseat

import (
	"absolutecinema/internal/database/models"

	"github.com/google/uuid"
)

type ReservedSeat struct {
	ID            uuid.UUID
	ReservationID uuid.UUID
	SeatID        uuid.UUID
}

func ToDBReservedSeat(r *ReservedSeat) *models.ReservedSeat {
	return &models.ReservedSeat{
		ID:            r.ID,
		ReservationID: r.ReservationID,
		SeatID:        r.SeatID,
	}
}

func ToDomainReservedSeat(r *models.ReservedSeat) *ReservedSeat {
	return &ReservedSeat{
		ID:            r.ID,
		ReservationID: r.ReservationID,
		SeatID:        r.SeatID,
	}
}
