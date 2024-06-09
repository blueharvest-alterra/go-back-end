package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FarmInvest struct {
	ID               uuid.UUID
	CustomerID       uuid.UUID
	FarmID           uuid.UUID
	InvestmentAmount float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

type FarmInvestRepositoryInterface interface {
	Create(farmInvest *FarmInvest) error
	GetById(farmInvest *FarmInvest) error
	GetAll(customerID uuid.UUID, farmInvests *[]FarmInvest) error
}

type FarmInvestUseCaseInterface interface {
	Create(farmInvest *FarmInvest) (FarmInvest, error)
	GetById(id uuid.UUID) (FarmInvest, error)
	GetAll(customerID uuid.UUID) ([]FarmInvest, error)
}
