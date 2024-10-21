package handlers

import (
	"net/http"

	"github.com/daydreme/classcharts-server-mock/pkg/models"
	"github.com/daydreme/classcharts-server-mock/pkg/responses"
)

type timeTableResponseData []models.Lesson

type timeTableResponseMeta struct {
	Dates          []string        `json:"dates"`
	TimetableDates []string        `json:"timetable_dates"`
	Periods        []models.Period `json:"periods"`
	StartTime      string          `json:"start_time"`
	EndTime        string          `json:"end_time"`
}

func TimetableHandler(w http.ResponseWriter, r *http.Request) {
	periods := models.NewMockPeriods()

	meta := timeTableResponseMeta{
		Dates:          []string{"2024-05-29", "2024-05-30"},
		TimetableDates: []string{"2024-05-29", "2024-05-30"},
		Periods:        periods,
		StartTime:      "08:00",
		EndTime:        "14:00",
	}

	data := timeTableResponseData(models.NewMockLessons())

	response := responses.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}
