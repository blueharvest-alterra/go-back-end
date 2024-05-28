package request

import "github.com/blueharvest-alterra/go-back-end/entities"

type CustomerRegister struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cr *CustomerRegister) ToEntities() *entities.Customer {
	return &entities.Customer{
		FullName: cr.FullName,
		Auth: entities.Auth{
			Email:    cr.Email,
			Password: cr.Password,
		},
	}
}
