package entities

import (
	"mime/multipart"
	"time"

	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	ID        uuid.UUID
	Title     string
	Content   string
	Picture   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type ArticleRepositoryInterface interface {
	Create(article *Article) error
	GetById(article *Article) error
	Update(article *Article) error
	Delete(article *Article) error
	GetAll(articles *[]Article) error
}

type ArticleUseCaseInterface interface {
	Create(article *Article, userData *middlewares.Claims, picture []*multipart.FileHeader) (Article, error)
	GetById(id uuid.UUID) (Article, error)
	Update(article *Article, userData *middlewares.Claims, picture []*multipart.FileHeader) (Article, error)
	Delete(id uuid.UUID, userData *middlewares.Claims) (Article, error)
	GetAll(articles *[]Article) ([]Article, error)
}
