package database

import (
	"absolutecinema/internal/auth"
	"absolutecinema/internal/database/models"
	"log"
	"time"

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
	db.Exec("DELETE FROM reserved_seats")
	db.Exec("DELETE FROM reservations")
	db.Exec("DELETE FROM screenings")
	db.Exec("DELETE FROM seats")
	db.Exec("DELETE FROM rooms")
	db.Exec("DELETE FROM cinemas")
	db.Exec("DELETE FROM movie_actors")
	db.Exec("DELETE FROM actors")
	db.Exec("DELETE FROM movies")
	db.Exec("DELETE FROM users")

	passwordHash, err := bcrypt.GenerateFromPassword([]byte("user"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	user := models.User{
		Name:     "User",
		Email:    "user@example.com",
		Role:     auth.Role("user"),
		Password: string(passwordHash),
	}
	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	adminPassHash, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	admin := models.User{
		Name:     "Admin",
		Email:    "admin@example.com",
		Role:     auth.Role("admin"),
		Password: string(adminPassHash),
	}
	if err := db.Create(&admin).Error; err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	actor1 := models.Actor{Name: "Robert Downey Jr."}
	actor2 := models.Actor{Name: "Scarlett Johansson"}
	if err := db.Create(&actor1).Error; err != nil {
		log.Fatalf("Failed to create actor1: %v", err)
	}
	if err := db.Create(&actor2).Error; err != nil {
		log.Fatalf("Failed to create actor2: %v", err)
	}

	movie1 := models.Movie{
		Title:       "Avengers: Endgame",
		Director:    "Anthony and Joe Russo",
		Description: "Superhero film",
		PhotoPath:   "/images/avengers.jpg",
		Actors:      []models.Actor{actor1, actor2},
	}
	if err := db.Create(&movie1).Error; err != nil {
		log.Fatalf("Failed to create movie: %v", err)
	}

	cinema := models.Cinema{
		Name:    "Downtown Cinema",
		Address: "123 Main St",
	}
	if err := db.Create(&cinema).Error; err != nil {
		log.Fatalf("Failed to create cinema: %v", err)
	}

	room := models.Room{
		Name:     "Room 1",
		CinemaID: cinema.ID,
	}
	if err := db.Create(&room).Error; err != nil {
		log.Fatalf("Failed to create room: %v", err)
	}

	seats := []models.Seat{
		{Row: "A", Number: 1, RoomID: room.ID},
		{Row: "A", Number: 2, RoomID: room.ID},
		{Row: "B", Number: 1, RoomID: room.ID},
		{Row: "B", Number: 2, RoomID: room.ID},
	}
	for _, seat := range seats {
		if err := db.Create(&seat).Error; err != nil {
			log.Fatalf("Failed to create seat: %v", err)
		}
	}

	var seatA1 models.Seat
	if err := db.Where("row = ? AND number = ? AND room_id = ?", "A", 1, room.ID).First(&seatA1).Error; err != nil {
		log.Fatalf("Failed to find seat A1: %v", err)
	}

	screening := models.Screening{
		MovieID:   movie1.ID,
		RoomID:    room.ID,
		StartTime: time.Now().Add(24 * time.Hour),
	}
	if err := db.Create(&screening).Error; err != nil {
		log.Fatalf("Failed to create screening: %v", err)
	}

	reservation := models.Reservation{
		ScreeningID: screening.ID,
		GuestName:   "Guest User",
		GuestEmail:  "guest@example.com",
		PDFPath:     "/reservations/guest1.pdf",
	}
	if err := db.Create(&reservation).Error; err != nil {
		log.Fatalf("Failed to create reservation: %v", err)
	}

	reservedSeat := models.ReservedSeat{
		ReservationID: reservation.ID,
		SeatID:        seatA1.ID,
	}
	if err := db.Create(&reservedSeat).Error; err != nil {
		log.Fatalf("Failed to create reserved seat: %v", err)
	}

	log.Println("Database seeding completed.")
}
