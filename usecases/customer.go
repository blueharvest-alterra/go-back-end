package usecases

import (
	"context"
	"mime/multipart"
	"path/filepath"

	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/utils/google"
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

func (cu *CustomerUseCase) AddAddress(customer *entities.Customer) (entities.Customer, error) {
	if customer.Addresses[0].Address == "" || customer.Addresses[0].Latitude == "" || customer.Addresses[0].ZipCode == "" || customer.Addresses[0].State == "" || customer.Addresses[0].Country == "" || customer.Addresses[0].City == "" {
		return entities.Customer{}, constant.ErrEmptyInput
	}

	customer.Addresses[0].ID = uuid.New()

	if err := cu.repository.AddAddress(customer); err != nil {
		return entities.Customer{}, err
	}

	return *customer, nil
}

func (cu *CustomerUseCase) GetAddresses(customer *entities.Customer) (entities.Customer, error) {
	if err := cu.repository.GetAddresses(customer); err != nil {
		return entities.Customer{}, err
	}

	return *customer, nil
}

func (cu *CustomerUseCase) GetProfile(customer *entities.Customer) (entities.Customer, error) {
	if err := cu.repository.GetProfile(customer); err != nil {
		return entities.Customer{}, err
	}
	return *customer, nil
}

func (cu *CustomerUseCase) EditProfile(customer *entities.Customer, avatar []*multipart.FileHeader) (entities.Customer, error) {
	if customer.FullName == "" || customer.Auth.Email == "" {
		return entities.Customer{}, constant.ErrEmptyInput
	}

	if len(avatar) != 0 {
		file, err := avatar[0].Open()
		if err != nil {
			return entities.Customer{}, err
		}
		defer file.Close()

		ext := filepath.Ext(avatar[0].Filename)

		ctx := context.Background()

		objectName := "avatar" + customer.ID.String() + ext
		url, err := google.Upload.UploadFile(ctx, file, objectName)
		if err != nil {
			return entities.Customer{}, err
		}
		customer.Avatar = url
	}

	if err := cu.repository.EditProfile(customer); err != nil {
		return entities.Customer{}, err
	}

	return *customer, nil
}
