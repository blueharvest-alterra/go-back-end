package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type customerAddress struct {
	ID        uuid.UUID `json:"id"`
	Address   string    `json:"address"`
	CityID    string    `json:"city_id"`
	City      string    `json:"city"`
	StateID   string    `json:"state_id"`
	State     string    `json:"state"`
	ZipCode   string    `json:"zip_code"`
	Country   string    `json:"country"`
	Longitude string    `json:"longitude"`
	Latitude  string    `json:"latitude"`
}

type GetAddressesResponse struct {
	CustomerID uuid.UUID         `json:"customer_id"`
	Addresses  []customerAddress `json:"addresses"`
}

func AddressesResponseFromUseCase(customer *entities.Customer) *GetAddressesResponse {
	addresses := make([]customerAddress, len(customer.Addresses))
	for i, address := range customer.Addresses {
		addresses[i] = customerAddress{
			ID:        address.ID,
			Address:   address.Address,
			CityID:    address.CityID,
			City:      address.City,
			StateID:   address.StateID,
			State:     address.State,
			ZipCode:   address.ZipCode,
			Country:   address.Country,
			Longitude: address.Longitude,
			Latitude:  address.Latitude,
		}
	}
	return &GetAddressesResponse{
		CustomerID: customer.ID,
		Addresses:  addresses,
	}
}
