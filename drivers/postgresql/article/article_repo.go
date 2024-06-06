package article

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewArticleRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Create(article *entities.Article) error {
	articleDb := FromUseCase(article)

	if err := r.DB.Create(&articleDb).Error; err != nil {
		return err
	}

	*article = *articleDb.ToUseCase()
	return nil
}

func (r *Repo) Update(article *entities.Article) error {
	articleDb := FromUseCase(article)

	db := r.DB.Where("id = ?", articleDb.ID).Updates(&articleDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}
	if err := db.Error; err != nil {
		return err
	}

	*article = *articleDb.ToUseCase()
	return nil
}

func (r *Repo) Delete(article *entities.Article) error {
	articleDb := FromUseCase(article)

	db := r.DB.Delete(&articleDb)
	if db.RowsAffected < 1 {
		return constant.ErrNotFound
	}
	if err := db.Error; err != nil {
		return err
	}

	*article = *articleDb.ToUseCase()
	return nil
}

func (r *Repo) GetById(article *entities.Article) error {
	var articleDb Article
	if err := r.DB.First(&articleDb, "id = ?", article.ID).Error; err != nil {
		if r.DB.RowsAffected < 1 {
			return constant.ErrNotFound
		}
		return err
	}

	*article = *articleDb.ToUseCase()
	return nil
}

func (r *Repo) GetAll(articles *[]entities.Article) error {
	var articleDb []Article

	if err := r.DB.Find(&articleDb).Error; err != nil {
		return err
	}

	for _, article := range articleDb {
		*articles = append(*articles, *article.ToUseCase())
	}
	return nil
}
