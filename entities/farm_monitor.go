package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FarmMonitor struct {
	ID              uuid.UUID
	FarmID          uuid.UUID
	Temperature     float64
	PH              float64
	DissolvedOxygen float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type FarmMonitorRepositoryInterface interface {
	Create(farmMonitor *FarmMonitor) error
	GetById(farmMonitor *FarmMonitor) error
	Update(farmMonitor *FarmMonitor) error
	GetAllByFarmId(farmId uuid.UUID, farmMonitors *[]FarmMonitor) error
}

type FarmMonitorUseCaseInterface interface {
	Create(farmMonitor *FarmMonitor) (FarmMonitor, error)
	GetById(farmMonitorId uuid.UUID) (FarmMonitor, error)
	Update(farmMonitor *FarmMonitor) (FarmMonitor, error)
	GetAllByFarmId(farmId uuid.UUID) ([]FarmMonitor, error)
}
