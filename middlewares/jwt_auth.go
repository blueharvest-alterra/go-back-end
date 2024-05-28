package middlewares

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
)

type Claims struct {
	ID       uuid.UUID
	Email    string
	FullName string
	Role     string
	jwt.RegisteredClaims
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			err := constant.ErrTokenNotFound
			return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				err := constant.ErrTokenNotValid
				return nil, c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
			}

			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			err := constant.ErrTokenNotValid
			return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
		}

		if !token.Valid {
			err := constant.ErrTokenNotValid
			return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			err := constant.ErrTokenNotValid
			return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
		}

		c.Set("claims", claims)

		return next(c)
	}
}
