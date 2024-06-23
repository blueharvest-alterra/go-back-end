package chat_bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/google/uuid"
	"net/http"
	"os"
	"time"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatBot struct {
	CustomerID  uuid.UUID
	AssistantID string
	ThreadID    string
	Message     []Message
}

func FromUseCase(chat *entities.ChatBot) *ChatBot {
	allMessages := make([]Message, len(chat.Message))
	for i, msg := range chat.Message {
		allMessages[i] = Message{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	return &ChatBot{
		CustomerID:  chat.CustomerID,
		AssistantID: chat.AssistantID,
		ThreadID:    chat.ThreadID,
		Message:     allMessages,
	}
}

func (c *ChatBot) ToUseCase() *entities.ChatBot {
	allMessages := make([]entities.Message, len(c.Message))
	for i, msg := range c.Message {
		allMessages[i] = entities.Message{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	return &entities.ChatBot{
		CustomerID:  c.CustomerID,
		AssistantID: c.AssistantID,
		ThreadID:    c.ThreadID,
		Message:     allMessages,
	}
}

type ResponseOpenAIThread struct {
	ID string `json:"id"`
}

func (c *ChatBot) CreateOpenAIThread() error {
	url := "https://api.openai.com/v1/threads"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("OPEN_API_KEY"))
	req.Header.Add("OpenAI-Beta", "assistants=v2")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		fmt.Println("url:", url)
		fmt.Println("response Status:", res.Status)
		return constant.ErrOpenAICallAPI
	}

	var response ResponseOpenAIThread
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}

	c.ThreadID = response.ID

	return nil
}

func (c *ChatBot) CreateOpenAIMessage() error {
	url := "https://api.openai.com/v1/threads/" + c.ThreadID + "/messages"
	method := "POST"

	jsonPayload, err := json.Marshal(c.Message[0])
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("OPEN_API_KEY"))
	req.Header.Add("OpenAI-Beta", "assistants=v2")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		fmt.Println("url:", url)
		fmt.Println("response Status:", res.Status)
		return constant.ErrOpenAICallAPI
	}

	return nil
}

type RequestOpenAIRunThread struct {
	AssistantID         string `json:"assistant_id"`
	Stream              bool   `json:"stream"`
	MaxCompletionTokens int64  `json:"max_completion_tokens"`
}

func (c *ChatBot) RunOpenAIThread() error {
	url := "https://api.openai.com/v1/threads/" + c.ThreadID + "/runs"
	method := "POST"

	var payload RequestOpenAIRunThread
	payload.AssistantID = c.AssistantID
	payload.Stream = false
	payload.MaxCompletionTokens = 100
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("OPEN_API_KEY"))
	req.Header.Add("OpenAI-Beta", "assistants=v2")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		fmt.Println("url:", url)
		fmt.Println("response Status:", res.Status)
		return constant.ErrOpenAICallAPI
	}

	return nil
}

type GetOpenAIMessageListsResponse struct {
	Object  string                `json:"object"`
	Data    []MessageListResponse `json:"data"`
	FirstID string                `json:"first_id"`
	LastID  string                `json:"last_id"`
	HasMore bool                  `json:"has_more"`
}

type MessageListResponse struct {
	ID          string            `json:"id"`
	Object      string            `json:"object"`
	CreatedAt   int64             `json:"created_at"`
	AssistantID *string           `json:"assistant_id"`
	ThreadID    string            `json:"thread_id"`
	RunID       *string           `json:"run_id"`
	Role        string            `json:"role"`
	Content     []ContentResponse `json:"content"`
	Attachments []interface{}     `json:"attachments"`
	Metadata    interface{}       `json:"metadata"`
}

type ContentResponse struct {
	Type string       `json:"type"`
	Text TextResponse `json:"text"`
}

type TextResponse struct {
	Value       string        `json:"value"`
	Annotations []interface{} `json:"annotations"`
}

func (c *ChatBot) GetOpenAIMessageLists() error {
	url := "https://api.openai.com/v1/threads/" + c.ThreadID + "/messages?order=asc&limit=100"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("OPEN_API_KEY"))
	req.Header.Add("OpenAI-Beta", "assistants=v2")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		fmt.Println("url:", url)
		fmt.Println("response Status:", res.Status)
		return constant.ErrOpenAICallAPI
	}

	var response GetOpenAIMessageListsResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}

	fmt.Println("response", utils.PrettyPrint(response))
	fmt.Println("len(response.Data)", len(response.Data))

	allMessages := make([]Message, len(response.Data))
	for i, msg := range response.Data {
		if len(msg.Content) < 1 {
			time.Sleep(5 * time.Second)
			return c.GetOpenAIMessageLists()
		}
		allMessages[i] = Message{
			Role:    msg.Role,
			Content: msg.Content[0].Text.Value,
		}
	}

	c.Message = allMessages

	return nil
}
