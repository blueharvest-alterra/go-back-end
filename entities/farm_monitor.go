package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FarmMonitor struct {
	ID              uuid.UUID
	FarmID          uuid.UUID
	Farm            Farm
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
	GetAll(farmMonitors *[]FarmMonitor) error
}

type FarmMonitorUseCaseInterface interface {
	Create(farmMonitor *FarmMonitor) (FarmMonitor, error)
	GetById(id uuid.UUID) (FarmMonitor, error)
	Update(farmMonitor *FarmMonitor) (FarmMonitor, error)
	GetAll(farmMonitors *[]FarmMonitor) ([]FarmMonitor, error)
}
