package request

import (
	"mime/multipart"

	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type CreateArticleRequest struct {
	ID          uuid.UUID
	AdminID     uuid.UUID
	Title       string `form:"title"`
	Content     string `form:"content"`
	Picture     string
	PictureFile []*multipart.FileHeader `form:"picture_file"`
}

func (r *CreateArticleRequest) ToEntities() *entities.Article {
	return &entities.Article{
		ID:      r.ID,
		AdminID: r.AdminID,
		Title:   r.Title,
		Content: r.Content,
		Picture: r.Picture,
	}
}
