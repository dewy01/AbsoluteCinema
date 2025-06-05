package service

import (
	"absolutecinema/internal/auth"
	"absolutecinema/internal/database/repository"
	actorService "absolutecinema/internal/service/actor"
	cinemaService "absolutecinema/internal/service/cinema"
	movieService "absolutecinema/internal/service/movie"
	reservationService "absolutecinema/internal/service/reservation"
	reserved_seat_Service "absolutecinema/internal/service/reserved_seat"
	room_service "absolutecinema/internal/service/room"
	screening_service "absolutecinema/internal/service/screening"
	seat_service "absolutecinema/internal/service/seat"
	userService "absolutecinema/internal/service/user"
)

type Services struct {
	User         userService.Service
	Actor        actorService.Service
	Cinema       cinemaService.Service
	Movie        movieService.Service
	Reservation  reservationService.Service
	ReservedSeat reserved_seat_Service.Service
	Room         room_service.Service
	Screening    screening_service.Service
	Seat         seat_service.Service
}

func NewServices(repos *repository.Repositories, sessionService *auth.Service) *Services {
	return &Services{
		User:         userService.NewUserService(repos.User, sessionService),
		Actor:        actorService.NewActorService(repos.Actor),
		Cinema:       cinemaService.NewCinemaService(repos.Cinema),
		Movie:        movieService.NewMovieService(repos.Movie),
		Reservation:  reservationService.NewReservationService(repos.Reservation),
		ReservedSeat: reserved_seat_Service.NewReservedSeatService(repos.ReservedSeat),
		Room:         room_service.NewRoomService(repos.Room),
		Screening:    screening_service.NewScreeningService(repos.Screening),
		Seat:         seat_service.NewSeatService(repos.Seat),
	}
}
