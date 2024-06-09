package response

import "github.com/blueharvest-alterra/go-back-end/entities"

type FarmMonitorGetAll struct {
	FarmMonitor []FarmMonitorResponse `json:"farmmonitor"`
}

func SliceFromUseCase(farmMonitors *[]entities.FarmMonitor) *FarmMonitorGetAll {
	allFarmMonitors := make([]FarmMonitorResponse, len(*farmMonitors))
	for i, _farmMonitor := range *farmMonitors {
		allFarmMonitors[i] = FarmMonitorResponse{
			ID:              _farmMonitor.ID,
			FarmID:          _farmMonitor.FarmID,
			Temperature:     _farmMonitor.Temperature,
			PH:              _farmMonitor.PH,
			DissolvedOxygen: _farmMonitor.DissolvedOxygen,
		}
	}

	return &FarmMonitorGetAll{
		FarmMonitor: allFarmMonitors,
	}
}
