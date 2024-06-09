package response

import "github.com/blueharvest-alterra/go-back-end/entities"

type ProductGetAll struct {
	Products []ProductDetail `json:"products"`
}

func SliceFromUseCase(products *[]entities.Product) *ProductGetAll {
	allProducts := make([]ProductDetail, len(*products))
	for i, product := range *products {
		allProducts[i] = *ProductDetailFromUseCase(&product)
	}

	return &ProductGetAll{allProducts}
}
