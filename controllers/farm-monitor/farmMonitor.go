package farm_monitor

import (
	"net/http"

	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/controllers/farm-monitor/request"
	"github.com/blueharvest-alterra/go-back-end/controllers/farm-monitor/response"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type FarmMonitorController struct {
	FarmMonitorUseCase entities.FarmMonitorUseCaseInterface
}

func NewFarmMonitorController(FarmMonitorUseCase entities.FarmMonitorUseCaseInterface) *FarmMonitorController {
	return &FarmMonitorController{
		FarmMonitorUseCase: FarmMonitorUseCase,
	}
}

func (fc *FarmMonitorController) Create(c echo.Context) error {
	var farMonitorCreate request.CreateFarmMonitorRequest
	if err := c.Bind(&farMonitorCreate); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	farmMonitor, errUseCase := fc.FarmMonitorUseCase.Create(farMonitorCreate.ToEntities())
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	farmMonitorResponse := response.FarmMonitorResponseFromUseCase(&farmMonitor)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Farm monitor created!", farmMonitorResponse))
}

func (fc *FarmMonitorController) Update(c echo.Context) error {
	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	var farmMonitorEdit request.EditFarmMonitorRequest
	if err := c.Bind(&farmMonitorEdit); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	farmMonitorId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		if uuid.IsInvalidLengthError(err) {
			return c.JSON(http.StatusNotFound, base.NewErrorResponse(err.Error()))
		}
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	farmMonitorEdit.ID = farmMonitorId

	farmMonitor, errUseCase := fc.FarmMonitorUseCase.Update(farmMonitorEdit.ToEntities(), userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	farmMonitorResponse := response.FarmMonitorResponseFromUseCase(&farmMonitor)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Farm monitor updated!", farmMonitorResponse))
}

func (fc *FarmMonitorController) GetById(c echo.Context) error {
	farmMonitorId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		if uuid.IsInvalidLengthError(err) {
			return c.JSON(http.StatusNotFound, base.NewErrorResponse(err.Error()))
		}
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	farmMonitor, errUseCase := fc.FarmMonitorUseCase.GetById(farmMonitorId)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	farmMonitorResponse := response.FarmMonitorResponseFromUseCase(&farmMonitor)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get farm monitor data!", farmMonitorResponse))

}

func (fc *FarmMonitorController) GetAllByFarmId(c echo.Context) error {
	farmId, err := uuid.Parse(c.Param("farmid"))
	if err != nil {
		if uuid.IsInvalidLengthError(err) {
			return c.JSON(http.StatusNotFound, base.NewErrorResponse(err.Error()))
		}
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	farmMonitors, errUseCase := fc.FarmMonitorUseCase.GetAllByFarmId(farmId)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	farmMonitorGetAllResponse := response.SliceFromUseCase(&farmMonitors)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get all farm monitor data!", farmMonitorGetAllResponse))
}
