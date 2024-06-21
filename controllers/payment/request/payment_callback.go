package request

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"strings"
)

type PaymentCallbackRequest struct {
	ID         string `json:"id"`
	ExternalID string `json:"external_id"`
	Status     string `json:"status"`
}

func (p *PaymentCallbackRequest) ToEntities() *entities.Payment {
	split := strings.Split(p.ExternalID, ":")
	context := split[0]
	contextID := split[1]
	paymentId, _ := uuid.Parse(split[2])

	return &entities.Payment{
		ID:         paymentId,
		ExternalID: p.ID,
		Status:     p.Status,
		Context:    context,
		ContextID:  contextID,
	}
}
