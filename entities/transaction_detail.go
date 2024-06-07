package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TransactionDetail struct {
	ID            uuid.UUID
	TransactionID uuid.UUID
	ProductID     uuid.UUID
	Product       Product
	Quantity      uint
	TotalPrice    float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
