package promo

import (
	"time"

	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PromoStatus string

const (
	Available   PromoStatus = "available"
	Unavailable PromoStatus = "unavailable"
)

type Promo struct {
	ID        uuid.UUID      `gorm:"type:varchar(100)"`
	Name      string         `gorm:"type:varchar(100)"`
	Code      string         `gorm:"type:varchar(100)"`
	Status    PromoStatus    `gorm:"type:promo_status;default:'available'"`
	Amount    float64        `gorm:"type:decimal(10,2)"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromUseCase(promo *entities.Promo) *Promo {
	return &Promo{
		ID:        promo.ID,
		Name:      promo.Name,
		Code:      promo.Code,
		Status:    PromoStatus(promo.Status),
		Amount:    promo.Amount,
		CreatedAt: promo.CreatedAt,
		UpdatedAt: promo.UpdatedAt,
		DeletedAt: promo.DeletedAt,
	}
}

func (u *Promo) ToUseCase() *entities.Promo {
	return &entities.Promo{
		ID:        u.ID,
		Name:      u.Name,
		Code:      u.Code,
		Status:    entities.PromoStatus(u.Status),
		Amount:    u.Amount,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}
