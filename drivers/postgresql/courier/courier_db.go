package courier

import (
	"github.com/google/uuid"
)

type Courier struct {
	ID   uuid.UUID `gorm:"type:varchar(100)"`
	Name string    `gorm:"type:varchar(255);not null"`
	Fee  float64   `gorm:"type:decimal;not null"`
	Type string    `gorm:"type:varchar(50);not null"`
}
