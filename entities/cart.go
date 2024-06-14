package entities

import (
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	ID         uuid.UUID
	CustomerID uuid.UUID `gorm:"type:varchar(100)"`
	ProductID  uuid.UUID `gorm:"type:varchar(100)"`
	Product    *Product
	Quantity   int64 `gorm:"type:decimal"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

type CartRepositoryInterface interface {
	Create(cart *Cart) error
	GetById(cart *Cart) error
	Update(cart *Cart) error
	Delete(cart *Cart) error
	GetAll(carts *[]Cart, userData *middlewares.Claims) error
}

type CartUseCaseInterface interface {
	Create(cart *Cart) (Cart, error)
	GetById(id uuid.UUID) (Cart, error)
	Update(cart *Cart) (Cart, error)
	Delete(id uuid.UUID) (Cart, error)
	GetAll(carts *[]Cart, userData *middlewares.Claims) ([]Cart, error)
}
