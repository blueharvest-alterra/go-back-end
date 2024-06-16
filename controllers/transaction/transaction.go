package transaction

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/controllers/transaction/request"
	"github.com/blueharvest-alterra/go-back-end/controllers/transaction/response"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/google/uuid"
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

	transaction, errUseCase := ac.transactionUseCase.Create(transactionCreate.ToEntities(), userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	transactionResponse := response.GetTransactionFromUseCase(&transaction)
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("transaction created", transactionResponse))
}

func (ac *TransactionController) GetByID(c echo.Context) error {
	transactionID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	transaction, errUseCase := ac.transactionUseCase.GetByID(&entities.Transaction{ID: transactionID}, userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	transactionResponse := response.GetTransactionFromUseCase(&transaction)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("get transaction succesful", transactionResponse))
}

func (ac *TransactionController) GetAll(c echo.Context) error {
	var transactionsEntities []entities.Transaction

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	transactions, errUseCase := ac.transactionUseCase.GetAll(&transactionsEntities, userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	transactionResponse := response.SliceFromUseCase(&transactions)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("get transaction succesful", transactionResponse))
}

func (ac *TransactionController) CheckoutSummary(c echo.Context) error {
	var checkoutSummary request.TransactionCreateResponse
	if err := c.Bind(&checkoutSummary); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	transaction, errUseCase := ac.transactionUseCase.CheckoutSummary(checkoutSummary.ToEntities(), userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	transactionResponse := response.GetCheckoutSummaryFromUseCase(&transaction)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("get checkout summary successfully", transactionResponse))
}
