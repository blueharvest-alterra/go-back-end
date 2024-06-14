package entities

import (
	"time"

	"github.com/blueharvest-alterra/go-back-end/middlewares"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FarmInvest struct {
	ID               uuid.UUID
	CustomerID       uuid.UUID
	FarmID           uuid.UUID
	PaymentID        uuid.UUID
	InvestmentAmount float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	Farm             Farm
	Payment          Payment
	Customer         Customer
}

type FarmInvestRepositoryInterface interface {
	Create(farmInvest *FarmInvest) error
	GetById(farmInvest *FarmInvest, userData *middlewares.Claims) error
	GetAll(farmInvests *[]FarmInvest, userData *middlewares.Claims) error
}

type FarmInvestUseCaseInterface interface {
	Create(farmInvest *FarmInvest, userData *middlewares.Claims) (FarmInvest, error)
	GetById(farmInvest *FarmInvest, userData *middlewares.Claims) (FarmInvest, error)
	GetAll(farmInvests *[]FarmInvest, userData *middlewares.Claims) ([]FarmInvest, error)
}
