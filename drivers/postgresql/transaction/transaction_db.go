package transaction

import (
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/courier"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/customer"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/transactionDetail"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID                 uuid.UUID `gorm:"type:varchar(100)"`
	Type               string    `gorm:"type:varchar(100)"`
	Status             string    `gorm:"type:varchar(50)"`
	CustomerID         uuid.UUID `gorm:"type:varchar(100)"`
	Customer           customer.Customer
	SubTotal           float64 `gorm:"type:decimal"`
	Tax                float64 `gorm:"type:decimal"`
	Discount           float64 `gorm:"type:decimal"`
	Total              float64 `gorm:"type:decimal"`
	Quantity           uint
	PaymentExternalID  string    `gorm:"type:text"`
	CourierID          uuid.UUID `gorm:"type:varchar(100)"`
	Courier            courier.Courier
	TransactionDetails transactionDetail.TransactionDetail
	CreatedAt          time.Time      `gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
