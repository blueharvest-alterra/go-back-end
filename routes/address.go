package routes

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/address"
	"github.com/labstack/echo/v4"
)

type AddressRouteController struct {
	AddressController *address.AddressController
}

func (r *AddressRouteController) InitRoute(e *echo.Echo) {
	e.GET("/v1/addresses/states", r.AddressController.GetAllStates)
	e.GET("/v1/addresses/cities/:stateID", r.AddressController.GetAllCities)
}
