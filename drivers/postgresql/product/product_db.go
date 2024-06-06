package product

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID                uuid.UUID `gorm:"type:varchar(100);"`
	Name              string    `gorm:"type:varchar(255);not null"`
	Description       string    `gorm:"type:text;not null"`
	Price             float64   `gorm:"type:decimal;not null"`
	Status            string    `gorm:"type:varchar(50);not null"`
	PickupAvailable   bool      `gorm:"type:boolean;not null"`
	DeliveryAvailable bool      `gorm:"type:boolean;not null"`
	Thumbnail         string    `gorm:"type:text;not null"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         gorm.DeletedAt
}

func FromUseCase(product *entities.Product) *Product {
	return &Product{
		ID:                product.ID,
		Name:              product.Name,
		Description:       product.Description,
		Price:             product.Price,
		Status:            product.Status,
		PickupAvailable:   product.PickupAvailable,
		DeliveryAvailable: product.DeliveryAvailable,
		CreatedAt:         product.CreatedAt,
		UpdatedAt:         product.UpdatedAt,
		DeletedAt:         product.DeletedAt,
		Thumbnail:         product.Thumbnail,
	}
}

func (product *Product) ToUseCase() *entities.Product {
	return &entities.Product{
		ID:                product.ID,
		Name:              product.Name,
		Description:       product.Description,
		Price:             product.Price,
		Status:            product.Status,
		PickupAvailable:   product.PickupAvailable,
		DeliveryAvailable: product.DeliveryAvailable,
		CreatedAt:         product.CreatedAt,
		UpdatedAt:         product.UpdatedAt,
		DeletedAt:         product.DeletedAt,
		Thumbnail:         product.Thumbnail,
	}
}

//func (p *Product) UploadThumbnail(thumbnail []*multipart.FileHeader) error {
//	file, err := thumbnail[0].Open()
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//
//	ext := filepath.Ext(thumbnail[0].Filename)
//
//	ctx := context.Background()
//
//	objectName := uuid.NewString() + ext
//	url, err := google.Upload.UploadFile(ctx, file, objectName)
//	if err != nil {
//		return err
//	}
//
//	p.Thumbnail = url
//
//	return nil
//}
