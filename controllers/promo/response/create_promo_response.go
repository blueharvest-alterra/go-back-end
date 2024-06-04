package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type PromoResponse struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Code   string    `json:"code"`
	Status string    `json:"status"`
	Amount float64   `json:"amount"`
}

func FarmResponseFromUseCase(promo *entities.Promo) *PromoResponse {
	return &PromoResponse{
		ID:     promo.ID,
		Name:   promo.Name,
		Code:   promo.Code,
		Status: string(promo.Status),
		Amount: promo.Amount,
	}
}
