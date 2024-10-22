package data

import (
	"net/http"

	"github.com/CommunityCharts/CCModels/shared"
)

type timeTableResponseData []shared.Lesson

type timeTableResponseMeta struct {
	Dates          []string        `json:"dates"`
	TimetableDates []string        `json:"timetable_dates"`
	Periods        []shared.Period `json:"periods"`
	StartTime      string          `json:"start_time"`
	EndTime        string          `json:"end_time"`
}

func TimetableHandler(w http.ResponseWriter, _ *http.Request) {
	periods := shared.NewMockPeriods()

	meta := timeTableResponseMeta{
		Dates:          []string{"2024-05-29", "2024-05-30"},
		TimetableDates: []string{"2024-05-29", "2024-05-30"},
		Periods:        periods,
		StartTime:      "08:00",
		EndTime:        "14:00",
	}

	data := timeTableResponseData(shared.NewMockLessons())

	response := shared.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}
