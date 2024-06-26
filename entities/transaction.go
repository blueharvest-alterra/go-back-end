package entities

import (
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID                   uuid.UUID
	CustomerID           uuid.UUID
	Customer             Customer
	SubTotal             float64
	Tax                  float64
	Discount             float64
	Total                float64
	PaymentID            uuid.UUID
	Payment              Payment
	PromoID              uuid.UUID
	Promo                Promo
	DestinationAddressID uuid.UUID
	CourierID            uuid.UUID
	Courier              Courier
	TransactionDetails   []TransactionDetail
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt
}

type TransactionRepositoryInterface interface {
	Create(transaction *Transaction) error
	GetByID(transaction *Transaction, userData *middlewares.Claims) error
	GetAll(transactions *[]Transaction, userData *middlewares.Claims) error
	CheckoutSummary(transaction *Transaction, userData *middlewares.Claims) error
}

type TransactionUseCaseInterface interface {
	Create(transaction *Transaction, userData *middlewares.Claims) (Transaction, error)
	GetByID(transaction *Transaction, userData *middlewares.Claims) (Transaction, error)
	GetAll(transactions *[]Transaction, userData *middlewares.Claims) ([]Transaction, error)
	CheckoutSummary(transaction *Transaction, userData *middlewares.Claims) (Transaction, error)
}
