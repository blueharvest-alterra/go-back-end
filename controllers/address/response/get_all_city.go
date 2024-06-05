package response

import "github.com/blueharvest-alterra/go-back-end/entities"

type CityDetail struct {
	ID   string `json:"id"`
	City string `json:"name"`
}

type CityGetAll struct {
	Cities []CityDetail `json:"cities"`
}

func GetAllCityFromUseCase(addresses *[]entities.Address) *CityGetAll {
	allCategories := make([]CityDetail, len(*addresses))
	for i, _address := range *addresses {
		allCategories[i] = CityDetail{
			ID:   _address.CityID,
			City: _address.City,
		}
	}

	return &CityGetAll{
		Cities: allCategories,
	}
}
