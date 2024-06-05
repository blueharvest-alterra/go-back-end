package customer

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

func NewCustomerRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Login(customer *entities.Customer) error {
	customerDb := FromUseCase(customer)

	customerAuth := auth.Auth{Email: customerDb.Auth.Email}
	if err := r.DB.Model(&customerAuth).Where("email = ?", customerDb.Auth.Email).First(&customerAuth).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrInvalidEmailOrPassword
		}
		return err
	}

	customerDb.Auth.ID = customerAuth.ID
	customerDb.Auth.Email = customerAuth.Email
	customerDb.Auth.Password = customerAuth.Password

	if err := r.DB.Model(&customerDb).Where("auth_id = ?", customerAuth.ID).First(&customerDb).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrInvalidEmailOrPassword
		}
		return err
	}

	*customer = *customerDb.ToUseCase()
	return nil
}

func (r *Repo) Register(customer *entities.Customer) error {
	customerDb := FromUseCase(customer)

	if err := r.DB.Create(&customerDb).Error; err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return constant.ErrDuplicatedData
		}
		return err
	}

	*customer = *customerDb.ToUseCase()
	return nil
}

func (r *Repo) AddAddress(user *entities.Customer) error {
	customerDb := FromUseCase(user)

	if err := r.DB.Model(&customerDb).Association("Addresses").Append(&customerDb); err != nil {
		return constant.ErrInsertDatabase
	}

	*user = *customerDb.ToUseCase()
	return nil
}

func (r *Repo) GetAddresses(customer *entities.Customer) error {
	customerDb := FromUseCase(customer)

	if err := r.DB.Preload("Addresses").First(&customerDb, "id = ?", customerDb.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	*customer = *customerDb.ToUseCase()
	return nil
}
