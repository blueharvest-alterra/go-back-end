package usecases

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type PromoUseCase struct {
	repository entities.PromoRepositoryInterface
}

func NewPromoUseCase(repository entities.PromoRepositoryInterface) *PromoUseCase {
	return &PromoUseCase{repository: repository}
}

func (c *PromoUseCase) Create(promo *entities.Promo) (entities.Promo, error) {
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

func (c *PromoUseCase) Update(promo *entities.Promo) (entities.Promo, error) {
	if err := c.repository.Update(promo); err != nil {
		return entities.Promo{}, err
	}

	return *promo, nil
}

func (c *PromoUseCase) Delete(id uuid.UUID) (entities.Promo, error) {
	var promo entities.Promo
	promo.ID = id

	if err := c.repository.Delete(&promo); err != nil {
		return entities.Promo{}, err
	}

	return promo, nil
}

func (c *PromoUseCase) GetAll(promo *[]entities.Promo) ([]entities.Promo, error) {
	if err := c.repository.GetAll(promo); err != nil {
		return []entities.Promo{}, err
	}

	return *promo, nil
}
