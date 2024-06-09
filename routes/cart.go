package routes

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/cart"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/labstack/echo/v4"
)

type CartRouteController struct {
	CartController *cart.CartController
}

func (r *CartRouteController) InitRoute(e *echo.Echo) {
	c := e.Group("/v1/carts")
	c.Use(middlewares.JWTMiddleware)
	c.POST("", r.CartController.Create)
	c.PUT("/:id", r.CartController.Update)
	c.DELETE("/:id", r.CartController.Delete)
	c.GET("", r.CartController.GetAll)
	c.GET("/:id", r.CartController.GetById)
}
