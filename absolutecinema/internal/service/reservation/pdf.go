package reservation_service

import (
	"absolutecinema/internal/database/repository/reservation"
	"bytes"
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func generateReservationPDF(res *reservation.Reservation) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(40, 10, fmt.Sprintf("Reservation ID: %s", res.ID.String()))
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Screening ID: %s", res.ScreeningID.String()))
	pdf.Ln(8)

	if res.UserID != nil {
		pdf.Cell(40, 10, fmt.Sprintf("User ID: %s", res.UserID.String()))
		pdf.Ln(8)
	} else {
		pdf.Cell(40, 10, fmt.Sprintf("Guest Name: %s", res.GuestName))
		pdf.Ln(8)
		pdf.Cell(40, 10, fmt.Sprintf("Guest Email: %s", res.GuestEmail))
		pdf.Ln(8)
	}

	pdf.Cell(40, 10, "Reserved Seats:")
	pdf.Ln(8)
	for _, seat := range res.ReservedSeats {
		pdf.Cell(40, 10, fmt.Sprintf("- Seat ID: %s", seat.SeatID.String()))
		pdf.Ln(6)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
