package farmInvest

import (
	"time"

	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/customer"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/farm"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/payment"

	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FarmInvest struct {
	ID               uuid.UUID         `gorm:"type:varchar(100)"`
	CustomerID       uuid.UUID         `gorm:"type:varchar(100)"`
	FarmID           uuid.UUID         `gorm:"type:varchar(100)"`
	PaymentID        uuid.UUID         `gorm:"type:varchar(100)"`
	InvestmentAmount float64           `gorm:"type:decimal"`
	CreatedAt        time.Time         `gorm:"autoCreateTime"`
	UpdatedAt        time.Time         `gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt    `gorm:"index"`
	Farm             farm.Farm         `gorm:"foreignKey:FarmID"`
	Payment          payment.Payment   `gorm:"foreignKey:PaymentID"`
	Customer         customer.Customer `gorm:"foreignKey:CustomerID"`
}

func FromUseCase(farminvest *entities.FarmInvest) *FarmInvest {
	return &FarmInvest{
		ID:               farminvest.ID,
		CustomerID:       farminvest.CustomerID,
		FarmID:           farminvest.FarmID,
		PaymentID:        farminvest.PaymentID,
		InvestmentAmount: farminvest.InvestmentAmount,
		CreatedAt:        farminvest.CreatedAt,
		UpdatedAt:        farminvest.UpdatedAt,
		DeletedAt:        farminvest.DeletedAt,
		Customer: customer.Customer{
			ID:          farminvest.Customer.ID,
			FullName:    farminvest.Customer.FullName,
			PhoneNumber: farminvest.Customer.PhoneNumber,
			BirthDate:   farminvest.Customer.BirthDate,
		},
		Farm: farm.Farm{
			ID:          farminvest.Farm.ID,
			Title:       farminvest.Farm.Title,
			Description: farminvest.Farm.Description,
			Picture:     farminvest.Farm.Picture,
			CreatedAt:   farminvest.Farm.CreatedAt,
			UpdatedAt:   farminvest.Farm.UpdatedAt,
			DeletedAt:   farminvest.Farm.DeletedAt,
		},
		Payment: payment.Payment{
			ID:         farminvest.Payment.ID,
			ExternalID: farminvest.Payment.ExternalID,
			InvoiceURL: farminvest.Payment.InvoiceURL,
			Status:     farminvest.Payment.Status,
			Amount:     farminvest.Payment.Amount,
		},
	}
}

func (u *FarmInvest) ToUseCase() *entities.FarmInvest {
	return &entities.FarmInvest{
		ID:               u.ID,
		CustomerID:       u.CustomerID,
		FarmID:           u.FarmID,
		PaymentID:        u.PaymentID,
		InvestmentAmount: u.InvestmentAmount,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
		DeletedAt:        u.DeletedAt,
		Farm: entities.Farm{
			ID:          u.Farm.ID,
			Title:       u.Farm.Title,
			Description: u.Farm.Description,
			Picture:     u.Farm.Picture,
			CreatedAt:   u.Farm.CreatedAt,
			UpdatedAt:   u.Farm.UpdatedAt,
			DeletedAt:   u.Farm.DeletedAt,
		},
		Payment: entities.Payment{
			ID:         u.Payment.ID,
			ExternalID: u.Payment.ExternalID,
			InvoiceURL: u.Payment.InvoiceURL,
			Status:     u.Payment.Status,
			Amount:     u.Payment.Amount,
		},
	}
}
