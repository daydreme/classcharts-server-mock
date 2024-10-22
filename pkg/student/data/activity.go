package data

import (
	"net/http"
	"time"

	"github.com/CommunityCharts/CCModels/shared"
)

type getActivityMeta struct {
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	LastId           int    `json:"last_id"`
	StepSize         string `json:"step_size"`
	DetentionAliasUC string `json:"detention_alias_uc"`
}

func GetActivityHandler(w http.ResponseWriter, _ *http.Request) {
	meta := getActivityMeta{
		StartDate:        time.Now().Format(time.RFC3339),
		EndDate:          time.Now().AddDate(0, 1, 0).Format(time.RFC3339),
		LastId:           1,
		StepSize:         "week",
		DetentionAliasUC: "Detention",
	}

	data := shared.NewMockActivities()
	response := shared.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}
