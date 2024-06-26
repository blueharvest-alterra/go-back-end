package customer

import (
	"net/http"
	"os"
	"time"

	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/controllers/customer/request"
	"github.com/blueharvest-alterra/go-back-end/controllers/customer/response"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CustomerController struct {
	customerUseCase entities.CustomerUseCaseInterface
}

func (ac *CustomerController) Login(c echo.Context) error {
	var customerLogin request.CustomerLogin
	if err := c.Bind(&customerLogin); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	customer, errUseCase := ac.customerUseCase.Login(customerLogin.ToEntities())
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	tokenExpires := jwt.NewNumericDate(time.Now().Add(time.Hour * 730000))

	claims := &middlewares.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: tokenExpires,
		},
		ID:       customer.ID,
		Email:    customer.Auth.Email,
		FullName: customer.FullName,
		Role:     "customer",
	}

	token, errTokenCreation := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if errTokenCreation != nil {
		return c.JSON(utils.ConvertResponseCode(errTokenCreation), base.NewErrorResponse(errTokenCreation.Error()))
	}

	customerResponse := response.AuthResponseFromUseCase(&customer, token)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Berhasil login!", customerResponse))
}

func (ac *CustomerController) Register(c echo.Context) error {
	var customerRegister request.CustomerRegister
	if err := c.Bind(&customerRegister); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	customer, errUseCase := ac.customerUseCase.Register(customerRegister.ToEntities())
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	tokenExpires := jwt.NewNumericDate(time.Now().Add(time.Hour * 24))

	claims := &middlewares.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: tokenExpires,
		},
		ID:       customer.ID,
		Email:    customer.Auth.Email,
		FullName: customer.FullName,
		Role:     "customer",
	}

	token, errTokenCreation := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if errTokenCreation != nil {
		return c.JSON(utils.ConvertResponseCode(errTokenCreation), base.NewErrorResponse(errTokenCreation.Error()))
	}

	customerResponse := response.AuthResponseFromUseCase(&customer, token)
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Berhasil mendaftarkan akun!", customerResponse))
}

func (ac *CustomerController) CreateAddress(c echo.Context) error {
	customerData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	var customerAddAddress request.CustomerAddAddress
	if err := c.Bind(&customerAddAddress); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	customer, errUseCase := ac.customerUseCase.AddAddress(customerAddAddress.AddAddressToEntities(customerData.ID))
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	customerResponse := response.AddressResponseFromUseCase(&customer)
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Berhasil menambahkan Alamat baru!", customerResponse))
}

func (ac *CustomerController) GetAddresses(c echo.Context) error {
	customerData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	customer, errUseCase := ac.customerUseCase.GetAddresses(&entities.Customer{ID: customerData.ID})
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	customerResponse := response.AddressesResponseFromUseCase(&customer)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Berhasil mendapatkan semua data addresses!", customerResponse))
}

func (ac *CustomerController) GetProfile(c echo.Context) error {
	customerData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	customer, errUseCase := ac.customerUseCase.GetProfile(&entities.Customer{ID: customerData.ID})
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	customerResponse := response.ProfileResponseFromUseCase(&customer)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Berhasil mendapatkan data profile pengguna!", customerResponse))
}

func (ac *CustomerController) EditProfile(c echo.Context) error {
	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	var profileEdit request.CustomerEditProfile
	if err := c.Bind(&profileEdit); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	profileEdit.ID = userData.ID

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constant.ErrEmptyInput.Error()))
	}

	picture := form.File["avatar_file"]

	if len(picture) > 1 {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Gambar Avatar Hanya Boleh Satu"))
	}
	for _, file := range picture {
		if !utils.IsImageFile(file.Filename) {
			return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Format file gambar tidak didukung"))
		}
	}

	customer, errUseCase := ac.customerUseCase.EditProfile(profileEdit.ToEntities(), picture)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	articleResponse := response.ProfileResponseFromUseCase(&customer)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("profile updated!", articleResponse))
}

func NewCustomerController(customerUseCase entities.CustomerUseCaseInterface) *CustomerController {
	return &CustomerController{
		customerUseCase: customerUseCase,
	}
}
