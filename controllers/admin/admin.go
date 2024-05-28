package admin

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/admin/request"
	"github.com/blueharvest-alterra/go-back-end/controllers/admin/response"
	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

type AdminController struct {
	adminUseCase entities.AdminUseCaseInterface
}

func (ac *AdminController) Login(c echo.Context) error {
	var adminLogin request.AdminLogin
	if err := c.Bind(&adminLogin); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	admin, errUseCase := ac.adminUseCase.Login(adminLogin.ToEntities())
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	tokenExpires := jwt.NewNumericDate(time.Now().Add(time.Hour * 24))

	claims := &middlewares.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: tokenExpires,
		},
		ID:       admin.ID,
		Email:    admin.Auth.Email,
		FullName: admin.FullName,
		Role:     "admin",
	}

	token, errTokenCreation := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if errTokenCreation != nil {
		return c.JSON(utils.ConvertResponseCode(errTokenCreation), base.NewErrorResponse(errTokenCreation.Error()))
	}

	adminResponse := response.AuthResponseFromUseCase(&admin, token)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Berhasil login!", adminResponse))
}

func NewAdminController(adminUseCase entities.AdminUseCaseInterface) *AdminController {
	return &AdminController{
		adminUseCase: adminUseCase,
	}
}
