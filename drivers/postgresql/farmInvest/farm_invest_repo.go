package farmInvest

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
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

	if err := r.DB.Create(&farmInvestDb).Error; err != nil {
		return err
	}

	*farmInvest = *farmInvestDb.ToUseCase()
	return nil
}

func (r *Repo) GetById(farmInvest *entities.FarmInvest) error {
	var farmInvestDb FarmInvest
	if err := r.DB.First(&farmInvestDb, "id = ?", farmInvest.ID).Error; err != nil {
		if r.DB.RowsAffected < 1 {
			return constant.ErrNotFound
		}
		return err
	}
	*farmInvest = *farmInvestDb.ToUseCase()
	return nil
}

func (r *Repo) GetAll(customerID uuid.UUID, farmInvests *[]entities.FarmInvest) error {
	var farmInvestDb []FarmInvest

	if err := r.DB.Where("customer_id = ?", customerID).Find(&farmInvestDb).Error; err != nil {
		return err
	}

	for _, farminvest := range farmInvestDb {
		*farmInvests = append(*farmInvests, *farminvest.ToUseCase())
	}
	return nil
}
