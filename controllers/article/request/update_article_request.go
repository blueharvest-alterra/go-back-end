package request

import (
	"mime/multipart"

	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type EditArticleRequest struct {
	ID          uuid.UUID
	Title       string `form:"title"`
	Content     string `form:"content"`
	Picture     string
	PictureFile []*multipart.FileHeader `form:"picture_file"`
}

func (r *EditArticleRequest) ToEntities() *entities.Article {
	return &entities.Article{
		ID:      r.ID,
		Title:   r.Title,
		Content: r.Content,
		Picture: r.Picture,
	}
}
