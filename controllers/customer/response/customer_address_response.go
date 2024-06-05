package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type CustomerAddressResponse struct {
	CustomerID uuid.UUID `json:"customer_id"`
	ID         uuid.UUID `json:"id"`
	Address    string    `json:"address"`
	CityID     string    `json:"city_id"`
	City       string    `json:"city"`
	StateID    string    `json:"state_id"`
	State      string    `json:"state"`
	ZipCode    string    `json:"zip_code"`
	Country    string    `json:"country"`
	Longitude  string    `json:"longitude"`
	Latitude   string    `json:"latitude"`
}

func AddressResponseFromUseCase(u *entities.Customer) *CustomerAddressResponse {
	address := &CustomerAddressResponse{}
	address.ID = u.Addresses[0].ID
	address.CustomerID = u.ID
	address.Address = u.Addresses[0].Address
	address.CityID = u.Addresses[0].CityID
	address.City = u.Addresses[0].City
	address.StateID = u.Addresses[0].StateID
	address.State = u.Addresses[0].State
	address.ZipCode = u.Addresses[0].ZipCode
	address.Country = u.Addresses[0].Country
	address.Longitude = u.Addresses[0].Longitude
	address.Latitude = u.Addresses[0].Latitude
	return address
}
