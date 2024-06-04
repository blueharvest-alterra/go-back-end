package address

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Address struct {
	ID        uuid.UUID      `gorm:"type:varchar(100);"`
	Address   string         `gorm:"type:varchar(255);not null"`
	City      string         `gorm:"type:varchar(100);not null"`
	State     string         `gorm:"type:varchar(50);not null"`
	ZipCode   string         `gorm:"type:varchar(10);not null"`
	Country   string         `gorm:"type:varchar(50);not null"`
	Longitude string         `gorm:"type:varchar(100);not null"`
	Latitude  string         `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
