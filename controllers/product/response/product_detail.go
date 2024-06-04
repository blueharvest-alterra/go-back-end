package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ThumbnailDetail struct {
	ID   uuid.UUID `json:"id"`
	Type string    `json:"type"`
	Key  string    `json:"key"`
}

type ProductDetail struct {
	ID                uuid.UUID      `json:"id"`
	Name              string         `json:"name"`
	Description       string         `json:"description"`
	Price             float64        `json:"price"`
	PickupAvailable   bool           `json:"pickup_available"`
	DeliveryAvailable bool           `json:"delivery_available"`
	Thumbnail         string         `json:"thumbnail"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at"`
}

func ProductDetailFromUseCase(product *entities.Product) *ProductDetail {
	return &ProductDetail{
		ID:                product.ID,
		Name:              product.Name,
		Description:       product.Description,
		Price:             product.Price,
		PickupAvailable:   product.PickupAvailable,
		DeliveryAvailable: product.DeliveryAvailable,
		Thumbnail:         product.Thumbnail,
		CreatedAt:         product.CreatedAt,
		UpdatedAt:         product.UpdatedAt,
		DeletedAt:         product.DeletedAt,
	}
}
