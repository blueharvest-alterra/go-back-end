package routes

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/promo"
	"github.com/labstack/echo/v4"
)

type PromoRouteController struct {
	PromoController *promo.PromoController
}

func (r *PromoRouteController) InitRoute(e *echo.Echo) {
	c := e.Group("/v1/promos")
	c.POST("", r.PromoController.Create)
	c.PUT("/:id", r.PromoController.Update)
	c.DELETE("/:id", r.PromoController.Delete)
	c.GET("", r.PromoController.GetAll)
	c.GET("/:id", r.PromoController.GetById)
}
