package cart

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repo struct {
	DB *gorm.DB
}

func NewCartRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Create(cart *entities.Cart) error {
	cartDb := FromUseCase(cart)

	if err := r.DB.Create(&cartDb).Error; err != nil {
		return err
	}

	*cart = *cartDb.ToUseCase()
	return nil
}

func (r *Repo) GetById(cart *entities.Cart) error {
	var cartDb Cart
	if err := r.DB.First(&cartDb, "id = ?", cart.ID).Error; err != nil {
		if r.DB.RowsAffected < 1 {
			return constant.ErrNotFound
		}
		return err
	}
	*cart = *cartDb.ToUseCase()
	return nil
}

func (r *Repo) Update(cart *entities.Cart) error {
	cartDb := FromUseCase(cart)

	db := r.DB.Where("id = ?", cartDb.ID).Updates(&cartDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}
	if err := db.Error; err != nil {
		return err
	}

	*cart = *cartDb.ToUseCase()
	return nil
}

func (r *Repo) Delete(cart *entities.Cart) error {
	cartDb := FromUseCase(cart)

	db := r.DB.Delete(&cartDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}
	if err := db.Error; err != nil {
		return err
	}

	*cart = *cartDb.ToUseCase()
	return nil
}

func (r *Repo) GetAll(carts *[]entities.Cart, userData *middlewares.Claims) error {
	var cartDb []Cart

	if err := r.DB.Preload(clause.Associations).Where("customer_id = ?", userData.ID).Find(&cartDb).Error; err != nil {
		return err
	}

	for _, cart := range cartDb {
		*carts = append(*carts, *cart.ToUseCase())
	}
	return nil
}
