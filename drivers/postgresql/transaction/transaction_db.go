package transaction

import (
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/courier"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/customer"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/payment"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/product"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/transactionDetail"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

const TaxFee = float64(10000)

type Transaction struct {
	ID                 uuid.UUID `gorm:"type:varchar(100)"`
	CustomerID         uuid.UUID `gorm:"type:varchar(100)"`
	Customer           customer.Customer
	SubTotal           float64   `gorm:"type:decimal"`
	Tax                float64   `gorm:"type:decimal"`
	Discount           float64   `gorm:"type:decimal"`
	Total              float64   `gorm:"type:decimal"`
	PaymentID          uuid.UUID `gorm:"type:varchar(100)"`
	Payment            payment.Payment
	CourierID          uuid.UUID `gorm:"type:varchar(100)"`
	PromoID            uuid.UUID `gorm:"type:varchar(100)"`
	Courier            courier.Courier
	TransactionDetails []transactionDetail.TransactionDetail
	CreatedAt          time.Time      `gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

func FromUseCase(transaction *entities.Transaction) *Transaction {
	allTransactionDetails := make([]transactionDetail.TransactionDetail, len(transaction.TransactionDetails))
	for i, _transactionDetail := range transaction.TransactionDetails {
		allTransactionDetails[i] = transactionDetail.TransactionDetail{
			ID:            _transactionDetail.ID,
			TransactionID: _transactionDetail.TransactionID,
			Product: product.Product{
				ID:          _transactionDetail.Product.ID,
				Name:        _transactionDetail.Product.Name,
				Description: _transactionDetail.Product.Description,
				Price:       _transactionDetail.Product.Price,
				Status:      _transactionDetail.Product.Status,
				Thumbnail:   _transactionDetail.Product.Thumbnail,
			},
			Quantity:   _transactionDetail.Quantity,
			TotalPrice: _transactionDetail.TotalPrice,
		}
	}

	return &Transaction{
		ID:         transaction.ID,
		CustomerID: transaction.CustomerID,
		Customer: customer.Customer{
			ID:          transaction.Customer.ID,
			FullName:    transaction.Customer.FullName,
			PhoneNumber: transaction.Customer.PhoneNumber,
			BirthDate:   transaction.Customer.BirthDate,
		},
		SubTotal:  transaction.SubTotal,
		Tax:       transaction.Tax,
		Discount:  transaction.Discount,
		Total:     transaction.Total,
		PaymentID: transaction.PaymentID,
		Payment: payment.Payment{
			ID:         transaction.Payment.ID,
			ExternalID: transaction.Payment.ExternalID,
			InvoiceURL: transaction.Payment.InvoiceURL,
			Status:     transaction.Payment.Status,
		},
		CourierID: transaction.CourierID,
		PromoID:   transaction.PromoID,
		Courier: courier.Courier{
			ID:                   transaction.Courier.ID,
			DestinationAddressID: transaction.Courier.DestinationAddressID,
			Name:                 transaction.Courier.Name,
			Fee:                  transaction.Courier.Fee,
			Type:                 transaction.Courier.Type,
		},
		TransactionDetails: allTransactionDetails,
	}
}

func (u *Transaction) ToUseCase() *entities.Transaction {
	allTransactionDetails := make([]entities.TransactionDetail, len(u.TransactionDetails))
	for i, _transactionDetail := range u.TransactionDetails {
		allTransactionDetails[i] = entities.TransactionDetail{
			ID:            _transactionDetail.ID,
			TransactionID: _transactionDetail.TransactionID,
			ProductID:     _transactionDetail.ProductID,
			Product: entities.Product{
				ID:          _transactionDetail.Product.ID,
				Name:        _transactionDetail.Product.Name,
				Description: _transactionDetail.Product.Description,
				Price:       _transactionDetail.Product.Price,
				Status:      _transactionDetail.Product.Status,
				Thumbnail:   _transactionDetail.Product.Thumbnail,
			},
			Quantity:   _transactionDetail.Quantity,
			TotalPrice: _transactionDetail.TotalPrice,
		}
	}

	return &entities.Transaction{
		ID:         u.ID,
		CustomerID: u.CustomerID,
		Customer: entities.Customer{
			ID:          u.Customer.ID,
			FullName:    u.Customer.FullName,
			PhoneNumber: u.Customer.PhoneNumber,
			BirthDate:   u.Customer.BirthDate,
		},
		SubTotal: u.SubTotal,
		Tax:      u.Tax,
		Discount: u.Discount,
		Total:    u.Total,
		Payment: entities.Payment{
			ID:         u.Payment.ID,
			ExternalID: u.Payment.ExternalID,
			InvoiceURL: u.Payment.InvoiceURL,
			Status:     u.Payment.Status,
		},
		CourierID: u.CourierID,
		PromoID:   u.PromoID,
		Courier: entities.Courier{
			ID:                   u.Courier.ID,
			DestinationAddressID: u.Courier.DestinationAddressID,
			Name:                 u.Courier.Name,
			Fee:                  u.Courier.Fee,
			Type:                 u.Courier.Type,
		},
		TransactionDetails: allTransactionDetails,
	}
}
