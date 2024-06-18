package request

import (
	"mime/multipart"

	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type CustomerEditProfile struct {
	ID          uuid.UUID
	Avatar      string
	FullName    string                  `form:"full_name"`
	Nickname    string                  `form:"nick_name"`
	PhoneNumber string                  `form:"phone_number"`
	Gender      string                  `form:"gender"`
	Email       string                  `form:"email"`
	AvatarFile  []*multipart.FileHeader `form:"avatar_file"`
}

func (r *CustomerEditProfile) ToEntities() *entities.Customer {
	return &entities.Customer{
		ID:          r.ID,
		Avatar:      r.Avatar,
		FullName:    r.FullName,
		NickName:    r.Nickname,
		PhoneNumber: r.PhoneNumber,
		Gender:      entities.GenderEnum(r.Gender),
		Auth: entities.Auth{
			Email: r.Email,
		},
	}
}
