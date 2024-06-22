package response

import "github.com/blueharvest-alterra/go-back-end/entities"

type FarmInvestGetAll struct {
	FarmInvests []FarmInvestResponse `json:"farm-invest"`
}

func SliceFromUseCase(farmInvests *[]entities.FarmInvest) *FarmInvestGetAll {
	allFarmInvests := make([]FarmInvestResponse, len(*farmInvests))
	for i, _farmInvest := range *farmInvests {
		allFarmInvests[i] = FarmInvestResponse{
			ID:               _farmInvest.ID,
			CustomerID:       _farmInvest.CustomerID,
			FarmID:           _farmInvest.FarmID,
			InvestmentAmount: _farmInvest.InvestmentAmount,
			Payment: PaymentResponse{
				ID:         _farmInvest.Payment.ID,
				ExternalID: _farmInvest.Payment.ExternalID,
				InvoiceURL: _farmInvest.Payment.InvoiceURL,
				Status:     _farmInvest.Payment.Status,
				Amount:     _farmInvest.Payment.Amount,
			},
		}
	}

	return &FarmInvestGetAll{
		FarmInvests: allFarmInvests,
	}
}
