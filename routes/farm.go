package routes

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/farm"
	"github.com/labstack/echo/v4"
)

type FarmRouteController struct {
	FarmController *farm.FarmController
}

func (r *FarmRouteController) InitRoute(e *echo.Echo) {
	c := e.Group("/v1/farms")
	c.POST("", r.FarmController.Create)
	c.PUT("/:id", r.FarmController.Update)
	c.DELETE("/:id", r.FarmController.Delete)
	c.GET("", r.FarmController.GetAll)
	c.GET("/:id", r.FarmController.GetById)
}
