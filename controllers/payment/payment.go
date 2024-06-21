package payment

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/controllers/payment/request"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	paymentUseCase entities.PaymentUseCaseInterface
}

func NewPaymentController(paymentUseCase entities.PaymentUseCaseInterface) *PaymentController {
	return &PaymentController{
		paymentUseCase: paymentUseCase,
	}
}

func (ac *PaymentController) Callback(c echo.Context) error {
	var PaymentCallback request.PaymentCallbackRequest
	if err := c.Bind(&PaymentCallback); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	payment, errUseCase := ac.paymentUseCase.UpdateStatus(PaymentCallback.ToEntities())
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	_, errUseCaseUpdatePaymentContext := ac.paymentUseCase.UpdatePaymentContext(&payment)
	if errUseCaseUpdatePaymentContext != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	return nil
}
