package entities

import (
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mime/multipart"
	"time"
)

type ProductStatus string

const (
	ProductAvailable   ProductStatus = "available"
	ProductUnavailable ProductStatus = "unavailable"
)

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Price       float64
	Status      ProductStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	Thumbnail   string
}

type ProductRepositoryInterface interface {
	Create(product *Product, thumbnail []*multipart.FileHeader) error
	Update(product *Product, thumbnail []*multipart.FileHeader) error
	Delete(product *Product) error
	GetByID(product *Product) error
	GetAll(product *[]Product) error
}

type ProductUseCaseInterface interface {
	Create(product *Product, userData *middlewares.Claims, thumbnail []*multipart.FileHeader) (Product, error)
	Update(product *Product, userData *middlewares.Claims, thumbnail []*multipart.FileHeader) (Product, error)
	Delete(product *Product, userData *middlewares.Claims) (Product, error)
	GetByID(product *Product) (Product, error)
	GetAll(product *[]Product) ([]Product, error)
}
