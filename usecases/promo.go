package usecases

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
)

type PromoUseCase struct {
	repository entities.PromoRepositoryInterface
}

func NewPromoUseCase(repository entities.PromoRepositoryInterface) *PromoUseCase {
	return &PromoUseCase{repository: repository}
}

func (c *PromoUseCase) Create(promo *entities.Promo, userData *middlewares.Claims) (entities.Promo, error) {
	if userData.Role != "admin" {
		return entities.Promo{}, constant.ErrNotAuthorized
	}

	if promo.Name == "" || promo.Code == "" || promo.Status == "" || promo.Amount == 0 {
		return entities.Promo{}, constant.ErrEmptyInput
	}
	promo.ID = uuid.New()

	if err := c.repository.Create(promo); err != nil {
		return entities.Promo{}, err
	}

	return *promo, nil
}

func (c *PromoUseCase) GetById(id uuid.UUID) (entities.Promo, error) {
	var promo entities.Promo
	promo.ID = id

	if err := c.repository.GetById(&promo); err != nil {
		return entities.Promo{}, err
	}

	return promo, nil
}

func (c *PromoUseCase) Update(promo *entities.Promo, userData *middlewares.Claims) (entities.Promo, error) {
	if userData.Role != "admin" {
		return entities.Promo{}, constant.ErrNotAuthorized
	}

	if promo.Name == "" || promo.Code == "" || promo.Status == "" || promo.Amount == 0 {
		return entities.Promo{}, constant.ErrEmptyInput
	}
	if err := c.repository.Update(promo); err != nil {
		return entities.Promo{}, err
	}

	return *promo, nil
}

func (c *PromoUseCase) Delete(id uuid.UUID, userData *middlewares.Claims) (entities.Promo, error) {
	if userData.Role != "admin" {
		return entities.Promo{}, constant.ErrNotAuthorized
	}

	var promo entities.Promo
	promo.ID = id

	if err := c.repository.Delete(&promo); err != nil {
		return entities.Promo{}, err
	}

	return promo, nil
}

func (c *PromoUseCase) GetAll(promo *[]entities.Promo, userData *middlewares.Claims) ([]entities.Promo, error) {
	if userData.Role != "admin" {
		return []entities.Promo{}, constant.ErrNotAuthorized
	}

	if err := c.repository.GetAll(promo); err != nil {
		return []entities.Promo{}, err
	}

	return *promo, nil
}
