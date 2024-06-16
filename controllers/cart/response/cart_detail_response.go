package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Status      string    `json:"status"`
	Thumbnail   string    `json:"thumbnail"`
}

type CartResponse struct {
	ID         uuid.UUID       `json:"id"`
	CustomerID uuid.UUID       `json:"customer_id"`
	ProductID  uuid.UUID       `json:"product_id"`
	Quantity   int64           `json:"quantity"`
	Product    ProductResponse `json:"product"`
}

func CartResponseFromUseCase(cart *entities.Cart) *CartResponse {
	return &CartResponse{
		ID:         cart.ID,
		CustomerID: cart.CustomerID,
		ProductID:  cart.ProductID,
		Quantity:   cart.Quantity,
		Product: ProductResponse{
			ID:          cart.Product.ID,
			Name:        cart.Product.Name,
			Description: cart.Product.Description,
			Price:       cart.Product.Price,
			Status:      string(cart.Product.Status),
			Thumbnail:   cart.Product.Thumbnail,
		},
	}
}
