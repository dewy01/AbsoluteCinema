package handlers

import (
	actorHandler "absolutecinema/internal/handlers/actor"
	cinemaHandler "absolutecinema/internal/handlers/cinema"
	movieHandler "absolutecinema/internal/handlers/movie"
	reservationHandler "absolutecinema/internal/handlers/reservation"
	reserved_seat_Handler "absolutecinema/internal/handlers/reserved_seat"
	roomHandler "absolutecinema/internal/handlers/room"
	screeningHandler "absolutecinema/internal/handlers/screening"
	seatHandler "absolutecinema/internal/handlers/seat"
	userHandler "absolutecinema/internal/handlers/user"
	"absolutecinema/internal/openapi/gen/actorgen"
	"absolutecinema/internal/openapi/gen/cinemagen"
	"absolutecinema/internal/openapi/gen/moviegen"
	"absolutecinema/internal/openapi/gen/reservationgen"
	"absolutecinema/internal/openapi/gen/reserved_seatgen"
	"absolutecinema/internal/openapi/gen/roomgen"
	"absolutecinema/internal/openapi/gen/screeninggen"
	"absolutecinema/internal/openapi/gen/seatgen"
	"absolutecinema/internal/openapi/gen/usergen"
	"absolutecinema/internal/service"
)

type Handlers struct {
	User         usergen.ServerInterface
	Actor        actorgen.ServerInterface
	Cinema       cinemagen.ServerInterface
	Movie        moviegen.ServerInterface
	Reservation  reservationgen.ServerInterface
	ReservedSeat reserved_seatgen.ServerInterface
	Room         roomgen.ServerInterface
	Screening    screeninggen.ServerInterface
	Seat         seatgen.ServerInterface
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		User:         userHandler.NewUserHandler(services.User),
		Actor:        actorHandler.NewActorHandler(services.Actor),
		Cinema:       cinemaHandler.NewCinemaHandler(services.Cinema),
		Movie:        movieHandler.NewMovieHandler(services.Movie),
		Reservation:  reservationHandler.NewReservationHandler(services.Reservation),
		ReservedSeat: reserved_seat_Handler.NewReservedSeatHandler(services.ReservedSeat),
		Room:         roomHandler.NewRoomHandler(services.Room),
		Screening:    screeningHandler.NewScreeningHandler(services.Screening),
		Seat:         seatHandler.NewSeatHandler(services.Seat),
	}
}
