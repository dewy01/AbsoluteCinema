package handler

import (
	"absolutecinema/internal/openapi/gen/roomgen"
	room_service "absolutecinema/internal/service/room"
)

func convertSeatInputs(seats []roomgen.SeatInput) []room_service.SeatInput {
	var out []room_service.SeatInput
	for _, s := range seats {
		out = append(out, room_service.SeatInput{
			Number: s.Number,
			Row:    s.Row,
		})
	}
	return out
}

func convertSeatOutputs(seats []room_service.SeatOutput) *[]roomgen.SeatOutput {
	var out []roomgen.SeatOutput
	for _, s := range seats {
		out = append(out, roomgen.SeatOutput{
			Id:     &s.ID,
			Number: &s.Number,
			Row:    &s.Row,
		})
	}
	return &out
}
