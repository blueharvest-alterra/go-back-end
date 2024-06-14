package routes

import (
	farm_monitor "github.com/blueharvest-alterra/go-back-end/controllers/farm-monitor"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/labstack/echo/v4"
)

type FarmMonitorRouteController struct {
	FarmMonitorController *farm_monitor.FarmMonitorController
}

func (r *FarmMonitorRouteController) InitRoute(e *echo.Echo) {
	c := e.Group("/v1/farmmonitors")
	c.Use(middlewares.JWTMiddleware)
	c.POST("", r.FarmMonitorController.Create)
	c.PUT("/:id", r.FarmMonitorController.Update)
	c.GET("/farm/:farmid", r.FarmMonitorController.GetAllByFarmId)
	c.GET("/:id", r.FarmMonitorController.GetById)
}
