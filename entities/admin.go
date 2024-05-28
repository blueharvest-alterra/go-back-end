package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	ID        uuid.UUID
	FullName  string
	Auth      Auth
	AuthID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type AdminRepositoryInterface interface {
	Login(admin *Admin) error
}

type AdminUseCaseInterface interface {
	Login(admin *Admin) (Admin, error)
}
