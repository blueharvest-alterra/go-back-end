package request

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type CreatePromoRequest struct {
	ID     uuid.UUID
	Name   string  `json:"name"`
	Code   string  `json:"code"`
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
}

func (r *CreatePromoRequest) ToEntities() *entities.Promo {
	return &entities.Promo{
		ID:     r.ID,
		Name:   r.Name,
		Code:   r.Code,
		Status: entities.PromoStatus(r.Status),
		Amount: r.Amount,
	}
}
