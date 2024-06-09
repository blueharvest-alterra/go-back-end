package request

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"mime/multipart"
)

type ProductCreateRequest struct {
	Name        string                  `form:"name"`
	Description string                  `form:"description"`
	Price       float64                 `form:"price"`
	Thumbnail   []*multipart.FileHeader `form:"thumbnail"`
}

func (r *ProductCreateRequest) ToEntities() *entities.Product {
	return &entities.Product{
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
	}
}
