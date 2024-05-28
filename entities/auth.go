package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Auth struct {
	ID        uuid.UUID
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
