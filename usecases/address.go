package usecases

import "github.com/blueharvest-alterra/go-back-end/entities"

type AddressUseCase struct {
	repository entities.AddressRepositoryInterface
}

func NewAddressUseCase(repository entities.AddressRepositoryInterface) *AddressUseCase {
	return &AddressUseCase{
		repository: repository,
	}
}

func (a *AddressUseCase) GetAllStates(addresses *[]entities.Address) ([]entities.Address, error) {
	if err := a.repository.GetAllStates(addresses); err != nil {
		return []entities.Address{}, err
	}

	return *addresses, nil
}

func (a *AddressUseCase) GetAllCities(addresses *[]entities.Address, stateID string) ([]entities.Address, error) {
	if err := a.repository.GetAllCities(addresses, stateID); err != nil {
		return []entities.Address{}, err
	}

	return *addresses, nil
}
