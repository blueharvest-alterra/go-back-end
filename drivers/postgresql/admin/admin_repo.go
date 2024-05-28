package admin

import (
	"errors"
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/auth"
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

	adminAuth := auth.Auth{Email: adminDb.Auth.Email}
	if err := r.DB.Model(&adminAuth).Where("email = ?", adminDb.Auth.Email).First(&adminAuth).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrInvalidEmailOrPassword
		}
		return err
	}

	adminDb.Auth.ID = adminAuth.ID
	adminDb.Auth.Email = adminAuth.Email
	adminDb.Auth.Password = adminAuth.Password

	if err := r.DB.Model(&adminDb).Where("auth_id = ?", adminAuth.ID).First(&adminDb).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrInvalidEmailOrPassword
		}
		return err
	}

	*admin = *adminDb.ToUseCase()
	return nil
}

func (r *Repo) Create(admin *entities.Admin) error {
	adminDb := FromUseCase(admin)

	if err := r.DB.Create(&adminDb).Error; err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return constant.ErrDuplicatedData
		}
		return err
	}

	*admin = *adminDb.ToUseCase()
	return nil
}
