package routes

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/admin"
	"github.com/labstack/echo/v4"
)

type AdminRouteController struct {
	AdminController *admin.AdminController
}

func (r *AdminRouteController) InitRoute(e *echo.Echo) {
	e.POST("/v1/login/admin", r.AdminController.Login)
}
