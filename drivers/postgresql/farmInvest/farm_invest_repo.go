package farmInvest

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewFarmInvestRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Create(farmInvest *entities.FarmInvest) error {
	farmInvestDb := FromUseCase(farmInvest)

	tx := r.DB.Begin()
	defer func() {
		if re := recover(); re != nil {
			tx.Rollback()
		}
	}()

	// Create Payment
	farmInvestDb.PaymentID = uuid.New()
	farmInvestDb.Payment.ID = farmInvestDb.PaymentID
	farmInvestDb.Payment.Amount = farmInvestDb.InvestmentAmount
	farmInvestDb.Payment.Status = "UNPAID"

	if err := farmInvestDb.Payment.Create(); err != nil {
		return err
	}

	// Create FarmInvest
	if err := tx.Create(&farmInvestDb).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	*farmInvest = *farmInvestDb.ToUseCase()
	return nil
}

func (r *Repo) GetById(farmInvest *entities.FarmInvest, userData *middlewares.Claims) error {
	var farmInvestDb FarmInvest

	if err := r.DB.First(&farmInvest, "id = ?", farmInvest.ID).Error; err != nil {
		if r.DB.RowsAffected < 1 {
			return constant.ErrNotFound
		}
		return err
	}
	*farmInvest = *farmInvestDb.ToUseCase()
	return nil
}

func (r *Repo) GetAll(farmInvests *[]entities.FarmInvest, userData *middlewares.Claims) error {
	var farmInvestDb []FarmInvest

	if err := r.DB.Where("customer_id = ?", userData.ID).Find(&farmInvestDb).Error; err != nil {
		return err
	}

	for _, farminvest := range farmInvestDb {
		*farmInvests = append(*farmInvests, *farminvest.ToUseCase())
	}
	return nil
}
