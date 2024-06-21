package response

import (
	"time"

	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type ArticleResponse struct {
	ID        uuid.UUID `json:"id"`
	AdminID   uuid.UUID `json:"admin_id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Picture   string    `json:"picture"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ArticleResponseFromUseCase(article *entities.Article) *ArticleResponse {
	return &ArticleResponse{
		ID:        article.ID,
		AdminID:   article.AdminID,
		Author:    article.Admin.FullName,
		Title:     article.Title,
		Content:   article.Content,
		Picture:   article.Picture,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
}
