package handler

import (
	"absolutecinema/internal/openapi/gen/moviegen"
	movie_service "absolutecinema/internal/service/movie"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func ToMovieOutput(m movie_service.MovieOutput) moviegen.MovieOutput {
	actors := make([]struct {
		Id   *openapi_types.UUID `json:"id,omitempty"`
		Name *string             `json:"name,omitempty"`
	}, len(m.Actors))

	for i, a := range m.Actors {
		id := openapi_types.UUID(a.ID)
		name := a.Name
		actors[i] = struct {
			Id   *openapi_types.UUID `json:"id,omitempty"`
			Name *string             `json:"name,omitempty"`
		}{
			Id:   &id,
			Name: &name,
		}
	}

	return moviegen.MovieOutput{
		Id:          &m.ID,
		Title:       &m.Title,
		Director:    &m.Director,
		Description: &m.Description,
		PhotoPath:   &m.PhotoPath,
		Actors:      &actors,
	}
}
