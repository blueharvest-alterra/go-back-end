package product

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/controllers/product/request"
	"github.com/blueharvest-alterra/go-back-end/controllers/product/response"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductController struct {
	productUseCase entities.ProductUseCaseInterface
}

func (ac *ProductController) Create(c echo.Context) error {
	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	var productCreate request.ProductCreateRequest
	if err := c.Bind(&productCreate); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	thumbnail := form.File["thumbnail"]

	product, errUseCase := ac.productUseCase.Create(productCreate.ToEntities(), userData, thumbnail)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	productResponse := response.ProductDetailFromUseCase(&product)
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Berhasil membuat user Admin!", productResponse))
}

func NewProductController(productUseCase entities.ProductUseCaseInterface) *ProductController {
	return &ProductController{
		productUseCase: productUseCase,
	}
}
