package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Farm struct {
	ID          uuid.UUID
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type FarmRepositoryInterface interface {
	Create(farm *Farm) error
	GetById(farm *Farm) error
	Update(farm *Farm) error
	Delete(farm *Farm) error
	GetAll(farms *[]Farm) error
}

type FarmUseCaseInterface interface {
	Create(farm *Farm) (Farm, error)
	GetById(id uuid.UUID) (Farm, error)
	Update(farm *Farm) (Farm, error)
	Delete(id uuid.UUID) (Farm, error)
	GetAll(farms *[]Farm) ([]Farm, error)
}
