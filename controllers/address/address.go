package address

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/address/response"
	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AddressController struct {
	addressUseCase entities.AddressUseCaseInterface
}

func (ac *AddressController) GetAllStates(c echo.Context) error {
	addresses, errUseCase := ac.addressUseCase.GetAllStates(&[]entities.Address{})
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	addressResponse := response.GetAllStateFromUseCase(&addresses)
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("get all state successfully", addressResponse))
}

func (ac *AddressController) GetAllCities(c echo.Context) error {
	stateID := c.Param("stateID")

	addresses, errUseCase := ac.addressUseCase.GetAllCities(&[]entities.Address{}, stateID)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	addressResponse := response.GetAllCityFromUseCase(&addresses)
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("get all state successfully", addressResponse))
}

func NewAddressController(addressUseCase entities.AddressUseCaseInterface) *AddressController {
	return &AddressController{
		addressUseCase: addressUseCase,
	}
}
