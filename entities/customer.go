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
	Addresses   []Address
}

type CustomerRepositoryInterface interface {
	Login(customer *Customer) error
	Register(customer *Customer) error
	AddAddress(customer *Customer) error
	GetAddresses(customer *Customer) error
}

type CustomerUseCaseInterface interface {
	Login(customer *Customer) (Customer, error)
	Register(customer *Customer) (Customer, error)
	AddAddress(customer *Customer) (Customer, error)
	GetAddresses(customer *Customer) (Customer, error)
}
