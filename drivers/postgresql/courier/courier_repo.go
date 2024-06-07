package courier

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func (r Repo) GetAll(couriers *[]entities.Courier, request entities.CostCourierAPIRequest) error {
	var couriersDb []Courier

	requestBody := RajaOngkirCostRequest{
		Origin:      request.OriginCityID,
		Destination: request.DestinationCityID,
		Weight:      request.Weight,
	}

	if err := GetAllAvailableCouriers(&couriersDb, requestBody); err != nil {
		return err
	}

	for _, _courier := range couriersDb {
		*couriers = append(*couriers, *_courier.ToUseCase())
	}
	return nil
}

func NewCourierRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}
