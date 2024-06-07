package transaction

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r Repo) Create(transaction *entities.Transaction) error {
	//TODO implement me
	panic("implement me")
}
