package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type FarmMonitorResponse struct {
	ID              uuid.UUID `json:"id"`
	FarmID          uuid.UUID `json:"farm_id"`
	Temperature     float64   `json:"temperature"`
	PH              float64   `json:"ph"`
	DissolvedOxygen float64   `json:"dissolved_oxygen"`
}

func FarmMonitorResponseFromUseCase(farmMonitor *entities.FarmMonitor) *FarmMonitorResponse {
	return &FarmMonitorResponse{
		ID:              farmMonitor.ID,
		FarmID:          farmMonitor.FarmID,
		Temperature:     farmMonitor.Temperature,
		PH:              farmMonitor.PH,
		DissolvedOxygen: farmMonitor.DissolvedOxygen,
	}
}
