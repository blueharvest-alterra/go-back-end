package request

import "github.com/blueharvest-alterra/go-back-end/entities"

type CreateAdmin struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ca *CreateAdmin) ToEntities() *entities.Admin {
	return &entities.Admin{
		FullName: ca.FullName,
		Auth: entities.Auth{
			Email:    ca.Email,
			Password: ca.Password,
		},
	}
}
