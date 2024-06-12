package response

import "github.com/blueharvest-alterra/go-back-end/entities"

type SendMessage struct {
	Message string `json:"message"`
}

func SendMessageFromUseCase(chat *entities.ChatBot) *SendMessage {
	return &SendMessage{
		Message: chat.Message[0].Content,
	}
}
