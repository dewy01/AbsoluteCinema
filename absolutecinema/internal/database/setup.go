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
	"gorm.io/gorm"
)

func (db *Database) Gorm() *gorm.DB {
	return db.DB
}

func (db *Database) Setup(seed bool) error {
	db.initUUID()
	if err := db.migrate(); err != nil {
		return err
	}

	if seed {
		db.seed()
	}

	return nil
}

func (db *Database) migrate() error {
	return db.AutoMigrate(
		&models.User{},
		&models.Cinema{},
		&models.Room{},
		&models.Seat{},
		&models.Actor{},
		&models.Movie{},
		&models.Screening{},
		&models.Reservation{},
		&models.ReservedSeat{},
	)
}

func (db *Database) initUUID() {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
}

func (db *Database) isEmpty() bool {
	var count int64
	db.Model(&models.Cinema{}).Count(&count)
	if count > 0 {
		return false
	}

	return true
}

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

	actorNames := []string{"Robert Downey Jr.", "Scarlett Johansson", "Chris Evans", "Tom Holland", "Gal Gadot", "Chris Hemsworth"}
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
		{"Avengers: Endgame", "Russo Brothers", "Superhero film", "/movies/avengers.jpg", []int{0, 1, 2}},
		{"Wonder Woman", "Patty Jenkins", "Amazonian warrior", "/movies/wonder.jpg", []int{4}},
		{"Spider-Man: Homecoming", "Jon Watts", "Teen superhero", "/movies/spiderman.jpg", []int{3}},
		{"Thor: Ragnarok", "Taika Waititi", "Thor in space", "/movies/thor.jpg", []int{5}},
		{"Black Widow", "Cate Shortland", "Spy action film", "/movies/blackwidow.jpg", []int{1}},
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

		for r := 1; r <= 2; r++ {
			room := models.Room{Name: fmt.Sprintf("Room %d", r), CinemaID: cinema.ID}
			if err := db.Create(&room).Error; err != nil {
				log.Fatalf("Failed to create room %d in cinema %d: %v", r, c, err)
			}

			for row := 'A'; row <= 'C'; row++ {
				for seatNum := 1; seatNum <= 4; seatNum++ {
					seat := models.Seat{Row: string(row), Number: seatNum, RoomID: room.ID}
					if err := db.Create(&seat).Error; err != nil {
						log.Fatalf("Failed to create seat %c%d in room %d: %v", row, seatNum, r, err)
					}
				}
			}

			for i := 0; i < 2; i++ {
				movie := movies[(c+r+i)%len(movies)]
				start := time.Now().Add(time.Duration((c+r+i)*3) * time.Hour)
				screening := models.Screening{MovieID: movie.ID, RoomID: room.ID, StartTime: start}
				if err := db.Create(&screening).Error; err != nil {
					log.Fatalf("Failed to create screening: %v", err)
				}

				if (c+r+i)%2 == 0 {
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

	log.Println("Expanded database seeding completed.")
}
