package reservation

import (
	"absolutecinema/internal/database/models"
	reservedseat "absolutecinema/internal/database/repository/reserved_seat"

	"github.com/google/uuid"
)

type Reservation struct {
	ID            uuid.UUID
	ScreeningID   uuid.UUID
	UserID        *uuid.UUID
	GuestName     string
	GuestEmail    string
	PDFPath       string
	ReservedSeats []reservedseat.ReservedSeat
}

func ToDBReservation(r *Reservation) *models.Reservation {
	seats := make([]models.ReservedSeat, len(r.ReservedSeats))
	for i, s := range r.ReservedSeats {
		seats[i] = models.ReservedSeat{
			ID:            s.ID,
			ReservationID: r.ID,
			SeatID:        s.SeatID,
		}
	}
	return &models.Reservation{
		ID:            r.ID,
		ScreeningID:   r.ScreeningID,
		UserID:        r.UserID,
		GuestName:     r.GuestName,
		GuestEmail:    r.GuestEmail,
		PDFPath:       r.PDFPath,
		ReservedSeats: seats,
	}
}

func ToDomainReservation(r *models.Reservation) *Reservation {
	seats := make([]reservedseat.ReservedSeat, len(r.ReservedSeats))
	for i, s := range r.ReservedSeats {
		seats[i] = *reservedseat.ToDomainReservedSeat(&s)
	}
	return &Reservation{
		ID:            r.ID,
		ScreeningID:   r.ScreeningID,
		UserID:        r.UserID,
		GuestName:     r.GuestName,
		GuestEmail:    r.GuestEmail,
		PDFPath:       r.PDFPath,
		ReservedSeats: seats,
	}
}
