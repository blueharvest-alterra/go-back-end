package transaction

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	transactionUseCase entities.TransactionUseCaseInterface
}

func NewTransactionController(transactionUseCase entities.TransactionUseCaseInterface) *TransactionController {
	return &TransactionController{
		transactionUseCase: transactionUseCase,
	}
}

func (ac *TransactionController) Create(c echo.Context) error {
	return nil
}
