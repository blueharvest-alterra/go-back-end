package response

import "github.com/blueharvest-alterra/go-back-end/entities"

type GetAllTransactionResponse struct {
	Transactions []GetTransactionResponse `json:"transactions"`
}

func SliceFromUseCase(products *[]entities.Transaction) *GetAllTransactionResponse {
	allTransactions := make([]GetTransactionResponse, len(*products))
	for i, product := range *products {
		allTransactions[i] = *GetTransactionFromUseCase(&product)
	}

	return &GetAllTransactionResponse{Transactions: allTransactions}
}
