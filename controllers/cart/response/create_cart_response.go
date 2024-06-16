package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ProductResponse struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Thumbnail   string         `json:"thumbnail"`
}

type CartResponse struct {
	ID         uuid.UUID        `json:"id"`
	CustomerID uuid.UUID        `json:"customer_id"`
	ProductID  uuid.UUID        `json:"product_id"`
	Quantity   int64            `json:"quantity"`
	Product    *ProductResponse `json:"product"`
}

func CartResponseFromUseCase(cart *entities.Cart) *CartResponse {
	var productResponse *ProductResponse
	if cart.Product != nil {
		productResponse = &ProductResponse{
			ID:          cart.Product.ID,
			Name:        cart.Product.Name,
			Description: cart.Product.Description,
			Price:       cart.Product.Price,
			Status:      string(cart.Product.Status),
			CreatedAt:   cart.Product.CreatedAt,
			UpdatedAt:   cart.Product.UpdatedAt,
			DeletedAt:   cart.Product.DeletedAt,
			Thumbnail:   cart.Product.Thumbnail,
		}
	}

	return &CartResponse{
		ID:         cart.ID,
		CustomerID: cart.CustomerID,
		ProductID:  cart.ProductID,
		Quantity:   cart.Quantity,
		Product:    productResponse,
	}
}
