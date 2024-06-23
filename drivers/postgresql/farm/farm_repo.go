package farm

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewFarmRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Create(farm *entities.Farm) error {
	farmDb := FromUseCase(farm)

	if err := r.DB.Create(&farmDb).Error; err != nil {
		return err
	}

	*farm = *farmDb.ToUseCase()
	return nil
}

func (r *Repo) Update(farm *entities.Farm) error {
	farmDb := FromUseCase(farm)

	db := r.DB.Where("id = ?", farmDb.ID).Updates(&farmDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}

	*farm = *farmDb.ToUseCase()
	return nil
}

func (r *Repo) Delete(farm *entities.Farm) error {
	farmDb := FromUseCase(farm)

	db := r.DB.Delete(&farmDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}

	if err := db.Error; err != nil {
		return err
	}

	*farm = *farmDb.ToUseCase()
	return nil
}

func (r *Repo) GetById(farm *entities.Farm) error {
	var farmDb Farm
	if err := r.DB.First(&farmDb, "id = ?", farm.ID).Error; err != nil {
		if r.DB.RowsAffected < 1 {
			return constant.ErrNotFound
		}
		return err
	}
	*farm = *farmDb.ToUseCase()
	return nil
}

func (r *Repo) GetAll(farms *[]entities.Farm) error {
	var farmDb []Farm

	if err := r.DB.Find(&farmDb).Error; err != nil {
		return err
	}

	for _, farm := range farmDb {
		*farms = append(*farms, *farm.ToUseCase())
	}
	return nil
}
