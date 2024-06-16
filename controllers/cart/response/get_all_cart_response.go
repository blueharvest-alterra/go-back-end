package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
)

type CartGetAll struct {
	Carts []CartResponse `json:"carts"`
}

func SliceFromUseCase(carts *[]entities.Cart) *CartGetAll {
	allCarts := make([]CartResponse, len(*carts))
	for i, cart := range *carts {
		allCarts[i] = *CartResponseFromUseCase(&cart)
	}

	return &CartGetAll{
		Carts: allCarts,
	}
}
