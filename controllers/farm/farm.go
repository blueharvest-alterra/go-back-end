package farm

import (
	"net/http"

	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/controllers/farm/request"
	"github.com/blueharvest-alterra/go-back-end/controllers/farm/response"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type FarmController struct {
	farmUseCase entities.FarmUseCaseInterface
}

func NewFarmController(farmUseCase entities.FarmUseCaseInterface) *FarmController {
	return &FarmController{
		farmUseCase: farmUseCase,
	}
}

func (fc *FarmController) Create(c echo.Context) error {
	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	var farmCreate request.CreateFarmRequest
	if err := c.Bind(&farmCreate); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constant.ErrEmptyInput.Error()))
	}

	picture := form.File["picture_file"]
	if len(picture) == 0 {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Gambar Farm Tidak Boleh Kosong"))
	}

	if len(picture) > 1 {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Gambar Farm Hanya Boleh Satu"))
	}
	for _, file := range picture {
		if !utils.IsImageFile(file.Filename) {
			return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Format file gambar tidak didukung"))
		}
	}

	farm, errUseCase := fc.farmUseCase.Create(farmCreate.ToEntities(), userData, picture)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	farmResponse := response.FarmResponseFromUseCase(&farm)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Farm created!", farmResponse))
}

func (fc *FarmController) GetById(c echo.Context) error {
	farmId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		if uuid.IsInvalidLengthError(err) {
			return c.JSON(http.StatusNotFound, base.NewErrorResponse(err.Error()))
		}
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	farm, errUseCase := fc.farmUseCase.GetById(farmId)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	farmResponse := response.FarmResponseFromUseCase(&farm)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get farm data!", farmResponse))

}

func (fc *FarmController) Update(c echo.Context) error {
	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	var farmEdit request.EditFarmRequest
	if err := c.Bind(&farmEdit); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	farmId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		if uuid.IsInvalidLengthError(err) {
			return c.JSON(http.StatusNotFound, base.NewErrorResponse(err.Error()))
		}
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	farmEdit.ID = farmId

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constant.ErrEmptyInput.Error()))
	}

	picture := form.File["picture_file"]

	if len(picture) > 1 {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Gambar Tanaman Hanya Boleh Satu"))
	}
	for _, file := range picture {
		if !utils.IsImageFile(file.Filename) {
			return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Format file gambar tidak didukung"))
		}
	}

	farm, errUseCase := fc.farmUseCase.Update(farmEdit.ToEntities(), userData, picture)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	farmResponse := response.FarmResponseFromUseCase(&farm)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Farm updated!", farmResponse))
}

func (fc *FarmController) Delete(c echo.Context) error {
	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}
	farmId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		if uuid.IsInvalidLengthError(err) {
			return c.JSON(http.StatusNotFound, base.NewErrorResponse(err.Error()))
		}
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	farm, errUseCase := fc.farmUseCase.Delete(farmId, userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	farmResponse := response.FarmResponseFromUseCase(&farm)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success delete farm data!", farmResponse))
}

func (fc *FarmController) GetAll(c echo.Context) error {
	farms, errUseCase := fc.farmUseCase.GetAll(&[]entities.Farm{})
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	farmGetAllResponse := response.SliceFromUseCase(&farms)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get all farm data!", farmGetAllResponse))
}
