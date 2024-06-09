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
		allcarts[i] = CartResponse{
			ID:         _cart.ID,
			CustomerID: _cart.CustomerID,
			ProductID:  _cart.ProductID,
			Quantity:   _cart.Quantity,
		}
	}

	return &CartGetAll{
		Carts: allcarts,
	}
}
