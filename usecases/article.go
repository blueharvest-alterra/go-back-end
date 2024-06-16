package usecases

import (
	"context"
	"mime/multipart"
	"path/filepath"

	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/blueharvest-alterra/go-back-end/utils/google"
	"github.com/google/uuid"
)

type ArticleUseCase struct {
	repository entities.ArticleRepositoryInterface
}

func NewArticleUseCase(repository entities.ArticleRepositoryInterface) *ArticleUseCase {
	return &ArticleUseCase{repository: repository}
}

func (c *ArticleUseCase) Create(article *entities.Article, userData *middlewares.Claims, picture []*multipart.FileHeader) (entities.Article, error) {
	if userData.Role != "admin" {
		return entities.Article{}, constant.ErrNotAuthorized
	}

	if article.Title == "" || article.Content == "" {
		return entities.Article{}, constant.ErrEmptyInput
	}

	article.ID = uuid.New()

	file, err := picture[0].Open()
	if err != nil {
		return entities.Article{}, err
	}
	defer file.Close()

	ext := filepath.Ext(picture[0].Filename)

	ctx := context.Background()

	objectName := article.ID.String() + ext
	url, err := google.Upload.UploadFile(ctx, file, objectName)
	if err != nil {
		return entities.Article{}, err
	}

	article.Picture = url

	if err := c.repository.Create(article); err != nil {
		return entities.Article{}, err
	}

	return *article, nil
}

func (c *ArticleUseCase) GetById(id uuid.UUID) (entities.Article, error) {
	var article entities.Article
	article.ID = id

	if err := c.repository.GetById(&article); err != nil {
		return entities.Article{}, err
	}

	return article, nil
}

func (c *ArticleUseCase) Update(article *entities.Article, userData *middlewares.Claims, picture []*multipart.FileHeader) (entities.Article, error) {
	if userData.Role != "admin" {
		return entities.Article{}, constant.ErrNotAuthorized
	}

	if article.Title == "" || article.Content == "" {
		return entities.Article{}, constant.ErrEmptyInput
	}
	if len(picture) != 0 {
		file, err := picture[0].Open()
		if err != nil {
			return entities.Article{}, err
		}
		defer file.Close()

		ext := filepath.Ext(picture[0].Filename)

		ctx := context.Background()

		objectName := article.ID.String() + ext
		url, err := google.Upload.UploadFile(ctx, file, objectName)
		if err != nil {
			return entities.Article{}, err
		}
		article.Picture = url
	}

	if err := c.repository.Update(article); err != nil {
		return entities.Article{}, err
	}

	return *article, nil
}

func (c *ArticleUseCase) Delete(id uuid.UUID, userData *middlewares.Claims) (entities.Article, error) {
	if userData.Role != "admin" {
		return entities.Article{}, constant.ErrNotAuthorized
	}

	ctx := context.Background()

	var article entities.Article
	article.ID = id

	if err := c.repository.GetById(&article); err != nil {
		return entities.Article{}, err
	}

	objectName := filepath.Base(article.Picture)
	err := google.Upload.DeleteFileFromGCS(ctx, objectName)
	if err != nil {
		return entities.Article{}, err
	}

	if err := c.repository.Delete(&article); err != nil {
		return entities.Article{}, err
	}

	return article, nil
}

func (c *ArticleUseCase) GetAll(article *[]entities.Article) ([]entities.Article, error) {
	if err := c.repository.GetAll(article); err != nil {
		return []entities.Article{}, err
	}

	return *article, nil
}
