package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type CustomerProfileResponse struct {
	ID          uuid.UUID `json:"id"`
	Avatar      string    `json:"avatar"`
	FullName    string    `json:"full_name"`
	Nickname    string    `json:"nick_name"`
	PhoneNumber string    `json:"phone_number"`
	Gender      string    `json:"gender"`
	Email       string    `json:"email"`
}

func ProfileResponseFromUseCase(user *entities.Customer) *CustomerProfileResponse {
	return &CustomerProfileResponse{
		ID:          user.ID,
		Avatar:      user.Avatar,
		FullName:    user.FullName,
		Nickname:    user.NickName,
		PhoneNumber: user.PhoneNumber,
		Gender:      string(user.Gender),
		Email:       user.Auth.Email,
	}
}
