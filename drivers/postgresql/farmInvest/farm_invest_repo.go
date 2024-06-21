package farmInvest

import (
	"errors"
	"fmt"

	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repo struct {
	DB *gorm.DB
}

func NewFarmInvestRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Create(farmInvest *entities.FarmInvest) error {
	var farm entities.Farm
	if err := r.DB.First(&farm, "id = ?", farmInvest.FarmID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	// Check minimum invest amount validation
	if farmInvest.InvestmentAmount < farm.MinimumInvestmentAmount {
		return constant.ErrMinumumAmount
	}

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

	if err := farmInvestDb.Payment.Create("farm_invest", farmInvest.FarmID.String()); err != nil {
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
	farmInvestDb := FromUseCase(farmInvest)

	query := r.DB.Preload("Farm").Preload(clause.Associations)

	if userData.Role == "customer" {
		query.Where("customer_id = ?", userData.ID)
	}

	if err := query.First(&farmInvestDb).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrNotFound
		}
		return err
	}

	*farmInvest = *farmInvestDb.ToUseCase()
	return nil
}

func (r *Repo) GetAll(farmInvests *[]entities.FarmInvest, userData *middlewares.Claims) error {
	var farmInvestDb []FarmInvest

	if err := r.DB.Preload("Payment").Where("customer_id = ?", userData.ID).Find(&farmInvestDb).Error; err != nil {
		return err
	}

	fmt.Println("hit: ", farmInvestDb)

	for _, farminvest := range farmInvestDb {
		*farmInvests = append(*farmInvests, *farminvest.ToUseCase())
	}
	return nil
}
