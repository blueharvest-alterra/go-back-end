package usecases

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
)

type DashboardUseCase struct {
	repository entities.DashboardRepositoryInterface
}

func (d DashboardUseCase) AdminDashboard(dashboard *entities.Dashboard, userData *middlewares.Claims) (entities.Dashboard, error) {
	if userData.Role != "admin" {
		return entities.Dashboard{}, constant.ErrNotAuthorized
	}

	if err := d.repository.AdminDashboard(dashboard); err != nil {
		return entities.Dashboard{}, err
	}

	return *dashboard, nil
}

func (d DashboardUseCase) CustomerDashboard(dashboard *entities.Dashboard, userData *middlewares.Claims) (entities.Dashboard, error) {
	if userData.Role != "customer" {
		return entities.Dashboard{}, constant.ErrNotAuthorized
	}

	if err := d.repository.CustomerDashboard(dashboard); err != nil {
		return entities.Dashboard{}, err
	}

	return *dashboard, nil
}

func NewDashboardUseCase(repository entities.DashboardRepositoryInterface) *DashboardUseCase {
	return &DashboardUseCase{
		repository: repository,
	}
}
