package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type FarmResponse struct {
	ID                      uuid.UUID `json:"id"`
	Title                   string    `json:"title"`
	Description             string    `json:"description"`
	Picture                 string    `json:"picture"`
	MinimumInvestmentAmount float64   `json:"minimum_investment_amount"`
	CountInvestment         float64    `json:"count_investment"`
}

func FarmResponseFromUseCase(farm *entities.Farm) *FarmResponse {
	return &FarmResponse{
		ID:                      farm.ID,
		Title:                   farm.Title,
		Description:             farm.Description,
		Picture:                 farm.Picture,
		MinimumInvestmentAmount: farm.MinimumInvestmentAmount,
		CountInvestment:         farm.CountInvestment,
	}
}
