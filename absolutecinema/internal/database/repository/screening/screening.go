package screening

import (
	"absolutecinema/internal/database/models"
	"absolutecinema/internal/database/repository/cinema"
	"absolutecinema/internal/database/repository/movie"
	"absolutecinema/internal/database/repository/room"
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID         uuid.UUID
	GuestName  string
	GuestEmail string
	PDFPath    string
}

type Screening struct {
	ID           uuid.UUID
	StartTime    time.Time
	Movie        movie.Movie
	Room         room.Room
	Reservations []Reservation
}

type ScreeningInput struct {
	ID        uuid.UUID
	MovieID   uuid.UUID
	RoomID    uuid.UUID
	StartTime time.Time
}

func ToDBScreening(s *ScreeningInput) *models.Screening {
	return &models.Screening{
		ID:        s.ID,
		MovieID:   s.MovieID,
		RoomID:    s.RoomID,
		StartTime: s.StartTime,
	}
}

func ToDomainScreening(s *models.Screening) *Screening {
	screening := &Screening{
		ID:        s.ID,
		StartTime: s.StartTime,
		Movie: movie.Movie{
			ID:          s.MovieID,
			Title:       s.Movie.Title,
			Director:    s.Movie.Director,
			Description: s.Movie.Description,
			PhotoPath:   s.Movie.PhotoPath,
		},
		Room: room.Room{
			ID:   s.RoomID,
			Name: s.Room.Name,
			Cinema: cinema.Cinema{
				ID:   s.Room.Cinema.ID,
				Name: s.Room.Cinema.Name,
			},
		},
	}

	if len(s.Reservations) > 0 {
		for _, r := range s.Reservations {
			screening.Reservations = append(screening.Reservations, Reservation{
				ID:         r.ID,
				GuestName:  r.GuestName,
				GuestEmail: r.GuestEmail,
				PDFPath:    r.PDFPath,
			})
		}
	}

	return screening
}

func mapScreenings(src []models.Screening) []Screening {
	result := make([]Screening, len(src))
	for i, s := range src {
		result[i] = *ToDomainScreening(&s)
	}
	return result
}

func dayBounds(day time.Time) (time.Time, time.Time) {
	year, month, dayOnly := day.Date()
	start := time.Date(year, month, dayOnly, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 1)
	return start, end
}
