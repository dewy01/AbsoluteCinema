package reservation_service

import (
	"bytes"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
)

type ReservationFile struct {
	ID            uuid.UUID
	GuestName     string
	GuestEmail    string
	Movie         MovieFile
	ReservedSeats []SeatFile
	StartTime     time.Time
	Room          string
}

type SeatFile struct {
	Row    string
	Number int
}
type MovieFile struct {
	Title       string
	Description string
}

func generateReservationPDF(res *ReservationFile) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(0, 10, fmt.Sprintf("Reservation ID: %s", res.ID.String()))
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, fmt.Sprintf("Guest Name: %s", res.GuestName))
	pdf.Ln(8)

	pdf.Cell(0, 10, fmt.Sprintf("Guest Email: %s", res.GuestEmail))
	pdf.Ln(8)

	pdf.Cell(0, 10, fmt.Sprintf("Movie Title: %s", res.Movie.Title))
	pdf.Ln(8)

	pdf.MultiCell(0, 8, fmt.Sprintf("Movie Description: %s", res.Movie.Description), "", "", false)
	pdf.Ln(4)

	pdf.Cell(0, 10, fmt.Sprintf("Start Time: %s", res.StartTime.Format("02 Jan 2006 15:04")))
	pdf.Ln(8)

	pdf.Cell(0, 10, fmt.Sprintf("Room: %s", res.Room))
	pdf.Ln(12)

	pdf.Cell(0, 10, "Reserved Seats:")
	pdf.Ln(8)
	for _, seat := range res.ReservedSeats {
		pdf.Cell(0, 10, fmt.Sprintf("- Row: %s, Number: %d", seat.Row, seat.Number))
		pdf.Ln(6)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
