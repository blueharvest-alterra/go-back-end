package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	ID          uuid.UUID
	FullName    string
	PhoneNumber string
	BirthDate   time.Time
	Auth        Auth
	AuthID      uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type CustomerRepositoryInterface interface {
	Login(admin *Customer) error
	Register(admin *Customer) error
}

type CustomerUseCaseInterface interface {
	Login(admin *Customer) (Customer, error)
	Register(admin *Customer) (Customer, error)
}
