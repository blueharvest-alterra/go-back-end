package middlewares

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/labstack/echo/v4"
	"os"
)

func CallbackAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("x-callback-token")
		if authHeader != os.Getenv("XDT_CALLBACK_VERIFICATION") {
			err := constant.ErrNotAuthorized
			return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
		}

		return next(c)
	}
}
