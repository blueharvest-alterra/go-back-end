package admin

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Login(admin *entities.Admin) error {
	adminDb := FromUseCase(admin)

	if err := r.DB.Joins("Auth").First(&adminDb).Error; err != nil {
		return err
	}

	*admin = *adminDb.ToUseCase()
	return nil
}
