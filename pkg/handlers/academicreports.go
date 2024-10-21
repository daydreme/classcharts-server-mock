package handlers

import (
	"net/http"
	"slices"
	"strconv"

	"github.com/daydreme/classcharts-server-mock/pkg/models"
	"github.com/daydreme/classcharts-server-mock/pkg/responses"
	"github.com/gorilla/mux"
)

type listAcademicReportsData = []models.PreviewAcademicReport

func ListAcademicReportsHandler(w http.ResponseWriter, r *http.Request) {
	data := listAcademicReportsData(models.NewMockPreviewAcademicReports())

	response := responses.NewSuccessfulResponse(data)
	response.Write(w)
}

type getAcademicReportData = models.AcademicReport

func GetAcademicReportHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	allReports := models.NewMockAcademicReports()

	reportIndex := slices.IndexFunc(allReports, func(r models.AcademicReport) bool { return r.Id == id })

	data := getAcademicReportData(allReports[reportIndex])

	response := responses.NewSuccessfulResponse(data)
	response.Write(w)
}
