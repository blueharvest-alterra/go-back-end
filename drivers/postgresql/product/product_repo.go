package product

import (
	"errors"
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"gorm.io/gorm"
	"mime/multipart"
)

type Repo struct {
	DB *gorm.DB
}

func (r *Repo) Delete(product *entities.Product) error {
	productDb := FromUseCase(product)

	if err := r.DB.Where("id = ?", product.ID).Delete(&productDb).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	*product = *productDb.ToUseCase()
	return nil
}

func (r *Repo) Update(product *entities.Product, thumbnail []*multipart.FileHeader) error {
	productDb := FromUseCase(product)

	if len(thumbnail) != 0 {
		if err := productDb.UploadThumbnail(thumbnail); err != nil {
			return err
		}
	}

	if err := r.DB.Where("id = ?", product.ID).Updates(&productDb).Error; err != nil {
		return err
	}

	*product = *productDb.ToUseCase()
	return nil
}

func NewProductRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Create(product *entities.Product, thumbnail []*multipart.FileHeader) error {
	productDb := FromUseCase(product)

	if err := productDb.UploadThumbnail(thumbnail); err != nil {
		return err
	}

	if err := r.DB.Create(&productDb).Error; err != nil {
		return err
	}

	*product = *productDb.ToUseCase()
	return nil
}

func (r *Repo) GetByID(product *entities.Product) error {
	productDb := FromUseCase(product)

	if err := r.DB.First(&productDb).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	*product = *productDb.ToUseCase()
	return nil
}

func (r *Repo) GetAll(product *[]entities.Product) error {
	var productsDb []Product

	if err := r.DB.Find(&productsDb).Error; err != nil {
		return err
	}

	for _, _product := range productsDb {
		*product = append(*product, *_product.ToUseCase())
	}
	return nil
}
