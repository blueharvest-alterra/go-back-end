package courier

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/controllers/courier/request"
	"github.com/blueharvest-alterra/go-back-end/controllers/courier/response"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CourierController struct {
	courierUseCase entities.CourierUseCaseInterface
}

func (ac *CourierController) GetAll(c echo.Context) error {
	var getAllCourierRequest request.GetAllCourier
	if err := c.Bind(&getAllCourierRequest); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	courierRequestAPI := entities.CostCourierAPIRequest{
		OriginCityID:      getAllCourierRequest.OriginCityID,
		DestinationCityID: getAllCourierRequest.DestinationCityID,
		Weight:            getAllCourierRequest.Weight,
	}

	couriers, errUseCase := ac.courierUseCase.GetAll(&[]entities.Courier{}, courierRequestAPI)
	if errUseCase != nil {
		return c.JSON(http.StatusBadRequest, errUseCase)
	}

	jobOnProgressResponse := response.GetAllCourierFromUseCase(&couriers)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("get all courier successfully", jobOnProgressResponse))
}

func NewCourierController(courierUseCase entities.CourierUseCaseInterface) *CourierController {
	return &CourierController{
		courierUseCase: courierUseCase,
	}
}
