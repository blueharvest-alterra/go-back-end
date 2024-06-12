package usecases

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
)

type ChatBotUseCase struct {
	repository entities.ChatBotRepositoryInterface
}

func (c ChatBotUseCase) GetMessageLists(chat *entities.ChatBot, userData *middlewares.Claims) (entities.ChatBot, error) {
	if userData.Role != "customer" {
		return entities.ChatBot{}, constant.ErrNotAuthorized
	}

	if err := c.repository.GetMessageLists(chat, userData); err != nil {
		return entities.ChatBot{}, err
	}

	return *chat, nil
}

func (c ChatBotUseCase) SendMessage(chat *entities.ChatBot, userData *middlewares.Claims) (entities.ChatBot, error) {
	if userData.Role != "customer" {
		return entities.ChatBot{}, constant.ErrNotAuthorized
	}

	if chat.Message[0].Content == "" {
		return entities.ChatBot{}, constant.ErrEmptyInput
	}

	if err := c.repository.SendMessage(chat, userData); err != nil {
		return entities.ChatBot{}, err
	}

	return *chat, nil
}

func NewChatBotUseCase(repository entities.ChatBotRepositoryInterface) *ChatBotUseCase {
	return &ChatBotUseCase{
		repository: repository,
	}
}
