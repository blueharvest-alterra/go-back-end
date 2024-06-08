package farminvest

import (
	"net/http"

	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/controllers/farminvest/request"
	"github.com/blueharvest-alterra/go-back-end/controllers/farminvest/response"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type FarmInvestController struct {
	FarmInvestUseCase entities.FarmInvestUseCaseInterface
}

func NewFarmInvestController(FarmInvestUseCase entities.FarmInvestUseCaseInterface) *FarmInvestController {
	return &FarmInvestController{
		FarmInvestUseCase: FarmInvestUseCase,
	}
}

func (fc *FarmInvestController) Create(c echo.Context) error {
	var farmInvestCreate request.CreateFarmInvestRequest
	if err := c.Bind(&farmInvestCreate); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	farmInvestCreate.CustomerID = userData.ID

	farm, errUseCase := fc.FarmInvestUseCase.Create(farmInvestCreate.ToEntities())
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	farmResponse := response.FarmInvestResponseFromUseCase(&farm)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Farm invest created!", farmResponse))
}

func (fc *FarmInvestController) GetById(c echo.Context) error {
	farmId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	farm, errUseCase := fc.FarmInvestUseCase.GetById(farmId)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	farmResponse := response.FarmInvestResponseFromUseCase(&farm)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get farm data!", farmResponse))

}

func (fc *FarmInvestController) GetAll(c echo.Context) error {
	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}
	farms, errUseCase := fc.FarmInvestUseCase.GetAll(userData.ID)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	farmGetAllResponse := response.SliceFromUseCase(&farms)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get all farm invest user data!", farmGetAllResponse))
}