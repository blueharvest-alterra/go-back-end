package entities

import (
	"github.com/google/uuid"
)

type Courier struct {
	ID                   uuid.UUID
	DestinationAddressID uuid.UUID
	Name                 string
	Fee                  float64
	Type                 string
}

type CostCourierAPIRequest struct {
	OriginCityID      string
	DestinationCityID string
	Weight            int
}

type CourierRepositoryInterface interface {
	GetAll(couriers *[]Courier, request CostCourierAPIRequest) error
}

type CourierUseCaseInterface interface {
	GetAll(couriers *[]Courier, request CostCourierAPIRequest) ([]Courier, error)
}
