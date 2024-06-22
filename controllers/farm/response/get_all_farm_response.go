package response

import "github.com/blueharvest-alterra/go-back-end/entities"

type FarmGetAll struct {
	Farms []FarmResponse `json:"farms"`
}

func SliceFromUseCase(farms *[]entities.Farm) *FarmGetAll {
	allFarms := make([]FarmResponse, len(*farms))
	for i, _farm := range *farms {
		allFarms[i] = FarmResponse{
			ID:                      _farm.ID,
			Title:                   _farm.Description,
			Description:             _farm.Description,
			Picture:                 _farm.Picture,
			MinimumInvestmentAmount: _farm.MinimumInvestmentAmount,
			CountInvestment:         _farm.CountInvestment,
		}
	}

	return &FarmGetAll{
		Farms: allFarms,
	}	
}
