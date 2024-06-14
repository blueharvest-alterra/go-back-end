package entities

import (
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
)

type Message struct {
	Role    string
	Content string
}

type ChatBot struct {
	CustomerID  uuid.UUID
	AssistantID string
	ThreadID    string
	Message     []Message
}

type ChatBotRepositoryInterface interface {
	SendMessage(chat *ChatBot, userData *middlewares.Claims) error
	GetMessageLists(chat *ChatBot, userData *middlewares.Claims) error
}

type ChatBotUseCaseInterface interface {
	SendMessage(chat *ChatBot, userData *middlewares.Claims) (ChatBot, error)
	GetMessageLists(chat *ChatBot, userData *middlewares.Claims) (ChatBot, error)
}
