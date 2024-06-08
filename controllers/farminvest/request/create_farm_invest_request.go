package request

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type CreateFarmInvestRequest struct {
	ID               uuid.UUID
	CustomerID       uuid.UUID `json:"customer_id"`
	FarmID           uuid.UUID `json:"farm_id"`
	InvestmentAmount float64   `json:"investment_amount"`
}

func (r *CreateFarmInvestRequest) ToEntities() *entities.FarmInvest {
	return &entities.FarmInvest{
		ID:               r.ID,
		CustomerID:       r.CustomerID,
		FarmID:           r.FarmID,
		InvestmentAmount: r.InvestmentAmount,
	}
}
