package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type PaymentResponse struct {
	ID         uuid.UUID `json:"id"`
	ExternalID string    `json:"external_id"`
	InvoiceURL string    `json:"invoice_url"`
	Status     string    `json:"status"`
	Context    string
	ContextID  string
	Amount     float64 `json:"amount"`
}

type FarmInvestResponse struct {
	ID               uuid.UUID       `json:"id"`
	CustomerID       uuid.UUID       `json:"customer_id"`
	FarmID           uuid.UUID       `json:"farm_id"`
	InvestmentAmount float64         `json:"investment_amount"`
	Payment          PaymentResponse `json:"payment"`
}

func FarmInvestResponseFromUseCase(farmInvest *entities.FarmInvest) *FarmInvestResponse {
	return &FarmInvestResponse{
		ID:               farmInvest.ID,
		CustomerID:       farmInvest.CustomerID,
		FarmID:           farmInvest.FarmID,
		InvestmentAmount: farmInvest.InvestmentAmount,
		Payment: PaymentResponse{
			ID:         farmInvest.Payment.ID,
			ExternalID: farmInvest.Payment.ExternalID,
			InvoiceURL: farmInvest.Payment.InvoiceURL,
			Status:     farmInvest.Payment.Status,
			Amount:     farmInvest.Payment.Amount,
		},
	}
}
