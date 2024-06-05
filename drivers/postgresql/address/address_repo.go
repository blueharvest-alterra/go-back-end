package address

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewAddressRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) GetAllStates(addresses *[]entities.Address) error {
	var addressesDb []Address

	err := GetAllState(&addressesDb)
	if err != nil {
		return err
	}

	for _, _address := range addressesDb {
		*addresses = append(*addresses, *_address.ToUseCase())
	}
	return nil
}

func (r *Repo) GetAllCities(addresses *[]entities.Address, stateID string) error {
	var addressesDb []Address

	err := GetAllCity(&addressesDb, stateID)
	if err != nil {
		return err
	}

	for _, _address := range addressesDb {
		*addresses = append(*addresses, *_address.ToUseCase())
	}
	return nil
}
