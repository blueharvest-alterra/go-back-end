package transaction

import (
	"fmt"
	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/controllers/transaction/request"
	"github.com/blueharvest-alterra/go-back-end/controllers/transaction/response"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/labstack/echo/v4"
	"net/http"
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
	var transactionCreate request.TransactionCreateResponse
	if err := c.Bind(&transactionCreate); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	fmt.Println("cntrlrl", utils.PrettyPrint(transactionCreate))

	transaction, errUseCase := ac.transactionUseCase.Create(transactionCreate.ToEntities(), userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	transactionResponse := response.GetTransactionFromUseCase(&transaction)
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("transaction created", transactionResponse))
}

func (ac *TransactionController) PaymentCallback(c echo.Context) error {
	return nil
}
