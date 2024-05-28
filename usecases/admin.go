package usecases

import (
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase struct {
	repository entities.AdminRepositoryInterface
}

func NewAdminUseCase(repository entities.AdminRepositoryInterface) *AdminUseCase {
	return &AdminUseCase{
		repository: repository,
	}
}

func (a *AdminUseCase) Login(admin *entities.Admin) (entities.Admin, error) {
	if admin.Auth.Email == "" || admin.Auth.Password == "" {
		return entities.Admin{}, constant.ErrEmptyInput
	}

	password := admin.Auth.Password

	if err := a.repository.Login(admin); err != nil {
		return entities.Admin{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Auth.Password), []byte(password)); err != nil {
		return entities.Admin{}, constant.ErrInvalidEmailOrPassword
	}

	return *admin, nil
}
