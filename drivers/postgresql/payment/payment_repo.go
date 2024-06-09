package payment

import (
	"errors"
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func (r Repo) UpdateStatus(payment *entities.Payment) error {
	paymentDb := FromUseCase(payment)

	if err := r.DB.Model(&paymentDb).Updates(&paymentDb).Where("id = ?", paymentDb.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	*payment = *paymentDb.ToUseCase()
	return nil
}

func NewPaymentRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}
