package customer

import (
	"errors"
	"github.com/blueharvest-alterra/go-back-end/constant"
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

	if err := r.DB.Joins("Auth").First(&customerDb).Error; err != nil {
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
