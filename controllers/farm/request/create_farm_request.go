package request

import (
	"mime/multipart"

	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type CreateFarmRequest struct {
	ID          uuid.UUID
	Title       string `form:"title"`
	Description string `form:"description"`
	Picture     string
	PictureFile []*multipart.FileHeader `form:"picture_file"`
}

func (r *CreateFarmRequest) ToEntities() *entities.Farm {
	return &entities.Farm{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		Picture:     r.Picture,
	}
}
