package usecases

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
)

type FarmMonitorUseCase struct {
	repository entities.FarmMonitorRepositoryInterface
}

func NewFarmMonitorUseCase(repository entities.FarmMonitorRepositoryInterface) *FarmMonitorUseCase {
	return &FarmMonitorUseCase{repository: repository}
}

func (c *FarmMonitorUseCase) Create(farmMonitor *entities.FarmMonitor) (entities.FarmMonitor, error) {
	farmMonitor.ID = uuid.New()

	if err := c.repository.Create(farmMonitor); err != nil {
		return entities.FarmMonitor{}, err
	}

	return *farmMonitor, nil
}

func (c *FarmMonitorUseCase) Update(farmMonitor *entities.FarmMonitor, userData *middlewares.Claims) (entities.FarmMonitor, error) {
	if userData.Role != "admin" {
		return entities.FarmMonitor{}, constant.ErrNotAuthorized
	}

	if farmMonitor.FarmID == uuid.Nil || farmMonitor.Temperature == 0.0 || farmMonitor.PH == 0.0  || farmMonitor.DissolvedOxygen == 0.0  {
		return entities.FarmMonitor{}, constant.ErrEmptyInput
	}

	if err := c.repository.Update(farmMonitor); err != nil {
		return entities.FarmMonitor{}, err
	}

	return *farmMonitor, nil
}

func (c *FarmMonitorUseCase) GetById(id uuid.UUID) (entities.FarmMonitor, error) {
	var farmMonitor entities.FarmMonitor
	farmMonitor.ID = id

	if err := c.repository.GetById(&farmMonitor); err != nil {
		return entities.FarmMonitor{}, err
	}

	return farmMonitor, nil
}

func (c *FarmMonitorUseCase) GetAllByFarmId(farmId uuid.UUID) ([]entities.FarmMonitor, error) {
	var farmMonitors []entities.FarmMonitor
	if err := c.repository.GetAllByFarmId(farmId, &farmMonitors); err != nil {
		return nil, err
	}
	return farmMonitors, nil
}
