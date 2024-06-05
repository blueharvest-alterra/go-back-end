package entities

import (
	"github.com/google/uuid"
	"time"
)

type Address struct {
	ID        uuid.UUID
	Address   string
	CityID    string
	City      string
	StateID   string
	State     string
	ZipCode   string
	Country   string
	Longitude string
	Latitude  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AddressRepositoryInterface interface {
	GetAllStates(addresses *[]Address) error
	GetAllCities(addresses *[]Address, stateID string) error
}

type AddressUseCaseInterface interface {
	GetAllStates(addresses *[]Address) ([]Address, error)
	GetAllCities(addresses *[]Address, stateID string) ([]Address, error)
}
