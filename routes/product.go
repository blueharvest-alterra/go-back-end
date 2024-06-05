package routes

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/product"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/labstack/echo/v4"
)

type ProductRouteController struct {
	ProductController *product.ProductController
}

func (r *ProductRouteController) InitRoute(e *echo.Echo) {
	p := e.Group("/v1/products")
	p.Use(middlewares.JWTMiddleware)
	p.POST("", r.ProductController.Create)
	p.GET("/:id", r.ProductController.GetByID)
	p.GET("", r.ProductController.GetAll)
}
