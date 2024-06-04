package entities

import (
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mime/multipart"
	"time"
)

type Product struct {
	ID                uuid.UUID
	Name              string
	Description       string
	Price             float64
	Status            string
	PickupAvailable   bool
	DeliveryAvailable bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt
	Thumbnail         string
}

type ProductRepositoryInterface interface {
	Create(product *Product, thumbnail []*multipart.FileHeader) error
}

type ProductUseCaseInterface interface {
	Create(product *Product, userData *middlewares.Claims, thumbnail []*multipart.FileHeader) (Product, error)
}
