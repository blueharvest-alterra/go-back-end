package request

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type CreateFarmMonitorRequest struct {
	ID              uuid.UUID
	FarmID          uuid.UUID `json:"farm_id"`
	Temperature     float64   `json:"temperature"`
	PH              float64   `json:"ph"`
	DissolvedOxygen float64   `json:"dissolved_oxygen"`
}

func (r *CreateFarmMonitorRequest) ToEntities() *entities.FarmMonitor {
	return &entities.FarmMonitor{
		ID:              r.ID,
		FarmID:          r.FarmID,
		Temperature:     r.Temperature,
		PH:              r.PH,
		DissolvedOxygen: r.DissolvedOxygen,
	}
}
