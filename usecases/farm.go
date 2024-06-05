package usecases

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type FarmUseCase struct {
	repository entities.FarmRepositoryInterface
}

func NewFarmUseCase(repository entities.FarmRepositoryInterface) *FarmUseCase {
	return &FarmUseCase{repository: repository}
}

func (c *FarmUseCase) Create(farm *entities.Farm) (entities.Farm, error) {
	farm.ID = uuid.New()

	if err := c.repository.Create(farm); err != nil {
		return entities.Farm{}, err
	}

	return *farm, nil
}

func (c *FarmUseCase) GetById(id uuid.UUID) (entities.Farm, error) {
    var farm entities.Farm
    farm.ID = id  

    if err := c.repository.GetById(&farm); err != nil {
        return entities.Farm{}, err
    }

    return farm, nil
}

func (c *FarmUseCase) Update(farm *entities.Farm) (entities.Farm, error) {
	if err := c.repository.Update(farm); err != nil {
		return entities.Farm{}, err
	}

	return *farm, nil
}

func (c *FarmUseCase) Delete(id uuid.UUID) (entities.Farm, error) {
    var farm entities.Farm
    farm.ID = id  

    if err := c.repository.Delete(&farm); err != nil {
        return entities.Farm{}, err
    }

    return farm, nil
}

func (c *FarmUseCase) GetAll(farm *[]entities.Farm) ([]entities.Farm, error) {
	if err := c.repository.GetAll(farm); err != nil {
		return []entities.Farm{}, err
	}

	return *farm, nil
}
