package routes

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/customer"
	"github.com/labstack/echo/v4"
)

type CustomerRouteController struct {
	CustomerController *customer.CustomerController
}

func (r *CustomerRouteController) InitRoute(e *echo.Echo) {
	e.POST("/v1/login/customer", r.CustomerController.Login)
	e.POST("/v1/register/customer", r.CustomerController.Register)
}
