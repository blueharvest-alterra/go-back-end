package usecases

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type FarmInvestUseCase struct {
	repository entities.FarmInvestRepositoryInterface
}

func NewFarmInvestUseCase(repository entities.FarmInvestRepositoryInterface) *FarmInvestUseCase {
	return &FarmInvestUseCase{repository: repository}
}

func (c *FarmInvestUseCase) Create(farmInvest *entities.FarmInvest) (entities.FarmInvest, error) {
	farmInvest.ID = uuid.New()

	if err := c.repository.Create(farmInvest); err != nil {
		return entities.FarmInvest{}, err
	}

	return *farmInvest, nil
}

func (c *FarmInvestUseCase) GetById(id uuid.UUID) (entities.FarmInvest, error) {
    var farmInvest entities.FarmInvest
    farmInvest.ID = id  

    if err := c.repository.GetById(&farmInvest); err != nil {
        return entities.FarmInvest{}, err
    }

    return farmInvest, nil
}



func (c *FarmInvestUseCase) GetAll(customerID uuid.UUID) ([]entities.FarmInvest, error) {
	var farmInvests []entities.FarmInvest

	if err := c.repository.GetAll(customerID, &farmInvests); err != nil {
		return []entities.FarmInvest{}, err
	}

	return farmInvests, nil
}