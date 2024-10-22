package data

import (
	"net/http"
	"time"

	"github.com/CommunityCharts/CCModels/shared"
)

type getBehaviourMeta struct {
	EndDate   string `json:"end_date"`
	StartDate string `json:"start_date"`
	StepSize  string `json:"step_size"`
}

func GetBehaviourHandler(w http.ResponseWriter, _ *http.Request) {
	meta := getBehaviourMeta{
		StartDate: time.Now().Format(time.RFC3339),
		EndDate:   time.Now().AddDate(0, 1, 0).Format(time.RFC3339),
		StepSize:  "week",
	}

	data := shared.NewMockBehaviour()
	response := shared.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}
