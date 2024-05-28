package request

import "github.com/blueharvest-alterra/go-back-end/entities"

type AdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (al *AdminLogin) ToEntities() *entities.Admin {
	return &entities.Admin{
		Auth: entities.Auth{
			Email:    al.Email,
			Password: al.Password,
		},
	}
}
