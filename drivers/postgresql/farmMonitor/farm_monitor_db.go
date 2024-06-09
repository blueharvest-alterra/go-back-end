package farmmonitor

import (
	"time"

	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FarmMonitor struct {
	ID              uuid.UUID      `gorm:"type:varchar(100)"`
	FarmID          uuid.UUID      `gorm:"type:varchar(100)"`
	Temperature     float64        `gorm:"type:decimal"`
	PH              float64        `gorm:"type:decimal"`
	DissolvedOxygen float64        `gorm:"type:decimal"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func FromUseCase(farmmonitor *entities.FarmMonitor) *FarmMonitor {
	return &FarmMonitor{
		ID:              farmmonitor.ID,
		FarmID:          farmmonitor.FarmID,
		Temperature:     farmmonitor.Temperature,
		PH:              farmmonitor.PH,
		DissolvedOxygen: farmmonitor.DissolvedOxygen,
		CreatedAt:       farmmonitor.CreatedAt,
		UpdatedAt:       farmmonitor.UpdatedAt,
		DeletedAt:       farmmonitor.DeletedAt,
	}
}

func (u *FarmMonitor) ToUseCase() *entities.FarmMonitor {
	return &entities.FarmMonitor{
		ID:              u.ID,
		FarmID:          u.FarmID,
		Temperature:     u.Temperature,
		PH:              u.PH,
		DissolvedOxygen: u.DissolvedOxygen,
		CreatedAt:       u.CreatedAt,
		UpdatedAt:       u.UpdatedAt,
		DeletedAt:       u.DeletedAt,
	}
}
