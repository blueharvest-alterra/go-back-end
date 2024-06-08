package cart

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID         uuid.UUID      `gorm:"type:varchar(100)"`
	CustomerID uuid.UUID      `gorm:"type:varchar(100)"`
	ProductID  uuid.UUID      `gorm:"type:varchar(100)"`
	Quantity   int64          `gorm:"type:decimal"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
