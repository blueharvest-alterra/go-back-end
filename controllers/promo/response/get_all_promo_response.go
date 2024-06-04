package response

import "github.com/blueharvest-alterra/go-back-end/entities"

type PromoGetAll struct {
	Promos []PromoResponse `json:"promos"`
}

func SliceFromUseCase(promos *[]entities.Promo) *PromoGetAll {
	allPromos := make([]PromoResponse, len(*promos))
	for i, _promo := range *promos {
		allPromos[i] = PromoResponse{
			ID:     _promo.ID,
			Name:   _promo.Name,
			Code:   _promo.Code,
			Status: string(_promo.Status),
			Amount: _promo.Amount,
		}
	}

	return &PromoGetAll{
		Promos: allPromos,
	}
}
