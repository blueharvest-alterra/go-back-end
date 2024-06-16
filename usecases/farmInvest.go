package usecases

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
)

type FarmInvestUseCase struct {
	repository entities.FarmInvestRepositoryInterface
}

func NewFarmInvestUseCase(repository entities.FarmInvestRepositoryInterface) *FarmInvestUseCase {
	return &FarmInvestUseCase{repository: repository}
}

func (c *FarmInvestUseCase) Create(farmInvest *entities.FarmInvest, userData *middlewares.Claims) (entities.FarmInvest, error) {
	if userData.Role != "customer" {
		return entities.FarmInvest{}, constant.ErrNotAuthorized
	}

	if farmInvest.FarmID == uuid.Nil || farmInvest.InvestmentAmount == 0 {
		return entities.FarmInvest{}, constant.ErrEmptyInput
	}

	farmInvest.ID = uuid.New()
	farmInvest.Customer.ID = userData.ID
	farmInvest.Customer.FullName = userData.FullName
	farmInvest.Customer.PhoneNumber = ""

	if err := c.repository.Create(farmInvest); err != nil {
		return entities.FarmInvest{}, err
	}

	return *farmInvest, nil
}

func (c *FarmInvestUseCase) GetById(farmInvest *entities.FarmInvest, userData *middlewares.Claims) (entities.FarmInvest, error) {
	if farmInvest.ID == uuid.Nil {
		return entities.FarmInvest{}, constant.ErrEmptyInput
	}

	if userData.Role == "customer" {
		farmInvest.CustomerID = userData.ID
	}

	if err := c.repository.GetById(farmInvest, userData); err != nil {
		return entities.FarmInvest{}, err
	}

	return *farmInvest, nil
}

func (c *FarmInvestUseCase) GetAll(farmInvests *[]entities.FarmInvest, userData *middlewares.Claims) ([]entities.FarmInvest, error) {

	if err := c.repository.GetAll(farmInvests, userData); err != nil {
		return []entities.FarmInvest{}, err
	}

	return *farmInvests, nil
}
