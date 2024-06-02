package request

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type EditFarmRequest struct {
	ID          uuid.UUID
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (r *EditFarmRequest) ToEntities() *entities.Farm {
	return &entities.Farm{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
	}
}
