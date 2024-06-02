package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type FarmResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

func FarmResponseFromUseCase(farm *entities.Farm) *FarmResponse {
	return &FarmResponse{
		ID:          farm.ID,
		Title:       farm.Title,
		Description: farm.Description,
	}
}
