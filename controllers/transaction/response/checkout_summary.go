package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type AddressSummaryResponse struct {
	ID        uuid.UUID `json:"id"`
	Address   string    `json:"address"`
	CityID    string    `json:"city_id"`
	City      string    `json:"city"`
	StateID   string    `json:"state_id"`
	State     string    `json:"state"`
	ZipCode   string    `json:"zip_code"`
	Country   string    `json:"country"`
	Longitude string    `json:"longitude"`
	Latitude  string    `json:"latitude"`
}

type CourierSummaryResponse struct {
	DestinationAddress AddressSummaryResponse `json:"destination_address"`
	Name               string                 `json:"name"`
	Fee                float64                `json:"fee"`
	Type               string                 `json:"type"`
}

type TransactionDetailsSummaryResponse struct {
	Product    ProductResponse `json:"product"`
	Quantity   uint            `json:"quantity"`
	TotalPrice float64         `json:"total_price"`
}

type PromoSummaryResponse struct {
	ID     uuid.UUID            `json:"id"`
	Name   string               `json:"name"`
	Code   string               `json:"code"`
	Status entities.PromoStatus `json:"status"`
	Amount float64              `json:"amount"`
}

type CheckoutSummaryResponse struct {
	SubTotal           float64                             `json:"sub_total"`
	Tax                float64                             `json:"tax"`
	Discount           float64                             `json:"discount"`
	Total              float64                             `json:"total"`
	Promo              PromoSummaryResponse                `json:"promo"`
	Courier            CourierSummaryResponse              `json:"courier"`
	TransactionDetails []TransactionDetailsSummaryResponse `json:"transaction_details"`
}

func GetCheckoutSummaryFromUseCase(transaction *entities.Transaction) *CheckoutSummaryResponse {
	allTransactionDetail := make([]TransactionDetailsSummaryResponse, len(transaction.TransactionDetails))
	for i, transactionDetail := range transaction.TransactionDetails {
		allTransactionDetail[i] = TransactionDetailsSummaryResponse{
			Product: ProductResponse{
				ID:          transactionDetail.Product.ID,
				Name:        transactionDetail.Product.Name,
				Description: transactionDetail.Product.Description,
				Price:       transactionDetail.Product.Price,
				Thumbnail:   transactionDetail.Product.Thumbnail,
			},
			Quantity:   transactionDetail.Quantity,
			TotalPrice: transactionDetail.TotalPrice,
		}
	}

	return &CheckoutSummaryResponse{
		SubTotal: transaction.SubTotal,
		Tax:      transaction.Tax,
		Discount: transaction.Discount,
		Total:    transaction.Total,
		Promo: PromoSummaryResponse{
			ID:     transaction.Promo.ID,
			Name:   transaction.Promo.Name,
			Code:   transaction.Promo.Code,
			Status: transaction.Promo.Status,
			Amount: transaction.Promo.Amount,
		},
		Courier: CourierSummaryResponse{
			DestinationAddress: AddressSummaryResponse{
				ID:        transaction.Courier.DestinationAddress.ID,
				Address:   transaction.Courier.DestinationAddress.Address,
				CityID:    transaction.Courier.DestinationAddress.CityID,
				City:      transaction.Courier.DestinationAddress.City,
				StateID:   transaction.Courier.DestinationAddress.StateID,
				State:     transaction.Courier.DestinationAddress.State,
				ZipCode:   transaction.Courier.DestinationAddress.ZipCode,
				Country:   transaction.Courier.DestinationAddress.Country,
				Longitude: transaction.Courier.DestinationAddress.Longitude,
				Latitude:  transaction.Courier.DestinationAddress.Latitude,
			},
			Name: transaction.Courier.Name,
			Fee:  transaction.Courier.Fee,
			Type: transaction.Courier.Type,
		},
		TransactionDetails: allTransactionDetail,
	}
}
