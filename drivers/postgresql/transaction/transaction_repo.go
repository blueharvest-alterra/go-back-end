package transaction

import (
	"errors"
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/product"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/promo"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r Repo) Create(transaction *entities.Transaction) error {
	transactionDb := FromUseCase(transaction)

	tx := r.DB.Begin()
	defer func() {
		if re := recover(); re != nil {
			tx.Rollback()
		}
	}()

	promoData := promo.Promo{ID: transactionDb.PromoID}
	if promoData.ID != uuid.Nil {
		if err := tx.First(&promoData).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return constant.ErrNotFound
			}
			return err
		}

		transactionDb.Discount = promoData.Amount
	}

	var subTotal float64

	for i := range transactionDb.TransactionDetails {
		transactionItem := &transactionDb.TransactionDetails[i]
		if err := tx.Where(&product.Product{ID: transactionItem.ProductID}).First(&transactionItem.Product).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return constant.ErrNotFound
			}
			return err
		}
		transactionItem.TotalPrice = transactionItem.Product.Price * float64(transactionItem.Quantity)
		subTotal += transactionItem.TotalPrice
	}

	transactionDb.SubTotal = subTotal
	transactionDb.Tax = TaxFee
	transactionDb.Total = (transactionDb.SubTotal + transactionDb.Tax + transactionDb.Courier.Fee) - transactionDb.Discount

	transactionDb.PaymentID = uuid.New()
	transactionDb.Payment.ID = transactionDb.PaymentID
	transactionDb.Payment.Amount = transactionDb.Total
	transactionDb.Payment.Status = "UNPAID"

	if err := transactionDb.Payment.Create(); err != nil {
		return err
	}

	if err := tx.Create(&transactionDb).Error; err != nil {
		panic(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
	}

	*transaction = *transactionDb.ToUseCase()
	return nil
}
