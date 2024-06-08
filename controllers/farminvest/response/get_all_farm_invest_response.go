package response

import "github.com/blueharvest-alterra/go-back-end/entities"

type FarmInvestGetAll struct {
	FarmInvests []FarmInvestResponse `json:"farmInvest"`
}

func SliceFromUseCase(farmInvests *[]entities.FarmInvest) *FarmInvestGetAll {
	allFarmInvests := make([]FarmInvestResponse, len(*farmInvests))
	for i, _farmInvest := range *farmInvests {
		allFarmInvests[i] = FarmInvestResponse{
			ID:               _farmInvest.ID,
			CustomerID:       _farmInvest.CustomerID,
			FarmID:           _farmInvest.FarmID,
			InvestmentAmount: _farmInvest.InvestmentAmount,
		}
	}

	return &FarmInvestGetAll{
		FarmInvests: allFarmInvests,
	}
}
