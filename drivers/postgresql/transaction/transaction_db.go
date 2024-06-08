package transaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/courier"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/customer"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/product"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/transactionDetail"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

const TaxFee = float64(10000)

type Transaction struct {
	ID                 uuid.UUID `gorm:"type:varchar(100)"`
	Type               string    `gorm:"type:varchar(100)"`
	Status             string    `gorm:"type:varchar(50)"`
	CustomerID         uuid.UUID `gorm:"type:varchar(100)"`
	Customer           customer.Customer
	SubTotal           float64   `gorm:"type:decimal"`
	Tax                float64   `gorm:"type:decimal"`
	Discount           float64   `gorm:"type:decimal"`
	Total              float64   `gorm:"type:decimal"`
	PaymentExternalID  string    `gorm:"type:text"`
	PaymentInvoiceURL  string    `gorm:"type:text"`
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
		Type:       transaction.Type,
		Status:     transaction.Status,
		CustomerID: transaction.CustomerID,
		Customer: customer.Customer{
			ID:          transaction.Customer.ID,
			FullName:    transaction.Customer.FullName,
			PhoneNumber: transaction.Customer.PhoneNumber,
			BirthDate:   transaction.Customer.BirthDate,
		},
		SubTotal:          transaction.SubTotal,
		Tax:               transaction.Tax,
		Discount:          transaction.Discount,
		Total:             transaction.Total,
		PaymentExternalID: transaction.PaymentExternalID,
		PaymentInvoiceURL: transaction.PaymentInvoiceURL,
		CourierID:         transaction.CourierID,
		PromoID:           transaction.PromoID,
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
		Type:       u.Type,
		Status:     u.Status,
		CustomerID: u.CustomerID,
		Customer: entities.Customer{
			ID:          u.Customer.ID,
			FullName:    u.Customer.FullName,
			PhoneNumber: u.Customer.PhoneNumber,
			BirthDate:   u.Customer.BirthDate,
		},
		SubTotal:          u.SubTotal,
		Tax:               u.Tax,
		Discount:          u.Discount,
		Total:             u.Total,
		PaymentExternalID: u.PaymentExternalID,
		PaymentInvoiceURL: u.PaymentInvoiceURL,
		CourierID:         u.CourierID,
		PromoID:           u.PromoID,
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

type xdtInvoiceResponse struct {
	ID         string `json:"id"`
	ExternalID string `json:"external_id"`
	Amount     int64  `json:"amount"`
	InvoiceURL string `json:"invoice_url"`
}

type xdtInvoicePayload struct {
	ExternalID  string `json:"external_id"`
	Amount      int64  `json:"amount"`
	PayerEmail  string `json:"payer_email"`
	Description string `json:"description"`
}

func (p *Transaction) PaymentCreate() error {
	url := "https://api.xendit.co/v2/invoices"
	method := "POST"

	var payload xdtInvoicePayload
	payload.ExternalID = "buy_product:" + p.ID.String()
	payload.Amount = int64(p.Total)
	payload.PayerEmail = "blueharvest@gmail.com"
	payload.Description = "Blueharvest App Invoice"
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+os.Getenv("XDT_SECRET_API_KEY"))

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		fmt.Println("response Status:", res.Status)
		return constant.ErrPaymentGateway
	}

	var response xdtInvoiceResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}

	p.PaymentExternalID = response.ID
	p.PaymentInvoiceURL = response.InvoiceURL

	return nil
}
