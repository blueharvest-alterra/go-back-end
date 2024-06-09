package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type FarmInvestResponse struct {
	ID               uuid.UUID `json:"id"`
	CustomerID       uuid.UUID `json:"customer_id"`
	FarmID           uuid.UUID `json:"farm_id"`
	InvestmentAmount float64   `json:"investment_amount"`
}

func FarmInvestResponseFromUseCase(farmInvest *entities.FarmInvest) *FarmInvestResponse {
	return &FarmInvestResponse{
		ID:               farmInvest.ID,
		CustomerID:       farmInvest.CustomerID,
		FarmID:           farmInvest.FarmID,
		InvestmentAmount: farmInvest.InvestmentAmount,
	}
}
