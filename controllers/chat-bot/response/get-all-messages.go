package response

import "github.com/blueharvest-alterra/go-back-end/entities"

type MessagesResponse struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GetAllMessagesResponse struct {
	Messages []MessagesResponse `json:"messages"`
}

func GetAllFromUseCase(chat *entities.ChatBot) *GetAllMessagesResponse {
	allMessages := make([]MessagesResponse, len(chat.Message))
	for i, msg := range chat.Message {
		allMessages[i] = MessagesResponse{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}
	return &GetAllMessagesResponse{
		Messages: allMessages,
	}
}
