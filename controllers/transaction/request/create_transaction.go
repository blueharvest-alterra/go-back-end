package request

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type DetailTransactionResponse struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  uint      `json:"quantity"`
}

type CourierResponse struct {
	DestinationAddressID uuid.UUID `json:"destination_address_id"`
	Name                 string    `json:"name"`
	Type                 string    `json:"type"`
	Fee                  float64   `json:"fee"`
}

type TransactionCreateResponse struct {
	TransactionDetails []DetailTransactionResponse `json:"transaction_details"`
	PromoID            string                      `json:"promo_id" binding:"omitempty"`
	Courier            CourierResponse             `json:"courier"`
}

func (r *TransactionCreateResponse) ToEntities() *entities.Transaction {
	allTransactionDetail := make([]entities.TransactionDetail, len(r.TransactionDetails))
	for i, transaction := range r.TransactionDetails {
		allTransactionDetail[i] = entities.TransactionDetail{
			ID:       uuid.New(),
			Product:  entities.Product{ID: transaction.ProductID},
			Quantity: transaction.Quantity,
		}
	}

	var promoID uuid.UUID
	if r.PromoID != "" {
		promoIDUUID, _ := uuid.Parse(r.PromoID)
		promoID = promoIDUUID
	}

	return &entities.Transaction{
		TransactionDetails: allTransactionDetail,
		PromoID:            promoID,
		Courier: entities.Courier{
			DestinationAddressID: r.Courier.DestinationAddressID,
			Name:                 r.Courier.Name,
			Type:                 r.Courier.Type,
			Fee:                  r.Courier.Fee,
		},
	}
}
