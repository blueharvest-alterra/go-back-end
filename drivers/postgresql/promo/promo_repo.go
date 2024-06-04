package promo

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewPromoRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Create(promo *entities.Promo) error {
	promoDb := FromUseCase(promo)

	if err := r.DB.Create(&promoDb).Error; err != nil {
		return err
	}

	*promo = *promoDb.ToUseCase()
	return nil
}

func (r *Repo) Update(promo *entities.Promo) error {
	promoDb := FromUseCase(promo)

	db := r.DB.Where("id = ?", promoDb.ID).Updates(&promoDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}
	if err := db.Error; err != nil {
		return err
	}

	*promo = *promoDb.ToUseCase()
	return nil
}

func (r *Repo) Delete(promo *entities.Promo) error {
	promoDb := FromUseCase(promo)

	db := r.DB.Delete(&promoDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}
	if err := db.Error; err != nil {
		return err
	}

	*promo = *promoDb.ToUseCase()
	return nil
}

func (r *Repo) GetById(promo *entities.Promo) error {
	var promoDb Promo
	if err := r.DB.First(&promoDb, "id = ?", promo.ID).Error; err != nil {
		if r.DB.RowsAffected < 1 {
			return constant.ErrNotFound
		}
		return err
	}
	
	*promo = *promoDb.ToUseCase()
	return nil
}

func (r *Repo) GetAll(promos *[]entities.Promo) error {
	var promoDb []Promo

	if err := r.DB.Find(&promoDb).Error; err != nil {
		return err
	}

	for _, promo := range promoDb {
		*promos = append(*promos, *promo.ToUseCase())
	}
	return nil
}
