package farmMonitor

import (
	"fmt"

	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewFarmMonitorRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Create(farmInvest *entities.FarmMonitor) error {
	farmInvestDb := FromUseCase(farmInvest)

	if err := r.DB.Create(&farmInvestDb).Error; err != nil {
		return err
	}

	*farmInvest = *farmInvestDb.ToUseCase()
	return nil
}

func (r *Repo) GetById(farmMonitor *entities.FarmMonitor) error {
	var farmMonitorDb FarmMonitor
	if err := r.DB.First(&farmMonitorDb, "id = ?", farmMonitor.ID).Error; err != nil {
		if r.DB.RowsAffected < 1 {
			return constant.ErrNotFound
		}
		return err
	}

	*farmMonitor = *farmMonitorDb.ToUseCase()
	return nil
}

func (r *Repo) GetAllByFarmId(farmID uuid.UUID, farmMonitors *[]entities.FarmMonitor) error {
	var farmMonitorDb []FarmMonitor
	if err := r.DB.Where("farm_id = ?", farmID).Find(&farmMonitorDb).Error; err != nil {
		return err
	}

	a := r.DB.Where("farm_id = ?", farmID).Find(&farmMonitorDb)

	fmt.Println("hit", a)

	for _, farm := range farmMonitorDb {
		*farmMonitors = append(*farmMonitors, *farm.ToUseCase())
	}
	return nil
}

func (r *Repo) Update(farmMonitor *entities.FarmMonitor) error {
	farmMonitorDb := FromUseCase(farmMonitor)

	db := r.DB.Where("id = ?", farmMonitorDb.ID).Updates(&farmMonitorDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}
	if err := db.Error; err != nil {
		return err
	}

	*farmMonitor = *farmMonitorDb.ToUseCase()
	return nil
}
