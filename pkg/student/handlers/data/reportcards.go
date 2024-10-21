package data

import (
	"fmt"
	"github.com/daydreme/classcharts-server-mock/pkg/global/models"
	"github.com/daydreme/classcharts-server-mock/pkg/global/models/responses"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func ListOnReportCardsHandler(w http.ResponseWriter, _ *http.Request) {
	data := models.NewMockPreviewOnReportCards()

	response := responses.NewSuccessfulResponse(data)
	response.Write(w)
}

type getOnReportCardData = models.OnReportCard
type getOnReportCardMeta struct {
	ForbiddenDaysOfWeek []int `json:"forbidden_days_of_week"`
}

func GetOnReportCardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	allReports := models.NewMockOnReportCards()

	reportIndex := slices.IndexFunc(allReports, func(r models.OnReportCard) bool { return r.Id == id })

	data := getOnReportCardData(allReports[reportIndex])
	meta := getOnReportCardMeta{
		ForbiddenDaysOfWeek: []int{2},
	}

	response := responses.NewSuccessfulMetaResponse(data, meta)
	response.Write(w)
}

type getOnReportCardSummaryCommentData = models.OnReportCardSummaryComment

func GetOnReportCardSummaryCommentHandler(w http.ResponseWriter, r *http.Request) {
	layoutIn := "02/01/2006 15:04"
	layoutOut := "2006-01-02"

	date := r.FormValue("date")

	allComments := models.NewMockOnReportCardSummaryComments()
	commentIndex := slices.IndexFunc(allComments, func(r models.OnReportCardSummaryComment) bool {
		t, err := time.Parse(layoutIn, r.Date)
		if err != nil {
			panic(fmt.Sprintf("Error parsing date: %v", err))
		}

		output := t.Format(layoutOut)
		fmt.Printf("Output: %v Date: %v\n", output, date)

		return output == date
	})

	data := getOnReportCardSummaryCommentData(allComments[commentIndex])

	response := responses.NewSuccessfulResponse(data)
	response.Write(w)
}

type getOnReportCardTargetData struct {
	Target  models.OnReportCardTarget         `json:"target"`
	Periods []models.OnReportCardTargetPeriod `json:"periods"`
}

func GetOnReportCardTargetHandler(w http.ResponseWriter, _ *http.Request) {
	allTargets := models.NewMockOnReportCardTargets()
	allPeriods := models.NewMockOnReportCardTargetPeriods()

	target := allTargets[0]
	periods := allPeriods

	data := getOnReportCardTargetData{target, periods}

	response := responses.NewSuccessfulResponse(data)
	response.Write(w)
}
