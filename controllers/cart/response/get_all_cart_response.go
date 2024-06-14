package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
)

type CartGetAll struct {
	Carts []CartResponse `json:"carts"`
}


func SliceFromUseCase(carts *[]entities.Cart) *CartGetAll {
    allcarts := make([]CartResponse, len(*carts))
    for i, _cart := range *carts {
        var productResponse ProductResponse
        if _cart.Product != nil {
            productResponse = ProductResponse{
                ID:          _cart.Product.ID,
                Name:        _cart.Product.Name,
                Description: _cart.Product.Description,
                Price:       _cart.Product.Price,
                Status:      _cart.Product.Status,
                Thumbnail:   _cart.Product.Thumbnail,
                CreatedAt:   _cart.Product.CreatedAt,
                UpdatedAt:   _cart.Product.UpdatedAt,
            }
        }

        allcarts[i] = CartResponse{
            ID:         _cart.ID,
            CustomerID: _cart.CustomerID,
            ProductID:  _cart.ProductID,
            Quantity:   _cart.Quantity,
            Product:    productResponse,
        }
    }

    return &CartGetAll{
        Carts: allcarts,
    }
}
