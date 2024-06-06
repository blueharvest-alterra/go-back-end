package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type ArticleResponse struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Picture string    `json:"picture"`
}

func ArticleResponseFromUseCase(article *entities.Article) *ArticleResponse {
	return &ArticleResponse{
		ID:      article.ID,
		Title:   article.Title,
		Content: article.Content,
		Picture: article.Picture,
	}
}
