package cart

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID         uuid.UUID      `gorm:"type:varchar(100)"`
	CustomerID uuid.UUID      `gorm:"type:varchar(100)"`
	ProductID  uuid.UUID      `gorm:"type:varchar(100)"`
	Quantity   int64          `gorm:"type:decimal"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
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
	}
}
