package customer

import (
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/auth"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	ID          uuid.UUID `gorm:"type:varchar(100)"`
	FullName    string    `gorm:"type:varchar(100)"`
	PhoneNumber string    `gorm:"type:varchar(20)"`
	BirthDate   time.Time `gorm:"type:date"`
	Auth        auth.Auth
	AuthID      uuid.UUID      `gorm:"type:varchar(100)"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func FromUseCase(customer *entities.Customer) *Customer {
	return &Customer{
		ID:          customer.ID,
		FullName:    customer.FullName,
		PhoneNumber: customer.PhoneNumber,
		BirthDate:   customer.BirthDate,
		Auth: auth.Auth{
			ID:       customer.Auth.ID,
			Email:    customer.Auth.Email,
			Password: customer.Auth.Password,
		},
		AuthID:    customer.AuthID,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
		DeletedAt: customer.DeletedAt,
	}
}

func (u *Customer) ToUseCase() *entities.Customer {
	return &entities.Customer{
		ID:          u.ID,
		FullName:    u.FullName,
		PhoneNumber: u.PhoneNumber,
		BirthDate:   u.BirthDate,
		Auth: entities.Auth{
			ID:       u.AuthID,
			Email:    u.Auth.Email,
			Password: u.Auth.Password,
		},
		AuthID:    u.AuthID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}
