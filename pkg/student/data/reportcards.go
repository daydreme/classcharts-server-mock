package data

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/CommunityCharts/CCModels/shared"
	"github.com/gorilla/mux"
)

func ListOnReportCardsHandler(w http.ResponseWriter, _ *http.Request) {
	data := shared.NewMockPreviewOnReportCards()

	response := shared.NewSuccessfulResponse(data)
	response.Write(w)
}

type getOnReportCardData = shared.OnReportCard
type getOnReportCardMeta struct {
	ForbiddenDaysOfWeek []int `json:"forbidden_days_of_week"`
}

func GetOnReportCardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	allReports := shared.NewMockOnReportCards()

	reportIndex := slices.IndexFunc(allReports, func(r shared.OnReportCard) bool { return r.Id == id })

	data := getOnReportCardData(allReports[reportIndex])
	meta := getOnReportCardMeta{
		ForbiddenDaysOfWeek: []int{2},
	}

	response := shared.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}

type getOnReportCardSummaryCommentData = shared.OnReportCardSummaryComment

func GetOnReportCardSummaryCommentHandler(w http.ResponseWriter, r *http.Request) {
	layoutIn := "02/01/2006 15:04"
	layoutOut := "2006-01-02"

	date := r.FormValue("date")

	allComments := shared.NewMockOnReportCardSummaryComments()
	commentIndex := slices.IndexFunc(allComments, func(r shared.OnReportCardSummaryComment) bool {
		t, err := time.Parse(layoutIn, r.Date)
		if err != nil {
			panic(fmt.Sprintf("Error parsing date: %v", err))
		}

		output := t.Format(layoutOut)
		fmt.Printf("Output: %v Date: %v\n", output, date)

		return output == date
	})

	data := getOnReportCardSummaryCommentData(allComments[commentIndex])

	response := shared.NewSuccessfulResponse(data)
	response.Write(w)
}

type getOnReportCardTargetData struct {
	Target  shared.OnReportCardTarget         `json:"target"`
	Periods []shared.OnReportCardTargetPeriod `json:"periods"`
}

func GetOnReportCardTargetHandler(w http.ResponseWriter, _ *http.Request) {
	allTargets := shared.NewMockOnReportCardTargets()
	allPeriods := shared.NewMockOnReportCardTargetPeriods()

	target := allTargets[0]
	periods := allPeriods

	data := getOnReportCardTargetData{target, periods}

	response := shared.NewSuccessfulResponse(data)
	response.Write(w)
}
