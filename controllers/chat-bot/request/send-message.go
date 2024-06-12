package request

import "github.com/blueharvest-alterra/go-back-end/entities"

type SendMessageRequest struct {
	Message string `json:"message"`
}

func (r *SendMessageRequest) ToEntities() *entities.ChatBot {
	allMessages := make([]entities.Message, 1)
	allMessages[0] = entities.Message{Content: r.Message, Role: "user"}

	return &entities.ChatBot{
		Message: allMessages,
	}
}
