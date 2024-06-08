package entities

import (
	"github.com/google/uuid"
)

type TransactionDetail struct {
	ID            uuid.UUID
	TransactionID uuid.UUID
	ProductID     uuid.UUID
	Product       Product
	Quantity      uint
	TotalPrice    float64
}
