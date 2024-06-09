package routes

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/payment"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/labstack/echo/v4"
)

type PaymentRouteController struct {
	PaymentController *payment.PaymentController
}

func (r *PaymentRouteController) InitRoute(e *echo.Echo) {
	p := e.Group("/v1/payments/callback")
	p.Use(middlewares.CallbackAuth)
	p.POST("/xdt", r.PaymentController.Callback)
}
