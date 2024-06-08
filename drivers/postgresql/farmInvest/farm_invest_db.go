package farminvest

import (
	"time"

	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FarmInvest struct {
	ID               uuid.UUID      `gorm:"type:varchar(100)"`
	CustomerID       uuid.UUID      `gorm:"type:varchar(100)"`
	FarmID           uuid.UUID      `gorm:"type:varchar(100)"`
	InvestmentAmount float64        `gorm:"type:decimal"`
	CreatedAt        time.Time      `gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

func FromUseCase(farminvest *entities.FarmInvest) *FarmInvest {
	return &FarmInvest{
		ID:               farminvest.ID,
		CustomerID:       farminvest.CustomerID,
		FarmID:           farminvest.FarmID,
		InvestmentAmount: farminvest.InvestmentAmount,
		CreatedAt:        farminvest.CreatedAt,
		UpdatedAt:        farminvest.UpdatedAt,
		DeletedAt:        farminvest.DeletedAt,
	}
}

func (u *FarmInvest) ToUseCase() *entities.FarmInvest {
	return &entities.FarmInvest{
		ID:               u.ID,
		CustomerID:       u.CustomerID,
		FarmID:           u.FarmID,
		InvestmentAmount: u.InvestmentAmount,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
		DeletedAt:        u.DeletedAt,
	}
}
