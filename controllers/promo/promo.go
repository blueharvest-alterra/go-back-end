package promo

import (
	"net/http"

	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/controllers/promo/request"
	"github.com/blueharvest-alterra/go-back-end/controllers/promo/response"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PromoController struct {
	promoUseCase entities.PromoUseCaseInterface
}

func NewPromoController(promoUseCase entities.PromoUseCaseInterface) *PromoController {
	return &PromoController{
		promoUseCase: promoUseCase,
	}
}

func (pc *PromoController) Create(c echo.Context) error {
	var promoCreate request.CreatePromoRequest
	if err := c.Bind(&promoCreate); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	promo, errUseCase := pc.promoUseCase.Create(promoCreate.ToEntities())
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	promoResponse := response.FarmResponseFromUseCase(&promo)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Promo created!", promoResponse))
}

func (pc *PromoController) GetById(c echo.Context) error {
	promoId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	promo, errUseCase := pc.promoUseCase.GetById(promoId)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	promoResponse := response.FarmResponseFromUseCase(&promo)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get promo data!", promoResponse))

}

func (pc *PromoController) Update(c echo.Context) error {
	var promoEdit request.EditPromoRequest
	if err := c.Bind(&promoEdit); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	promoId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	promoEdit.ID = promoId

	promo, errUseCase := pc.promoUseCase.Update(promoEdit.ToEntities())
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	promoResponse := response.FarmResponseFromUseCase(&promo)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Promo updated!", promoResponse))
}

func (pc *PromoController) Delete(c echo.Context) error {
	promoId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	promo, errUseCase := pc.promoUseCase.Delete(promoId)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	promoResponse := response.FarmResponseFromUseCase(&promo)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success delete promo data!", promoResponse))
}

func (pc *PromoController) GetAll(c echo.Context) error {
	promos, errUseCase := pc.promoUseCase.GetAll(&[]entities.Promo{})
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	promoGetAllResponse := response.SliceFromUseCase(&promos)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get all promo data!", promoGetAllResponse))
}
