package admin

import (
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/auth"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	ID        uuid.UUID `gorm:"type:varchar(100);"`
	FullName  string    `gorm:"type:varchar(100);not null"`
	Auth      auth.Auth
	AuthID    uuid.UUID      `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromUseCase(admin *entities.Admin) *Admin {
	return &Admin{
		ID:       admin.ID,
		FullName: admin.FullName,
		Auth: auth.Auth{
			ID:       admin.Auth.ID,
			Email:    admin.Auth.Email,
			Password: admin.Auth.Password,
		},
		AuthID:    admin.AuthID,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
		DeletedAt: admin.DeletedAt,
	}
}

func (u *Admin) ToUseCase() *entities.Admin {
	return &entities.Admin{
		ID:       u.ID,
		FullName: u.FullName,
		Auth: entities.Auth{
			ID:       u.Auth.ID,
			Email:    u.Auth.Email,
			Password: u.Auth.Password,
		},
		AuthID:    u.AuthID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}
