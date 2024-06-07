package routes

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/courier"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/labstack/echo/v4"
)

type CourierRouteController struct {
	CourierController *courier.CourierController
}

func (r *CourierRouteController) InitRoute(e *echo.Echo) {
	p := e.Group("/v1/couriers")
	p.Use(middlewares.JWTMiddleware)
	p.POST("", r.CourierController.GetAll)
}
