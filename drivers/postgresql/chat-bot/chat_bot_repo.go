package chat_bot

import (
	"context"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"os"
	"time"
)

type Repo struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (r Repo) GetMessageLists(chat *entities.ChatBot, userData *middlewares.Claims) error {
	chatBotDb := FromUseCase(chat)

	threadID := r.Redis.Get(ctx, "chat_bot:"+userData.ID.String()+":thread_id").Val()
	if threadID == "" {
		return nil
	}

	chatBotDb.ThreadID = threadID

	if err := chatBotDb.GetOpenAIMessageLists(); err != nil {
		return err
	}

	*chat = *chatBotDb.ToUseCase()
	return nil
}

var ctx = context.Background()

func (r Repo) SendMessage(chat *entities.ChatBot, userData *middlewares.Claims) error {
	chatBotDb := FromUseCase(chat)

	threadID := r.Redis.Get(ctx, "chat_bot:"+userData.ID.String()+":thread_id").Val()
	chatBotDb.ThreadID = threadID
	if threadID == "" {
		if err := chatBotDb.CreateOpenAIThread(); err != nil {
			return err
		}

		if err := r.Redis.Set(ctx, "chat_bot:"+userData.ID.String()+":thread_id", chatBotDb.ThreadID, 8*time.Hour).Err(); err != nil {
			return err
		}
	}

	chatBotDb.AssistantID = os.Getenv("OPEN_API_ASSISTANT_ID")

	if err := chatBotDb.CreateOpenAIMessage(); err != nil {
		return err
	}

	if err := chatBotDb.RunOpenAIThread(); err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	if err := chatBotDb.GetOpenAIMessageLists(); err != nil {
		return err
	}

	*chat = *chatBotDb.ToUseCase()
	return nil
}

func NewChatBotRepo(db *gorm.DB, redis *redis.Client) *Repo {
	return &Repo{DB: db, Redis: redis}
}
