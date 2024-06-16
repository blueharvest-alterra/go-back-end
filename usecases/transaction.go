package usecases

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
)

type TransactionUseCase struct {
	repository entities.TransactionRepositoryInterface
}

func (t TransactionUseCase) CheckoutSummary(transaction *entities.Transaction, userData *middlewares.Claims) (entities.Transaction, error) {
	if userData.Role != "customer" {
		return entities.Transaction{}, constant.ErrNotAuthorized
	}

	if len(transaction.TransactionDetails) < 1 || transaction.Courier.DestinationAddressID == uuid.Nil || transaction.Courier.Name == "" || transaction.Courier.Type == "" || transaction.Courier.Fee < 0 {
		return entities.Transaction{}, constant.ErrEmptyInput
	}

	transaction.Customer.ID = userData.ID

	if err := t.repository.CheckoutSummary(transaction, userData); err != nil {
		return entities.Transaction{}, err
	}

	return *transaction, nil
}

func (t TransactionUseCase) GetAll(transactions *[]entities.Transaction, userData *middlewares.Claims) ([]entities.Transaction, error) {

	if err := t.repository.GetAll(transactions, userData); err != nil {
		return []entities.Transaction{}, err
	}

	return *transactions, nil
}

func (t TransactionUseCase) GetByID(transaction *entities.Transaction, userData *middlewares.Claims) (entities.Transaction, error) {
	if transaction.ID == uuid.Nil {
		return entities.Transaction{}, constant.ErrEmptyInput
	}

	if userData.Role == "customer" {
		transaction.CustomerID = userData.ID
	}

	if err := t.repository.GetByID(transaction, userData); err != nil {
		return entities.Transaction{}, err
	}

	return *transaction, nil
}

func NewTransactionUseCase(repository entities.TransactionRepositoryInterface) *TransactionUseCase {
	return &TransactionUseCase{repository: repository}
}

func (t TransactionUseCase) Create(transaction *entities.Transaction, userData *middlewares.Claims) (entities.Transaction, error) {
	if userData.Role != "customer" {
		return entities.Transaction{}, constant.ErrNotAuthorized
	}

	if len(transaction.TransactionDetails) < 1 || transaction.Courier.DestinationAddressID == uuid.Nil || transaction.Courier.Name == "" || transaction.Courier.Type == "" || transaction.Courier.Fee < 0 {
		return entities.Transaction{}, constant.ErrEmptyInput
	}

	transaction.ID = uuid.New()
	transaction.Customer.ID = userData.ID
	transaction.Customer.FullName = userData.FullName
	transaction.Customer.PhoneNumber = ""
	transaction.Courier.ID = uuid.New()

	if err := t.repository.Create(transaction); err != nil {
		return entities.Transaction{}, err
	}

	return *transaction, nil
}
