package request

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type PaymentCallbackRequest struct {
	ID         string `json:"id"`
	ExternalID string `json:"external_id"`
	Status     string `json:"status"`
}

func (p *PaymentCallbackRequest) ToEntities() *entities.Payment {
	paymentId, _ := uuid.Parse(p.ExternalID)

	return &entities.Payment{
		ID:         paymentId,
		ExternalID: p.ID,
		Status:     p.Status,
	}
}
