package response

import "github.com/blueharvest-alterra/go-back-end/entities"

type StateDetail struct {
	ID    string `json:"id"`
	State string `json:"name"`
}

type StateGetAll struct {
	States []StateDetail `json:"states"`
}

func GetAllStateFromUseCase(addresses *[]entities.Address) *StateGetAll {
	allCategories := make([]StateDetail, len(*addresses))
	for i, _address := range *addresses {
		allCategories[i] = StateDetail{
			ID:    _address.StateID,
			State: _address.State,
		}
	}

	return &StateGetAll{
		States: allCategories,
	}
}
