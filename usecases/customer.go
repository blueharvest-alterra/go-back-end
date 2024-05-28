package usecases

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CustomerUseCase struct {
	repository entities.CustomerRepositoryInterface
}

func NewCustomerUseCase(repository entities.CustomerRepositoryInterface) *CustomerUseCase {
	return &CustomerUseCase{
		repository: repository,
	}
}

func (cu *CustomerUseCase) Login(customer *entities.Customer) (entities.Customer, error) {
	if customer.Auth.Email == "" || customer.Auth.Password == "" {
		return entities.Customer{}, constant.ErrEmptyInput
	}

	password := customer.Auth.Password

	if err := cu.repository.Login(customer); err != nil {
		return entities.Customer{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Auth.Password), []byte(password)); err != nil {
		return entities.Customer{}, constant.ErrInvalidEmailOrPassword
	}

	return *customer, nil
}

func (cu *CustomerUseCase) Register(customer *entities.Customer) (entities.Customer, error) {
	if customer.FullName == "" || customer.Auth.Email == "" || customer.Auth.Password == "" {
		return entities.Customer{}, constant.ErrEmptyInput
	}

	customer.ID = uuid.New()
	customer.Auth.ID = uuid.New()

	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(customer.Auth.Password), bcrypt.DefaultCost)
	if errHash != nil {
		return entities.Customer{}, constant.ErrInvalidRequest
	}

	customer.Auth.Password = string(hashedPassword)

	if err := cu.repository.Register(customer); err != nil {
		return entities.Customer{}, err
	}

	return *customer, nil
}
