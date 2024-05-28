package auth

import (
	"github.com/google/uuid"
)

type Auth struct {
	ID       uuid.UUID `gorm:"type:varchar(100)"`
	Email    string    `gorm:"type:varchar(255);unique"`
	Password string    `gorm:"type:varchar(255)"`
}
