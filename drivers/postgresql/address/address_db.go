package address

import (
	"encoding/json"
	"fmt"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

type Address struct {
	ID        uuid.UUID      `gorm:"type:varchar(100);"`
	Address   string         `gorm:"type:varchar(255);not null"`
	CityID    string         `json:"city_id"`
	City      string         `gorm:"type:varchar(100);not null" json:"city_name"`
	StateID   string         `json:"province_id"`
	State     string         `gorm:"type:varchar(50);not null" json:"province"`
	ZipCode   string         `gorm:"type:varchar(10);not null"`
	Country   string         `gorm:"type:varchar(50);not null"`
	Longitude string         `gorm:"type:varchar(100);not null"`
	Latitude  string         `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromUseCase(address *entities.Address) *Address {
	return &Address{
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
		CreatedAt: address.CreatedAt,
		UpdatedAt: address.UpdatedAt,
	}
}

func (a *Address) ToUseCase() *entities.Address {
	return &entities.Address{
		ID:        a.ID,
		Address:   a.Address,
		CityID:    a.CityID,
		City:      a.City,
		StateID:   a.StateID,
		State:     a.State,
		ZipCode:   a.ZipCode,
		Country:   a.Country,
		Longitude: a.Longitude,
		Latitude:  a.Latitude,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

type RajaongkirResponse struct {
	Rajaongkir struct {
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Results []Address `json:"results"`
	} `json:"rajaongkir"`
}

func GetAllState(addresses *[]Address) error {
	url := fmt.Sprintf("https://api.rajaongkir.com/starter/province")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("key", os.Getenv("RAJAONGKIR_API_KEY"))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get all states, server responded with %s", res.Status)
	}

	var response RajaongkirResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}

	*addresses = response.Rajaongkir.Results

	if response.Rajaongkir.Status.Code != 200 {
		return fmt.Errorf("failed to get all states, API responded with code %d: %s", response.Rajaongkir.Status.Code, response.Rajaongkir.Status.Description)
	}

	return nil
}

func GetAllCity(addresses *[]Address, stateID string) error {
	url := fmt.Sprintf("https://api.rajaongkir.com/starter/city?province=%s", stateID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("key", os.Getenv("RAJAONGKIR_API_KEY"))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get all states, server responded with %s", res.Status)
	}

	var response RajaongkirResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}

	*addresses = response.Rajaongkir.Results

	if response.Rajaongkir.Status.Code != 200 {
		return fmt.Errorf("failed to get all states, API responded with code %d: %s", response.Rajaongkir.Status.Code, response.Rajaongkir.Status.Description)
	}

	return nil
}
