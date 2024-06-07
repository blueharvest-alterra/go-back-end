package courier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
)

type Courier struct {
	ID   uuid.UUID `gorm:"type:varchar(100)"`
	Name string    `gorm:"type:varchar(255);not null"`
	Fee  float64   `gorm:"type:decimal;not null"`
	Type string    `gorm:"type:varchar(50);not null"`
}

func FromUseCase(courier *entities.Courier) *Courier {
	return &Courier{
		ID:   courier.ID,
		Name: courier.Name,
		Fee:  courier.Fee,
		Type: courier.Type,
	}
}

func (c *Courier) ToUseCase() *entities.Courier {
	return &entities.Courier{
		ID:   c.ID,
		Name: c.Name,
		Fee:  c.Fee,
		Type: c.Type,
	}
}

type RajaOngkirCostRequest struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Weight      int    `json:"weight"`
	Courier     string `json:"courier"`
}

type RajaOngkirCostResponse struct {
	RajaOngkir struct {
		Results []struct {
			Code  string `json:"code"`
			Name  string `json:"name"`
			Costs []struct {
				Service string `json:"service"`
				Cost    []struct {
					Value float64 `json:"value"`
				} `json:"cost"`
			} `json:"costs"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

func GetAllAvailableCouriers(couriers *[]Courier, request RajaOngkirCostRequest) error {
	url := "https://api.rajaongkir.com/starter/cost"

	availableCouriers := []string{"jne", "pos", "tiki"}

	for _, courier := range availableCouriers {
		fmt.Println("courier", courier)
		request.Courier = courier
		jsonRequestBody, err := json.Marshal(request)
		if err != nil {
			fmt.Println("Error marshaling request body:", err)
			return err
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBody))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return err
		}

		req.Header.Add("key", os.Getenv("RAJAONGKIR_API_KEY"))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			fmt.Printf("Failed to get available couriers, server responded with %s\n", res.Status)
			return nil
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return err
		}

		var response RajaOngkirCostResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println("Error unmarshaling JSON response:", err)
			return err
		}

		for _, result := range response.RajaOngkir.Results {
			for _, cost := range result.Costs {
				courier := Courier{
					Name: result.Name,
					Fee:  cost.Cost[0].Value,
					Type: cost.Service,
				}
				*couriers = append(*couriers, courier)
			}
		}
	}

	return nil
}
