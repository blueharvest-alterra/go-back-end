package response

import "github.com/blueharvest-alterra/go-back-end/entities"

type CourierDetail struct {
	Name string  `json:"name"`
	Fee  float64 `json:"fee"`
	Type string  `json:"type"`
}

type GetAllCourierResponse struct {
	Couriers []CourierDetail `json:"couriers"`
}

func GetAllCourierFromUseCase(couriers *[]entities.Courier) *GetAllCourierResponse {
	var allCouriers = make([]CourierDetail, len(*couriers))
	for i, courier := range *couriers {
		allCouriers[i] = CourierDetail{
			Name: courier.Name,
			Fee:  courier.Fee,
			Type: courier.Type,
		}
	}
	return &GetAllCourierResponse{
		Couriers: allCouriers,
	}
}
