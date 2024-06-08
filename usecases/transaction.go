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

func NewTransactionUseCase(repository entities.TransactionRepositoryInterface) *TransactionUseCase {
	return &TransactionUseCase{repository: repository}
}

func (t TransactionUseCase) Create(transaction *entities.Transaction, userData *middlewares.Claims) (entities.Transaction, error) {
	if userData.Role != "customer" {
		return entities.Transaction{}, constant.ErrNotAuthorized
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
