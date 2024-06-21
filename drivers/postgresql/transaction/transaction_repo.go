package transaction

import (
	"errors"
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/cart"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repo struct {
	DB *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r Repo) GetAll(transactions *[]entities.Transaction, userData *middlewares.Claims) error {
	var transactionsDb []Transaction

	query := r.DB.Preload("TransactionDetails.Product").Preload(clause.Associations)

	if userData.Role == "customer" {
		query.Where("customer_id = ?", userData.ID)
	}

	if err := query.Find(&transactionsDb).Error; err != nil {
		return err
	}

	for _, _transaction := range transactionsDb {
		*transactions = append(*transactions, *_transaction.ToUseCase())
	}
	return nil
}

func (r Repo) GetByID(transaction *entities.Transaction, userData *middlewares.Claims) error {
	transactionDb := FromUseCase(transaction)

	query := r.DB.Preload("TransactionDetails.Product").Preload(clause.Associations)

	if userData.Role == "customer" {
		query.Where("customer_id = ?", userData.ID)
	}

	if err := query.First(&transactionDb).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	*transaction = *transactionDb.ToUseCase()
	return nil
}

func (r Repo) Create(transaction *entities.Transaction) error {
	transactionDb := FromUseCase(transaction)

	if err := transactionDb.SetCustomerData(r.DB); err != nil {
		return err
	}

	if transactionDb.PromoID != uuid.Nil {
		if err := transactionDb.SetPromoData(r.DB); err != nil {
			return err
		}
	}

	if err := transactionDb.SetAddressData(); err != nil {
		return err
	}

	var productIDs []uuid.UUID

	for i := range transactionDb.TransactionDetails {
		productIDs = append(productIDs, transactionDb.TransactionDetails[i].Product.ID)
		if err := transactionDb.SetTransactionDetail(r.DB, &transactionDb.TransactionDetails[i]); err != nil {
			return err
		}
	}

	transactionDb.Tax = TaxFee
	transactionDb.Total = (transactionDb.SubTotal + transactionDb.Tax + transactionDb.Courier.Fee) - transactionDb.Discount

	transactionDb.PaymentID = uuid.New()
	transactionDb.Payment.ID = transactionDb.PaymentID
	transactionDb.Payment.Amount = transactionDb.Total
	transactionDb.Payment.Status = "UNPAID"

	if err := transactionDb.Payment.Create("buy_product", transactionDb.ID.String()); err != nil {
		return err
	}

	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&transactionDb).Error; err != nil {
		tx.Rollback()
		return constant.ErrCreateTransaction
	}

	for _, productID := range productIDs {
		query := tx.Where("customer_id = ?", transactionDb.Customer.ID).Where("product_id = ?", productID).Delete(&cart.Cart{})
		if err := query.Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	*transaction = *transactionDb.ToUseCase()
	return nil
}

func (r Repo) CheckoutSummary(transaction *entities.Transaction, userData *middlewares.Claims) error {
	transactionDb := FromUseCase(transaction)

	if err := transactionDb.SetCustomerData(r.DB); err != nil {
		return err
	}

	if transactionDb.PromoID != uuid.Nil {
		if err := transactionDb.SetPromoData(r.DB); err != nil {
			return err
		}
	}

	if err := transactionDb.SetAddressData(); err != nil {
		return err
	}

	for i := range transactionDb.TransactionDetails {
		if err := transactionDb.SetTransactionDetail(r.DB, &transactionDb.TransactionDetails[i]); err != nil {
			return err
		}
	}

	transactionDb.Tax = TaxFee
	transactionDb.Total = (transactionDb.SubTotal + transactionDb.Tax + transactionDb.Courier.Fee) - transactionDb.Discount

	*transaction = *transactionDb.ToUseCase()
	return nil
}

//func UpdateBuyProductContext(paymentID uuid.UUID) error {
//	if err := Repo.DB.Preload("TransactionDetail").Where("payment_id = ?", paymentID).First(&transactionDb).Error; err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return constant.ErrNotFound
//		}
//		return err
//	}
//	return nil
//}
