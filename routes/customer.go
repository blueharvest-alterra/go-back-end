package routes

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/customer"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/labstack/echo/v4"
)

type CustomerRouteController struct {
	CustomerController *customer.CustomerController
}

func (r *CustomerRouteController) InitRoute(e *echo.Echo) {
	e.POST("/v1/login/customer", r.CustomerController.Login)
	e.POST("/v1/register/customer", r.CustomerController.Register)

	c := e.Group("/v1/customers")
	c.Use(middlewares.JWTMiddleware)
	c.POST("/addresses", r.CustomerController.CreateAddress)
	c.GET("/addresses", r.CustomerController.GetAddresses)

	c.GET("/profile", r.CustomerController.GetProfile)
	c.POST("/profile", r.CustomerController.EditProfile)

}
