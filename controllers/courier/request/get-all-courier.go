package request

import "github.com/blueharvest-alterra/go-back-end/entities"

type GetAllCourier struct {
	OriginCityID      string `json:"origin_city_id"`
	DestinationCityID string `json:"destination_city_id"`
	Weight            int    `json:"weight"`
}

func (ca *GetAllCourier) ToEntities() *entities.CostCourierAPIRequest {
	return &entities.CostCourierAPIRequest{
		OriginCityID:      ca.OriginCityID,
		DestinationCityID: ca.DestinationCityID,
		Weight:            ca.Weight,
	}
}
