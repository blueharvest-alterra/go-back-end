package product

import (
	"fmt"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"gorm.io/gorm"
	"mime/multipart"
)

type Repo struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Create(product *entities.Product, thumbnail []*multipart.FileHeader) error {
	productDb := FromUseCase(product)

	if err := productDb.UploadThumbnail(thumbnail); err != nil {
		fmt.Println("err repo Create", err)
		return err
	}

	if err := r.DB.Create(&productDb).Error; err != nil {
		return err
	}

	*product = *productDb.ToUseCase()
	return nil
}
