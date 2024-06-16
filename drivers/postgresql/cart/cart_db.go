package cart

import (
	"time"

	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/product"
	"github.com/blueharvest-alterra/go-back-end/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID         uuid.UUID      `gorm:"type:varchar(100)"`
	CustomerID uuid.UUID      `gorm:"type:varchar(100)"`
	Quantity   int64          `gorm:"type:decimal"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	ProductID  uuid.UUID      `gorm:"type:varchar(100)"`
	Product    product.Product
}

func FromUseCase(cart *entities.Cart) *Cart {
	return &Cart{
		ID:         cart.ID,
		CustomerID: cart.CustomerID,
		ProductID:  cart.ProductID,
		Quantity:   cart.Quantity,
		CreatedAt:  cart.CreatedAt,
		UpdatedAt:  cart.UpdatedAt,
		DeletedAt:  cart.DeletedAt,
		Product: product.Product{
			ID:          cart.ProductID,
			Name:        cart.Product.Name,
			Description: cart.Product.Description,
			Price:       cart.Product.Price,
			Status:      product.Status(cart.Product.Status),
			Thumbnail:   cart.Product.Thumbnail,
		},
	}
}

func (u *Cart) ToUseCase() *entities.Cart {
	return &entities.Cart{
		ID:         u.ID,
		CustomerID: u.CustomerID,
		ProductID:  u.ProductID,
		Quantity:   u.Quantity,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
		DeletedAt:  u.DeletedAt,
		Product: entities.Product{
			ID:          u.Product.ID,
			Name:        u.Product.Name,
			Description: u.Product.Description,
			Price:       u.Product.Price,
			Status:      entities.ProductStatus(u.Product.Status),
			Thumbnail:   u.Product.Thumbnail,
		},
	}
}
