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

type FarmUseCase struct {
	repository entities.FarmRepositoryInterface
}

func NewFarmUseCase(repository entities.FarmRepositoryInterface) *FarmUseCase {
	return &FarmUseCase{repository: repository}
}

func (c *FarmUseCase) Create(farm *entities.Farm, userData *middlewares.Claims, picture []*multipart.FileHeader) (entities.Farm, error) {
	if userData.Role != "admin" {
		return entities.Farm{}, constant.ErrNotAuthorized
	}

	if farm.Title == "" || farm.Description == "" {
		return entities.Farm{}, constant.ErrEmptyInput
	}
	farm.ID = uuid.New()

	file, err := picture[0].Open()
	if err != nil {
		return entities.Farm{}, err
	}
	defer file.Close()

	ext := filepath.Ext(picture[0].Filename)

	ctx := context.Background()

	objectName := farm.ID.String() + ext
	url, err := google.Upload.UploadFile(ctx, file, objectName)
	if err != nil {
		return entities.Farm{}, err
	}

	farm.Picture = url

	if err := c.repository.Create(farm); err != nil {
		return entities.Farm{}, err
	}

	return *farm, nil
}

func (c *FarmUseCase) GetById(id uuid.UUID) (entities.Farm, error) {
	var farm entities.Farm
	farm.ID = id

	if err := c.repository.GetById(&farm); err != nil {
		return entities.Farm{}, err
	}

	return farm, nil
}

func (c *FarmUseCase) Update(farm *entities.Farm, userData *middlewares.Claims, picture []*multipart.FileHeader) (entities.Farm, error) {
	if userData.Role != "admin" {
		return entities.Farm{}, constant.ErrNotAuthorized
	}

	if farm.Title == "" || farm.Description == "" {
		return entities.Farm{}, constant.ErrEmptyInput
	}

	if len(picture) != 0 {
		file, err := picture[0].Open()
		if err != nil {
			return entities.Farm{}, err
		}
		defer file.Close()

		ext := filepath.Ext(picture[0].Filename)

		ctx := context.Background()

		objectName := farm.ID.String() + ext
		url, err := google.Upload.UploadFile(ctx, file, objectName)
		if err != nil {
			return entities.Farm{}, err
		}
		farm.Picture = url
	}

	if err := c.repository.Update(farm); err != nil {
		return entities.Farm{}, err
	}

	return *farm, nil
}

func (c *FarmUseCase) Delete(id uuid.UUID, userData *middlewares.Claims) (entities.Farm, error) {
	if userData.Role != "admin" {
		return entities.Farm{}, constant.ErrNotAuthorized
	}
	var farm entities.Farm
	farm.ID = id

	if err := c.repository.Delete(&farm); err != nil {
		return entities.Farm{}, err
	}

	return farm, nil
}

func (c *FarmUseCase) GetAll(farm *[]entities.Farm) ([]entities.Farm, error) {
	if err := c.repository.GetAll(farm); err != nil {
		return []entities.Farm{}, err
	}

	return *farm, nil
}
