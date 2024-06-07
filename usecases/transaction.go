package usecases

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
)

type TransactionUseCase struct {
	repository entities.TransactionRepositoryInterface
}

func NewTransactionUseCase(repository entities.TransactionRepositoryInterface) *TransactionUseCase {
	return &TransactionUseCase{repository: repository}
}

func (t TransactionUseCase) Create(transaction *entities.Transaction, userData *middlewares.Claims) (entities.Transaction, error) {
	//TODO implement me
	panic("implement me")
}
