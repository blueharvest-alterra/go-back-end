package request

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type EditCartRequest struct {
	ID         uuid.UUID
	CustomerID uuid.UUID `json:"customer_id"`
	ProductID  uuid.UUID `json:"product_id"`
	Quantity   int64     `json:"quantity"`
}

func (r *EditCartRequest) ToEntities() *entities.Cart {
	return &entities.Cart{
		ID:         r.ID,
		CustomerID: r.CustomerID,
		ProductID:  r.ProductID,
		Quantity:   r.Quantity,
	}
}
