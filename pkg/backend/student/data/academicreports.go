package data

import (
	"net/http"
	"slices"
	"strconv"

	"github.com/CommunityCharts/CCModels/shared"
	"github.com/gorilla/mux"
)

func ListAcademicReportsHandler(w http.ResponseWriter, _ *http.Request) {
	data := shared.NewMockPreviewAcademicReports()

	response := shared.NewSuccessfulResponse(data)
	response.Write(w)
}

type getAcademicReportData = shared.AcademicReport

func GetAcademicReportHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	allReports := shared.NewMockAcademicReports()

	reportIndex := slices.IndexFunc(allReports, func(r shared.AcademicReport) bool { return r.Id == id })

	data := getAcademicReportData(allReports[reportIndex])

	response := shared.NewSuccessfulResponse(data)
	response.Write(w)
}
