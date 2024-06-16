package product

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/controllers/product/request"
	"github.com/blueharvest-alterra/go-back-end/controllers/product/response"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/google/uuid"
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
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("product created", productResponse))
}

func (ac *ProductController) GetByID(c echo.Context) error {
	var product entities.Product

	productID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	product.ID = productID

	product, errUseCase := ac.productUseCase.GetByID(&product)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	productResponse := response.ProductDetailFromUseCase(&product)
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("product created", productResponse))
}

func (ac *ProductController) GetAll(c echo.Context) error {
	products, errUseCase := ac.productUseCase.GetAll(&[]entities.Product{})
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	productResponse := response.SliceFromUseCase(&products)
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("product get all success", productResponse))
}

func (ac *ProductController) Update(c echo.Context) error {
	var productUpdate request.ProductUpdateRequest
	if err := c.Bind(&productUpdate); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	productID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	productEntities := productUpdate.ToEntities()
	productEntities.ID = productID

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	thumbnail := form.File["thumbnail"]

	product, errUseCase := ac.productUseCase.Update(productEntities, userData, thumbnail)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	productResponse := response.ProductDetailFromUseCase(&product)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("product updated", productResponse))
}

func (ac *ProductController) Delete(c echo.Context) error {
	var product entities.Product

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	productID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		if uuid.IsInvalidLengthError(err) {
			return c.JSON(http.StatusNotFound, base.NewErrorResponse(err.Error()))
		}
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	product.ID = productID

	product, errUseCase := ac.productUseCase.Delete(&product, userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	productResponse := response.ProductDetailFromUseCase(&product)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("product deleted", productResponse))
}

func NewProductController(productUseCase entities.ProductUseCaseInterface) *ProductController {
	return &ProductController{
		productUseCase: productUseCase,
	}
}
