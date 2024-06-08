package request

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type EditFarmInvestRequest struct {
	ID          uuid.UUID
}

func (r *EditFarmInvestRequest) ToEntities() *entities.FarmInvest {
	return &entities.FarmInvest{
		ID: r.ID,
	}
}
