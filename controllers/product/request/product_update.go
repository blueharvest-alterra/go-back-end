package request

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"mime/multipart"
)

type ProductUpdateRequest struct {
	Name        string                   `form:"name"`
	Description string                   `form:"description"`
	Price       float64                  `form:"price"`
	Thumbnail   *[]*multipart.FileHeader `form:"thumbnail"`
	Status      string                   `form:"status"`
}

func (p *ProductUpdateRequest) ToEntities() *entities.Product {
	return &entities.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Status:      p.Status,
	}
}
