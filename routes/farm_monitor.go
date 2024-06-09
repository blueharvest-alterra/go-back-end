package routes

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/farmmonitor"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/labstack/echo/v4"
)

type FarmMonitorRouteController struct {
	FarmMonitorController *farmmonitor.FarmMonitorController
}

func (r *FarmMonitorRouteController) InitRoute(e *echo.Echo) {
	c := e.Group("/v1/farmmonitors")
	c.Use(middlewares.JWTMiddleware)
	c.POST("", r.FarmMonitorController.Create)
	c.PUT("/:id", r.FarmMonitorController.Update)
	c.GET("", r.FarmMonitorController.GetAll)
	c.GET("/:id", r.FarmMonitorController.GetById)
}
