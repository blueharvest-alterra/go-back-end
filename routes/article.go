package routes

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/article"
	"github.com/labstack/echo/v4"
)

type ArticleRouteController struct {
	ArticleController *article.ArticleController
}

func (r *ArticleRouteController) InitRoute(e *echo.Echo) {
	c := e.Group("/v1/articles")
	c.POST("", r.ArticleController.Create)
	c.PUT("/:id", r.ArticleController.Update)
	c.DELETE("/:id", r.ArticleController.Delete)
	c.GET("", r.ArticleController.GetAll)
	c.GET("/:id", r.ArticleController.GetById)
}
