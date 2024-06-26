package routes

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/farm-invest"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/labstack/echo/v4"
)

type FarmInvestRouteController struct {
	FarmInvestController *farm_invest.FarmInvestController
}

func (r *FarmInvestRouteController) InitRoute(e *echo.Echo) {
	c := e.Group("/v1/farminvests")
	c.Use(middlewares.JWTMiddleware)
	c.POST("", r.FarmInvestController.Create)
	c.GET("", r.FarmInvestController.GetAll)
	c.GET("/:id", r.FarmInvestController.GetById)
}
