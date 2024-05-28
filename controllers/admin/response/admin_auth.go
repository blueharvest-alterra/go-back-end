package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type AdminAuthResponse struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	UserToken string    `json:"user_token"`
}

func AuthResponseFromUseCase(user *entities.Admin, token string) *AdminAuthResponse {
	return &AdminAuthResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Auth.Email,
		UserToken: token,
	}
}
