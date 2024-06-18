package entities

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GenderEnum string

const (
	Male   GenderEnum = "male"
	Female GenderEnum = "female"
	Choose GenderEnum = "choose"
)

type Customer struct {
	ID          uuid.UUID
	FullName    string
	NickName    string
	PhoneNumber string
	Avatar      string
	Gender      GenderEnum
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
	GetProfile(customer *Customer) error
	EditProfile(customer *Customer) error
}

type CustomerUseCaseInterface interface {
	Login(customer *Customer) (Customer, error)
	Register(customer *Customer) (Customer, error)
	AddAddress(customer *Customer) (Customer, error)
	GetAddresses(customer *Customer) (Customer, error)
	GetProfile(customer *Customer) (Customer, error)
	EditProfile(customer *Customer, picture []*multipart.FileHeader) (Customer, error)
}
