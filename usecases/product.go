package usecases

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
	"mime/multipart"
)

type ProductUseCase struct {
	repository entities.ProductRepositoryInterface
}

func NewProductUseCase(repository entities.ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{
		repository: repository,
	}
}

func (cu *ProductUseCase) Create(product *entities.Product, userData *middlewares.Claims, thumbnail []*multipart.FileHeader) (entities.Product, error) {
	if userData.Role != "admin" {
		return entities.Product{}, constant.ErrNotAuthorized
	}

	if len(thumbnail) == 0 || product.Name == "" || product.Description == "" || product.Price < 1 {
		return entities.Product{}, constant.ErrEmptyInput
	}

	product.ID = uuid.New()
	product.Status = "available"

	if err := cu.repository.Create(product, thumbnail); err != nil {
		return entities.Product{}, err
	}

	return *product, nil
}

func (cu *ProductUseCase) GetByID(product *entities.Product) (entities.Product, error) {
	if product.ID == uuid.Nil {
		return entities.Product{}, constant.ErrEmptyInput
	}

	if err := cu.repository.GetByID(product); err != nil {
		return entities.Product{}, err
	}

	return *product, nil
}

func (cu *ProductUseCase) GetAll(product *[]entities.Product) ([]entities.Product, error) {
	if err := cu.repository.GetAll(product); err != nil {
		return []entities.Product{}, err
	}

	return *product, nil
}
