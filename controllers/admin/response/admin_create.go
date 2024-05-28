package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type AdminAuth struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

type CreateAdmin struct {
	ID        uuid.UUID      `json:"id"`
	FullName  string         `json:"full_name"`
	Auth      AdminAuth      `json:"auth"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func CreateAdminFromUseCase(admin *entities.Admin) *CreateAdmin {
	return &CreateAdmin{
		ID:       admin.ID,
		FullName: admin.FullName,
		Auth: AdminAuth{
			ID:    admin.Auth.ID,
			Email: admin.Auth.Email,
		},
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
		DeletedAt: admin.DeletedAt,
	}
}
