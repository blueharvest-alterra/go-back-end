package entities

import (
	"time"

	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PromoStatus string

const (
	Available   PromoStatus = "available"
	Unavailable PromoStatus = "unavailable"
)

type Promo struct {
	ID        uuid.UUID
	Name      string
	Code      string
	Status    PromoStatus
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type PromoRepositoryInterface interface {
	Create(promo *Promo) error
	GetById(promo *Promo) error
	Update(promo *Promo) error
	Delete(promo *Promo) error
	GetAll(promo *[]Promo) error
}

type PromoUseCaseInterface interface {
	Create(promo *Promo, userData *middlewares.Claims) (Promo, error)
	GetById(id uuid.UUID) (Promo, error)
	Update(promo *Promo, userData *middlewares.Claims) (Promo, error)
	Delete(id uuid.UUID, userData *middlewares.Claims) (Promo, error)
	GetAll(promo *[]Promo, userData *middlewares.Claims) ([]Promo, error)
}
