package cart

import (
	"fmt"
	"net/http"

	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/controllers/cart/request"
	"github.com/blueharvest-alterra/go-back-end/controllers/cart/response"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CartController struct {
	CartUseCaseInterface entities.CartUseCaseInterface
}

func NewCartController(CartUseCaseInterface entities.CartUseCaseInterface) *CartController {
	return &CartController{
		CartUseCaseInterface: CartUseCaseInterface,
	}
}

func (cc *CartController) Create(c echo.Context) error {
	var cartCreate request.CreateCartRequest
	if err := c.Bind(&cartCreate); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse("failed parse token"))
	}

	cartCreate.CustomerID = userData.ID

	farm, errUseCase := cc.CartUseCaseInterface.Create(cartCreate.ToEntities())
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	farmResponse := response.CartResponseFromUseCase(&farm)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("cart created!", farmResponse))
}

func (cc *CartController) Update(c echo.Context) error {
	var cartEdit request.EditCartRequest
	if err := c.Bind(&cartEdit); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse("failed parse token"))
	}

	cartId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	cartEdit.ID = cartId
	cartEdit.CustomerID = userData.ID

	cart, errUseCase := cc.CartUseCaseInterface.Update(cartEdit.ToEntities())
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	cartResponse := response.CartResponseFromUseCase(&cart)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("cart updated!", cartResponse))
}

func (cc *CartController) GetById(c echo.Context) error {
	cartId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	farm, errUseCase := cc.CartUseCaseInterface.GetById(cartId)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	farmResponse := response.CartResponseFromUseCase(&farm)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get cart data!", farmResponse))
}

func (cc *CartController) Delete(c echo.Context) error {
	cartId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	farm, errUseCase := cc.CartUseCaseInterface.Delete(cartId)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	cartResponse := response.CartResponseFromUseCase(&farm)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success delete cart data!", cartResponse))

}

func (cc *CartController) GetAll(c echo.Context) error {
	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse("failed parse token"))
	}

	fmt.Println("value:", userData.ID)
	carts, errUseCase := cc.CartUseCaseInterface.GetAll(userData.ID)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	cartGetAllResponse := response.SliceFromUseCase(&carts)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get all cart user data!", cartGetAllResponse))
}
