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

func (r *Repo) GetProfile(customer *entities.Customer) error {
	customerDb := FromUseCase(customer)

	if err := r.DB.First(&customerDb, "id = ?", customerDb.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	customerAuth := auth.Auth{ID: customerDb.AuthID}
	if err := r.DB.Model(&customerAuth).Where("id = ?", customerDb.AuthID).First(&customerAuth).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	customerDb.Auth.Email = customerAuth.Email

	*customer = *customerDb.ToUseCase()
	return nil
}

func (r *Repo) EditProfile(customer *entities.Customer) error {
	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	customerDb := FromUseCase(customer)

	if err := tx.Omit("Addresses").Where("id = ?", customerDb.ID).Updates(customerDb).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	if err := tx.Where("id = ?", customerDb.ID).First(&customerDb).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	var existingAuth entities.Auth
	if err := tx.Where("email = ?", customerDb.Auth.Email).First(&existingAuth).Error; err == nil {
		tx.Rollback()
		return errors.New("email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return err
	}

	if err := tx.Where("id = ?", customerDb.AuthID).Updates(customerDb.Auth).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	customer.Auth.Email = customerDb.Auth.Email

	*customer = *customerDb.ToUseCase()
	return nil
}
