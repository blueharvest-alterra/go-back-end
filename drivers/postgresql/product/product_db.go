package product

import (
	"context"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/utils/google"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mime/multipart"
	"path/filepath"
	"time"
)

type Status string

const (
	Available   Status = "available"
	Unavailable Status = "unavailable"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:varchar(100);"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text;not null"`
	Price       float64   `gorm:"type:decimal;not null"`
	Status      Status    `gorm:"type:varchar(50);not null"`
	Thumbnail   string    `gorm:"type:text;not null"`
	CountSold   uint
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func FromUseCase(product *entities.Product) *Product {
	return &Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Status:      Status(product.Status),
		Thumbnail:   product.Thumbnail,
		CountSold:   product.CountSold,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		DeletedAt:   product.DeletedAt,
	}
}

func (product *Product) ToUseCase() *entities.Product {
	return &entities.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Status:      entities.ProductStatus(product.Status),
		Thumbnail:   product.Thumbnail,
		CountSold:   product.CountSold,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		DeletedAt:   product.DeletedAt,
	}
}

func (p *Product) UploadThumbnail(thumbnail []*multipart.FileHeader) error {
	file, err := thumbnail[0].Open()
	if err != nil {
		return err
	}
	defer file.Close()

	ext := filepath.Ext(thumbnail[0].Filename)

	ctx := context.Background()

	objectName := uuid.NewString() + ext
	url, err := google.Upload.UploadFile(ctx, file, objectName)
	if err != nil {
		return err
	}

	p.Thumbnail = url

	return nil
}
