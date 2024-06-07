package usecases

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
)

type CourierUseCase struct {
	repository entities.CourierRepositoryInterface
}

func (c CourierUseCase) GetAll(couriers *[]entities.Courier, request entities.CostCourierAPIRequest) ([]entities.Courier, error) {
	if request.OriginCityID == "" || request.DestinationCityID == "" || request.Weight < 1 {
		return []entities.Courier{}, constant.ErrEmptyInput
	}

	if err := c.repository.GetAll(couriers, request); err != nil {
		return nil, err
	}

	return *couriers, nil
}

func NewCourierUseCase(repository entities.CourierRepositoryInterface) *CourierUseCase {
	return &CourierUseCase{
		repository: repository,
	}
}
