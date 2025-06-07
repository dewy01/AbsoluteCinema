package database

import (
	"absolutecinema/internal/auth"
	"absolutecinema/internal/database/models"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (db *Database) seed() {
	tables := []string{
		"reserved_seats", "reservations", "screenings", "seats", "rooms", "cinemas",
		"movie_actors", "actors", "movies", "users",
	}
	for _, table := range tables {
		db.Exec(fmt.Sprintf("DELETE FROM %s", table))
	}

	passwords := map[string]string{"user": "user", "admin": "admin", "john": "123456", "jane": "password"}
	users := []models.User{}
	for name, pass := range passwords {
		hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		role := auth.Role("user")
		if name == "admin" {
			role = auth.Role("admin")
		}
		user := models.User{Name: strings.Title(name), Email: fmt.Sprintf("%s@example.com", name), Role: role, Password: string(hash)}
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Failed to create user %s: %v", name, err)
		}
		users = append(users, user)
	}

	actorNames := []string{
		"Robert Downey Jr.", "Scarlett Johansson", "Chris Evans", "Tom Holland",
		"Gal Gadot", "Chris Hemsworth", "Emma Stone", "Ryan Reynolds", "Zendaya", "Florence Pugh",
	}
	actors := []models.Actor{}
	for _, name := range actorNames {
		actor := models.Actor{Name: name}
		if err := db.Create(&actor).Error; err != nil {
			log.Fatalf("Failed to create actor %s: %v", name, err)
		}
		actors = append(actors, actor)
	}

	movieData := []struct {
		Title       string
		Director    string
		Description string
		PhotoPath   string
		ActorIdx    []int
	}{
		{"Avengers: Endgame", "Russo Brothers", "Superhero film", "/seeded/avengers.jpg", []int{0, 1, 2}},
		{"Wonder Woman", "Patty Jenkins", "Amazonian warrior", "/seeded/wonder.jpg", []int{4}},
		{"Spider-Man: Homecoming", "Jon Watts", "Teen superhero", "/seeded/spiderman.jpg", []int{3, 8}},
		{"Thor: Ragnarok", "Taika Waititi", "Thor in space", "/seeded/thor.jpg", []int{5}},
		{"Black Widow", "Cate Shortland", "Spy action film", "/seeded/blackwidow.jpg", []int{1, 9}},
		{"La La Land", "Damien Chazelle", "Jazz romance", "/seeded/lalaland.jpg", []int{6, 7}},
		{"Deadpool", "Tim Miller", "Antihero action comedy", "/seeded/deadpool.jpg", []int{7}},
		{"Dune", "Denis Villeneuve", "Sci-fi epic", "/seeded/dune.jpg", []int{9}},
		{"The Greatest Showman", "Michael Gracey", "Musical", "/seeded/showman.jpg", []int{6}},
		{"No Way Home", "Jon Watts", "Multiverse Spidey", "/seeded/nowayhome.jpg", []int{3, 8}},
	}

	movies := []models.Movie{}
	for _, m := range movieData {
		cast := []models.Actor{}
		for _, idx := range m.ActorIdx {
			cast = append(cast, actors[idx])
		}
		movie := models.Movie{Title: m.Title, Director: m.Director, Description: m.Description, PhotoPath: m.PhotoPath, Actors: cast}
		if err := db.Create(&movie).Error; err != nil {
			log.Fatalf("Failed to create movie %s: %v", m.Title, err)
		}
		movies = append(movies, movie)
	}

	for c := 1; c <= 3; c++ {
		cinema := models.Cinema{Name: fmt.Sprintf("Cinema %d", c), Address: fmt.Sprintf("%d Main Street", 100+c)}
		if err := db.Create(&cinema).Error; err != nil {
			log.Fatalf("Failed to create cinema %d: %v", c, err)
		}

		for r := 1; r <= 3; r++ {
			room := models.Room{Name: fmt.Sprintf("Room %d", r), CinemaID: cinema.ID}
			if err := db.Create(&room).Error; err != nil {
				log.Fatalf("Failed to create room %d in cinema %d: %v", r, c, err)
			}

			for row := 'A'; row <= 'F'; row++ {
				for seatNum := 1; seatNum <= 8; seatNum++ {
					seat := models.Seat{Row: string(row), Number: seatNum, RoomID: room.ID}
					if err := db.Create(&seat).Error; err != nil {
						log.Fatalf("Failed to create seat %c%d in room %d: %v", row, seatNum, r, err)
					}
				}
			}

			for dayOffset := 0; dayOffset < 7; dayOffset++ {
				for show := 0; show < 3; show++ {
					movie := movies[(c+r+dayOffset+show)%len(movies)]
					start := time.Now().Truncate(24 * time.Hour).
						Add(time.Duration(dayOffset) * 24 * time.Hour).
						Add(time.Duration(10+show*3) * time.Hour)

					screening := models.Screening{MovieID: movie.ID, RoomID: room.ID, StartTime: start}
					if err := db.Create(&screening).Error; err != nil {
						log.Fatalf("Failed to create screening: %v", err)
					}

					if (c+r+show)%3 == 0 {
						res := models.Reservation{
							ScreeningID: screening.ID,
							UserID:      &users[0].ID,
							PDFPath:     fmt.Sprintf("/reservations/%s.pdf", uuid.New().String()),
							GuestName:   "Sample Guest",
							GuestEmail:  "guest@example.com",
						}
						if err := db.Create(&res).Error; err != nil {
							log.Fatalf("Failed to create reservation: %v", err)
						}

						var seat models.Seat
						if err := db.Where("room_id = ?", room.ID).First(&seat).Error; err == nil {
							resSeat := models.ReservedSeat{ReservationID: res.ID, SeatID: seat.ID}
							if err := db.Create(&resSeat).Error; err != nil {
								log.Fatalf("Failed to create reserved seat: %v", err)
							}
						}
					}
				}
			}
		}
	}

	log.Println("Database seeding completed.")
}
