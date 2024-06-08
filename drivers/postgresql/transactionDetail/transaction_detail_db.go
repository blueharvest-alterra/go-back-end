package transactionDetail

import (
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/product"
	"github.com/google/uuid"
)

type TransactionDetail struct {
	ID            uuid.UUID `gorm:"type:varchar(100)"`
	TransactionID uuid.UUID `gorm:"type:varchar(100)"`
	ProductID     uuid.UUID `gorm:"type:varchar(100)"`
	Product       product.Product
	Quantity      uint
	TotalPrice    float64 `gorm:"type:decimal"`
}
