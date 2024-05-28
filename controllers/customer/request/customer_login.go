package request

import "github.com/blueharvest-alterra/go-back-end/entities"

type CustomerLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cl *CustomerLogin) ToEntities() *entities.Customer {
	return &entities.Customer{
		Auth: entities.Auth{
			Email:    cl.Email,
			Password: cl.Password,
		},
	}
}
