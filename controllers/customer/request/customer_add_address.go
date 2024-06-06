package request

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type CustomerAddAddress struct {
	Address   string `json:"address"`
	CityID    string `json:"city_id"`
	City      string `json:"city"`
	StateID   string `json:"state_id"`
	State     string `json:"state"`
	ZipCode   string `json:"zip_code"`
	Country   string `json:"country"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

func (r *CustomerAddAddress) AddAddressToEntities(customerId uuid.UUID) *entities.Customer {
	customer := &entities.Customer{}
	customer.ID = customerId
	newAddress := entities.Address{
		Address:   r.Address,
		CityID:    r.CityID,
		City:      r.City,
		StateID:   r.StateID,
		State:     r.State,
		ZipCode:   r.ZipCode,
		Country:   r.Country,
		Longitude: r.Longitude,
		Latitude:  r.Latitude,
	}
	customer.Addresses = append(customer.Addresses, newAddress)
	return customer
}
