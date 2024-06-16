package routes

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/transaction"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/labstack/echo/v4"
)

type TransactionRouteController struct {
	TransactionController *transaction.TransactionController
}

func (r *TransactionRouteController) InitRoute(e *echo.Echo) {
	t := e.Group("/v1/transactions")
	t.Use(middlewares.JWTMiddleware)
	t.POST("", r.TransactionController.Create)
	t.GET("/:id", r.TransactionController.GetByID)
	t.GET("", r.TransactionController.GetAll)
	t.POST("/checkout/summaries", r.TransactionController.CheckoutSummary)
}
