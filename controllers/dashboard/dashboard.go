package dashboard

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/controllers/dashboard/response"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DashboardController struct {
	dashboardUseCase entities.DashboardUseCaseInterface
}

func (dc *DashboardController) GetAdminDashboard(c echo.Context) error {
	var dashboard entities.Dashboard

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	dashboard, errUseCase := dc.dashboardUseCase.AdminDashboard(&dashboard, userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	dashboardAdminResponse := response.AdminDashboardFromUseCase(&dashboard)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("get admin dashboard successfully", dashboardAdminResponse))
}

func (dc *DashboardController) GetCustomerDashboard(c echo.Context) error {
	var dashboard entities.Dashboard

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	dashboard, errUseCase := dc.dashboardUseCase.CustomerDashboard(&dashboard, userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	dashboardCustomerResponse := response.CustomerDashboardFromUseCase(&dashboard)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("get customer dashboard successfully", dashboardCustomerResponse))
}

func NewDashboardController(dashboardUseCase entities.DashboardUseCaseInterface) *DashboardController {
	return &DashboardController{
		dashboardUseCase: dashboardUseCase,
	}
}
