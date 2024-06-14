package usecases

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
)

type CartUseCase struct {
	repository entities.CartRepositoryInterface
}

func NewCartUseCase(repository entities.CartRepositoryInterface) *CartUseCase {
	return &CartUseCase{repository: repository}
}

func (c *CartUseCase) Create(cart *entities.Cart) (entities.Cart, error) {
	cart.ID = uuid.New()

	if err := c.repository.Create(cart); err != nil {
		return entities.Cart{}, err
	}

	return *cart, nil
}

func (c *CartUseCase) Update(cart *entities.Cart) (entities.Cart, error) {
	if err := c.repository.Update(cart); err != nil {
		return entities.Cart{}, err
	}

	return *cart, nil
}

func (c *CartUseCase) Delete(id uuid.UUID) (entities.Cart, error) {
	var cart entities.Cart
	cart.ID = id
	if err := c.repository.Delete(&cart); err != nil {
		return entities.Cart{}, err
	}

	return cart, nil
}

func (c *CartUseCase) GetById(id uuid.UUID) (entities.Cart, error) {
	var cart entities.Cart
	cart.ID = id

	if err := c.repository.GetById(&cart); err != nil {
		return entities.Cart{}, err
	}

	return cart, nil
}

func (c *CartUseCase) GetAll(carts *[]entities.Cart, userData *middlewares.Claims) ([]entities.Cart, error) {
	if err := c.repository.GetAll(carts, userData); err != nil {
		return []entities.Cart{}, err
	}

	return *carts, nil
}
