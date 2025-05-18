package repository

import (
	"absolutecinema/internal/database/repository/actor"
	"absolutecinema/internal/database/repository/cinema"
	"absolutecinema/internal/database/repository/movie"
	"absolutecinema/internal/database/repository/reservation"
	reservedseat "absolutecinema/internal/database/repository/reserved_seat"
	"absolutecinema/internal/database/repository/room"
	"absolutecinema/internal/database/repository/screening"
	"absolutecinema/internal/database/repository/seat"
	"absolutecinema/internal/database/repository/user"

	"gorm.io/gorm"
)

type Repositories struct {
	User         user.Repository
	Cinema       cinema.Repository
	Room         room.Repository
	Seat         seat.Repository
	Movie        movie.Repository
	Actor        actor.Repository
	Screening    screening.Repository
	Reservation  reservation.Repository
	ReservedSeat reservedseat.Repository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User:         user.NewUserRepository(db),
		Cinema:       cinema.NewCinemaRepository(db),
		Room:         room.NewRoomRepository(db),
		Seat:         seat.NewSeatRepository(db),
		Movie:        movie.NewMovieRepository(db),
		Actor:        actor.NewActorRepository(db),
		Screening:    screening.NewScreeningRepository(db),
		Reservation:  reservation.NewReservationRepository(db),
		ReservedSeat: reservedseat.NewReservedSeatRepository(db),
	}
}
