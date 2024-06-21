package entities

import (
	"github.com/google/uuid"
)

type Payment struct {
	ID         uuid.UUID
	ExternalID string
	InvoiceURL string
	Status     string
	Context    string
	ContextID  string
	Amount     float64
}

type PaymentRepositoryInterface interface {
	UpdateStatus(payment *Payment) error
	UpdateBuyProductContext(payment *Payment) error
	UpdateFarmInvestContext(payment *Payment) error
}

type PaymentUseCaseInterface interface {
	UpdateStatus(payment *Payment) (Payment, error)
	UpdatePaymentContext(payment *Payment) (Payment, error)
}
