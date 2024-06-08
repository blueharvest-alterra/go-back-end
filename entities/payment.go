package entities

import "github.com/google/uuid"

type Payment struct {
	ID         uuid.UUID
	ExternalID string
	InvoiceURL string
	Status     string
	Amount     float64
}
