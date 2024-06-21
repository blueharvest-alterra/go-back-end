package article

import (
	"time"

	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/admin"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	ID        uuid.UUID `gorm:"type:varchar(100)"`
	AdminID   uuid.UUID `gorm:"type:varchar(100)"`
	Admin     admin.Admin
	Title     string         `gorm:"type:varchar(100)"`
	Content   string         `gorm:"type:varchar(65535)"`
	Picture   string         `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromUseCase(article *entities.Article) *Article {
	return &Article{
		ID:      article.ID,
		AdminID: article.AdminID,
		Admin: admin.Admin{
			ID:       article.Admin.ID,
			FullName: article.Admin.FullName,
		},
		Title:     article.Title,
		Content:   article.Content,
		Picture:   article.Picture,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		DeletedAt: article.DeletedAt,
	}
}

func (u *Article) ToUseCase() *entities.Article {
	return &entities.Article{
		ID:      u.ID,
		AdminID: u.AdminID,
		Admin: entities.Admin{
			ID:       u.Admin.ID,
			FullName: u.Admin.FullName,
		},
		Title:     u.Title,
		Content:   u.Content,
		Picture:   u.Picture,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}
