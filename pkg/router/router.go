package router

import (
	"net/http"

	"github.com/daydreme/classcharts-server-mock/pkg/handlers"
	"github.com/gorilla/mux"
)

func CreateMuxRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(handlers.ErrorHandler)

	r.HandleFunc("/apiv2student/hasdob", handlers.HasDOBHandler).Methods(http.MethodPost)
	r.HandleFunc("/apiv2student/login", handlers.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/apiv2student/getcode", handlers.GetCodeHandler).Methods(http.MethodPost)
	r.HandleFunc("/apiv2student/logout", handlers.LogoutHandler).Methods(http.MethodPost)

	r.HandleFunc("/apiv2student/ping", handlers.UserHandler).Methods(http.MethodPost)

	r.HandleFunc("/apiv2student/behaviour/{studentId}", handlers.GetBehaviourHandler).Methods(http.MethodGet)
	r.HandleFunc("/apiv2student/timetable/{studentId}", handlers.TimetableHandler).Methods(http.MethodGet)

	r.HandleFunc("/apiv2student/getacademicreports", handlers.ListAcademicReportsHandler).Methods(http.MethodGet)
	r.HandleFunc("/apiv2student/getacademicreport/{id}", handlers.GetAcademicReportHandler).Methods(http.MethodGet)

	r.HandleFunc("/apiv2student/getpupilreportcards", handlers.ListOnReportCardsHandler).Methods(http.MethodPost)
	r.HandleFunc("/apiv2student/getpupilreportcard/{id}", handlers.GetOnReportCardHandler).Methods(http.MethodGet)

	r.HandleFunc("/apiv2student/getpupilreportcardsummarycomment/{id}", handlers.GetOnReportCardSummaryCommentHandler).Methods(http.MethodGet)
	r.HandleFunc("/apiv2student/getpupilreportcardtarget/{id}", handlers.GetOnReportCardTargetHandler).Methods(http.MethodGet)

	return r
}
