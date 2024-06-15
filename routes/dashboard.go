package routes

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/dashboard"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/labstack/echo/v4"
)

type DashboardRouteController struct {
	DashboardController *dashboard.DashboardController
}

func (r *DashboardRouteController) InitRoute(e *echo.Echo) {
	d := e.Group("/v1/dashboards")
	d.Use(middlewares.JWTMiddleware)
	d.GET("/admin", r.DashboardController.GetAdminDashboard)
	d.GET("/customer", r.DashboardController.GetCustomerDashboard)
}
