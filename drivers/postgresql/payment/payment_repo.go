package payment

import (
	"errors"

	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/farm"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/product"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/transactionDetail"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func (r Repo) UpdateBuyProductContext(payment *entities.Payment) error {
	paymentDb := FromUseCase(payment)

	var transactionDetailsDb []transactionDetail.TransactionDetail
	if err := r.DB.Preload("Product").Where("transaction_id = ?", paymentDb.ContextID).Find(&transactionDetailsDb).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	for _, _transactionDetail := range transactionDetailsDb {
		if err := r.DB.Model(product.Product{}).Where("id = ?", _transactionDetail.ProductID).Update("count_sold", _transactionDetail.Product.CountSold+_transactionDetail.Quantity).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return constant.ErrNotFound
			}
			return err
		}
	}

	*payment = *paymentDb.ToUseCase()
	return nil
}

func (r Repo) UpdateFarmInvestContext(payment *entities.Payment) error {
		paymentDb := FromUseCase(payment)

		if err := r.DB.Model(&farm.Farm{}).Where("id = ?", paymentDb.ContextID).
			Update("count_investment", gorm.Expr("count_investment + ?", paymentDb.Amount)).Error; err != nil {
			return err
		}
	
		*payment = *paymentDb.ToUseCase()
		return nil
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
