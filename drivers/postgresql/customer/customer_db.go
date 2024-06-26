package customer

import (
	"time"

	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/address"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/auth"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GenderEnum string

const (
	Male   GenderEnum = "male"
	Female GenderEnum = "female"
	Choose GenderEnum = "choose"
)

type Customer struct {
	ID          uuid.UUID  `gorm:"type:varchar(100)"`
	FullName    string     `gorm:"type:varchar(100)"`
	NickName    string     `gorm:"type:varchar(100)"`
	PhoneNumber string     `gorm:"type:varchar(20)"`
	Avatar      string     `gorm:"type:varchar(100)"`
	Gender      GenderEnum `gorm:"type:gender;default:'choose'"`
	BirthDate   time.Time  `gorm:"type:date"`
	Auth        auth.Auth
	AuthID      uuid.UUID         `gorm:"type:varchar(100)"`
	CreatedAt   time.Time         `gorm:"autoCreateTime"`
	UpdatedAt   time.Time         `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt    `gorm:"index"`
	Addresses   []address.Address `gorm:"many2many:customer_addresses;"`
}

func FromUseCase(customer *entities.Customer) *Customer {
	addresses := make([]address.Address, len(customer.Addresses))
	for i, _address := range customer.Addresses {
		addresses[i] = address.Address{
			ID:        _address.ID,
			Address:   _address.Address,
			CityID:    _address.CityID,
			City:      _address.City,
			StateID:   _address.StateID,
			State:     _address.State,
			ZipCode:   _address.ZipCode,
			Country:   _address.Country,
			Longitude: _address.Longitude,
			Latitude:  _address.Latitude,
			CreatedAt: _address.CreatedAt,
			UpdatedAt: _address.UpdatedAt,
		}
	}

	return &Customer{
		ID:          customer.ID,
		FullName:    customer.FullName,
		PhoneNumber: customer.PhoneNumber,
		BirthDate:   customer.BirthDate,
		NickName:    customer.NickName,
		Gender:      GenderEnum(customer.Gender),
		Avatar:      customer.Avatar,
		Auth: auth.Auth{
			ID:       customer.Auth.ID,
			Email:    customer.Auth.Email,
			Password: customer.Auth.Password,
		},
		AuthID:    customer.AuthID,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
		DeletedAt: customer.DeletedAt,
		Addresses: addresses,
	}
}

func (u *Customer) ToUseCase() *entities.Customer {
	addresses := make([]entities.Address, len(u.Addresses))
	for i, _address := range u.Addresses {
		addresses[i] = entities.Address{
			ID:        _address.ID,
			Address:   _address.Address,
			CityID:    _address.CityID,
			City:      _address.City,
			StateID:   _address.StateID,
			State:     _address.State,
			ZipCode:   _address.ZipCode,
			Country:   _address.Country,
			Longitude: _address.Longitude,
			Latitude:  _address.Latitude,
			CreatedAt: _address.CreatedAt,
			UpdatedAt: _address.UpdatedAt,
		}
	}

	return &entities.Customer{
		ID:          u.ID,
		FullName:    u.FullName,
		PhoneNumber: u.PhoneNumber,
		BirthDate:   u.BirthDate,
		NickName:    u.NickName,
		Gender:      entities.GenderEnum(u.Gender),
		Avatar:      u.Avatar,
		Auth: entities.Auth{
			ID:       u.AuthID,
			Email:    u.Auth.Email,
			Password: u.Auth.Password,
		},
		AuthID:    u.AuthID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
		Addresses: addresses,
	}
}
