package farmmonitor

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FarmMonitor struct {
	ID              uuid.UUID      `gorm:"type:varchar(100)"`
	FarmID          uuid.UUID      `gorm:"type:varchar(100)"`
	Temperature     float64        `gorm:"type:varchar(100)"`
	PH              float64        `gorm:"type:decimal"`
	DissolvedOxygen float64        `gorm:"type:decimal"`
	TotalPrice      float64        `gorm:"type:decimal"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
