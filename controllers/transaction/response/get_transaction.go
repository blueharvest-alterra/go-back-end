package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"time"
)

type CustomerResponse struct {
	ID          uuid.UUID `json:"id"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
}

type CourierResponse struct {
	ID                   uuid.UUID `json:"id"`
	DestinationAddressID uuid.UUID `json:"destination_address_id"`
	Name                 string    `json:"name"`
	Fee                  float64   `json:"fee"`
	Type                 string    `json:"type"`
}

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Thumbnail   string    `json:"thumbnail"`
}

type TransactionDetailsResponse struct {
	ID         uuid.UUID       `json:"id"`
	Product    ProductResponse `json:"product"`
	Quantity   uint            `json:"quantity"`
	TotalPrice float64         `json:"total_price"`
}

type PaymentResponse struct {
	ID         uuid.UUID `json:"id"`
	ExternalID string    `json:"external_id"`
	InvoiceURL string    `json:"invoice_url"`
	Status     string    `json:"status"`
	Amount     float64   `json:"amount"`
}

type GetTransactionResponse struct {
	ID                 uuid.UUID                    `json:"id"`
	Customer           CustomerResponse             `json:"customer"`
	SubTotal           float64                      `json:"sub_total"`
	Tax                float64                      `json:"tax"`
	Discount           float64                      `json:"discount"`
	Total              float64                      `json:"total"`
	Payment            PaymentResponse              `json:"payment"`
	Courier            CourierResponse              `json:"courier"`
	TransactionDetails []TransactionDetailsResponse `json:"transaction_details"`
	CreatedAt          time.Time                    `json:"created_at"`
}

func GetTransactionFromUseCase(transaction *entities.Transaction) *GetTransactionResponse {
	allTransactionDetail := make([]TransactionDetailsResponse, len(transaction.TransactionDetails))
	for i, transactionDetail := range transaction.TransactionDetails {
		allTransactionDetail[i] = TransactionDetailsResponse{
			ID: transactionDetail.ID,
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

	return &GetTransactionResponse{
		ID: transaction.ID,
		Customer: CustomerResponse{
			ID:          transaction.Customer.ID,
			FullName:    transaction.Customer.FullName,
			PhoneNumber: transaction.Customer.PhoneNumber,
		},
		SubTotal: transaction.SubTotal,
		Tax:      transaction.Tax,
		Discount: transaction.Discount,
		Total:    transaction.Total,
		Payment: PaymentResponse{
			ID:         transaction.Payment.ID,
			ExternalID: transaction.Payment.ExternalID,
			InvoiceURL: transaction.Payment.InvoiceURL,
			Status:     transaction.Payment.Status,
			Amount:     transaction.Payment.Amount,
		},
		Courier: CourierResponse{
			ID:                   transaction.Courier.ID,
			DestinationAddressID: transaction.Courier.DestinationAddressID,
			Name:                 transaction.Courier.Name,
			Fee:                  transaction.Courier.Fee,
			Type:                 transaction.Courier.Type,
		},
		TransactionDetails: allTransactionDetail,
		CreatedAt:          transaction.CreatedAt,
	}
}
