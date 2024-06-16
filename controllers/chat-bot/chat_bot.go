package chat_bot

import (
	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/controllers/chat-bot/request"
	"github.com/blueharvest-alterra/go-back-end/controllers/chat-bot/response"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ChatBotController struct {
	chatBotUseCase entities.ChatBotUseCaseInterface
}

func NewChatBotController(chatBotUseCase entities.ChatBotUseCaseInterface) *ChatBotController {
	return &ChatBotController{
		chatBotUseCase: chatBotUseCase,
	}
}

func (ac *ChatBotController) SendMessage(c echo.Context) error {
	var sendMessageRequest request.SendMessageRequest
	if err := c.Bind(&sendMessageRequest); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	chatBot, errUseCase := ac.chatBotUseCase.SendMessage(sendMessageRequest.ToEntities(), userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	chatBotResponse := response.GetAllFromUseCase(&chatBot)
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("send message successful", chatBotResponse))
}

func (ac *ChatBotController) GetMessageLists(c echo.Context) error {
	var chatBotRequest entities.ChatBot

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	chatBot, errUseCase := ac.chatBotUseCase.GetMessageLists(&chatBotRequest, userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	chatBotResponse := response.GetAllFromUseCase(&chatBot)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("get all message successful", chatBotResponse))
}
