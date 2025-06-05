package handler

import (
	"absolutecinema/internal/openapi/gen/screeninggen"
	screening_service "absolutecinema/internal/service/screening"
	"encoding/json"
	"net/http"
	"time"

	"github.com/oapi-codegen/runtime/types"
)

func toScreeningOutput(s screening_service.ScreeningOutput) screeninggen.ScreeningOutput {
	return screeninggen.ScreeningOutput{
		Id:        &s.ID,
		StartTime: &s.StartTime,
		Movie: &screeninggen.MovieOutput{
			Id:    &s.Movie.ID,
			Title: &s.Movie.Title,
		},
		Room: &screeninggen.RoomOutput{
			Id:   &s.Room.ID,
			Name: &s.Room.Name,
		},
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func optionalDateToTime(d *types.Date) *time.Time {
	if d == nil {
		return nil
	}
	return &d.Time
}
