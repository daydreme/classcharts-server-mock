package data

import (
	"github.com/daydreme/classcharts-server-mock/pkg/models/global"
	"github.com/daydreme/classcharts-server-mock/pkg/models/responses"
	"net/http"
	"time"
)

type getBehaviourMeta struct {
	EndDate   string `json:"end_date"`
	StartDate string `json:"start_date"`
	StepSize  string `json:"step_size"`
}

func GetBehaviourHandler(w http.ResponseWriter, r *http.Request) {
	meta := getBehaviourMeta{
		StartDate: time.Now().Format(time.RFC3339),
		EndDate:   time.Now().AddDate(0, 1, 0).Format(time.RFC3339),
		StepSize:  "week",
	}

	data := global.NewMockBehaviour()
	response := responses.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}
