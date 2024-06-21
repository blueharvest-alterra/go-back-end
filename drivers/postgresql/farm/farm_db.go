package farm

import (
	"time"

	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Farm struct {
	ID                      uuid.UUID      `gorm:"type:varchar(100)"`
	Title                   string         `gorm:"type:varchar(100)"`
	Description             string         `gorm:"type:varchar(100)"`
	Picture                 string         `gorm:"type:varchar(100)"`
	MinimumInvestmentAmount float64        `gorm:"type:decimal"`
	CreatedAt               time.Time      `gorm:"autoCreateTime"`
	UpdatedAt               time.Time      `gorm:"autoUpdateTime"`
	DeletedAt               gorm.DeletedAt `gorm:"index"`
}

func FromUseCase(farm *entities.Farm) *Farm {
	return &Farm{
		ID:                      farm.ID,
		Title:                   farm.Title,
		Description:             farm.Description,
		Picture:                 farm.Picture,
		MinimumInvestmentAmount: farm.MinimumInvestmentAmount,
		CreatedAt:               farm.CreatedAt,
		UpdatedAt:               farm.UpdatedAt,
		DeletedAt:               farm.DeletedAt,
	}
}

func (u *Farm) ToUseCase() *entities.Farm {
	return &entities.Farm{
		ID:                      u.ID,
		Title:                   u.Title,
		Description:             u.Description,
		Picture:                 u.Picture,
		MinimumInvestmentAmount: u.MinimumInvestmentAmount,
		CreatedAt:               u.CreatedAt,
		UpdatedAt:               u.UpdatedAt,
		DeletedAt:               u.DeletedAt,
	}
}
