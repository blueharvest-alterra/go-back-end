package routes

import (
	chatBot "github.com/blueharvest-alterra/go-back-end/controllers/chat-bot"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/labstack/echo/v4"
)

type ChatBotRouteController struct {
	ChatBotController *chatBot.ChatBotController
}

func (c *ChatBotRouteController) InitRoute(e *echo.Echo) {
	p := e.Group("/v1/chat-bot")
	p.Use(middlewares.JWTMiddleware)
	p.POST("/messages/send", c.ChatBotController.SendMessage)
	p.GET("/messages", c.ChatBotController.GetMessageLists)
}
