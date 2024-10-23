package router

import (
	"github.com/CommunityCharts/CCServerMock/pkg/backend/student/data"
	studentUser "github.com/CommunityCharts/CCServerMock/pkg/backend/student/user"
	test2 "github.com/CommunityCharts/CCServerMock/pkg/backend/test"
	"github.com/CommunityCharts/CCServerMock/pkg/frontend"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateMuxRouter() *mux.Router {
	// StrictSlash(true) is used to make sure that the router will automatically redirect requests with a trailing slash to the equivalent URL without a trailing slash.
	// Honestly, this is more of a preference thing. I don't think it's necessary to have this, but it's good to have it anyway. Some dev out there might forget that this doesn't use trailing slashes, and end up spending 2 hours debugging why their code isn't working.
	// This is just a safety net to save some unobservant people :)
	router := mux.NewRouter().StrictSlash(true)
	router.Use(ErrorHandler)
	router.Use(RequestHandler)

	router.PathPrefix("/public").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir(frontend.Dir("public")))))
	DoFrontendStuff(router)

	CreateStudentRoutes(router.PathPrefix("/apiv2student").Subrouter(), true)
	CreateStudentV1Routes(router.PathPrefix("/student").Subrouter())

	CreateParentRoutes(router.PathPrefix("/apiv2parent").Subrouter())
	//CreateParentReportAbsenceRoutes(router.PathPrefix("/apiv2parentreportabsence").Subrouter())

	CreateTestRouter(router.PathPrefix("/test").Subrouter())

	return router
}

func CreateStudentRoutes(v2student *mux.Router, includeExtras bool) *mux.Router {
	restrictedv2Student := v2student.PathPrefix("").Subrouter()
	restrictedv2Student.Use(AuthHandler)

	if includeExtras {
		v2student.HandleFunc("/hasdob", studentUser.HasDOBHandler).Methods(http.MethodPost)
		v2student.HandleFunc("/login", studentUser.LoginHandler).Methods(http.MethodPost)
		restrictedv2Student.HandleFunc("/ping", studentUser.StudentUserHandler).Methods(http.MethodPost)
		restrictedv2Student.HandleFunc("/getcode", studentUser.GetCodeHandler).Methods(http.MethodPost)
		v2student.HandleFunc("/logout", studentUser.LogoutHandler).Methods(http.MethodPost)
	}

	restrictedv2Student.HandleFunc("/behaviour/{studentId}", data.GetBehaviourHandler).Methods(http.MethodGet)
	restrictedv2Student.HandleFunc("/activity/{studentId}", data.GetActivityHandler).Methods(http.MethodGet)

	restrictedv2Student.HandleFunc("/announcements/{studentId}", data.GetAnnouncementsHandler).Methods(http.MethodGet)

	restrictedv2Student.HandleFunc("/addconcern", data.AddConcernHandler).Methods(http.MethodPost)

	restrictedv2Student.HandleFunc("/getacademicreports", data.ListAcademicReportsHandler).Methods(http.MethodGet)
	restrictedv2Student.HandleFunc("/getacademicreport/{id}", data.GetAcademicReportHandler).Methods(http.MethodGet)

	restrictedv2Student.HandleFunc("/getpupilreportcards", data.ListOnReportCardsHandler).Methods(http.MethodPost)
	restrictedv2Student.HandleFunc("/getpupilreportcard/{id}", data.GetOnReportCardHandler).Methods(http.MethodGet)
	restrictedv2Student.HandleFunc("/getpupilreportcardsummarycomment/{id}", data.GetOnReportCardSummaryCommentHandler).Methods(http.MethodGet)
	restrictedv2Student.HandleFunc("/getpupilreportcardtarget/{id}", data.GetOnReportCardTargetHandler).Methods(http.MethodGet)

	restrictedv2Student.HandleFunc("/timetable/{studentId}", data.TimetableHandler).Methods(http.MethodGet)

	return v2student
}

func CreateParentRoutes(v2parent *mux.Router) *mux.Router {
	//v2parent.HandleFunc("/login", user2.LoginHandler).Methods(http.MethodPost)
	//v2parent.HandleFunc("/ping", user2.ParentUserHandler).Methods(http.MethodPost)
	//v2parent.HandleFunc("/logout", user2.LogoutHandler).Methods(http.MethodPost)
	//
	//v2parent.HandleFunc("/pupils", user2.GetPupilsHandler).Methods(http.MethodGet)
	//v2parent.HandleFunc("/announcements", data.GetAnnouncementsHandler).Methods(http.MethodGet)

	CreateStudentRoutes(v2parent, false) // Creates all the /apiv2parent/behaviour, /apiv2parent/activity, etc. routes

	return v2parent
}

//func CreateParentReportAbsenceRoutes(v2parentreportabs *mux.Router) *mux.Router {
//	v2parentreportabs.HandleFunc("/getreportedabsences/{studentId}", parentData.ListReportedAbsencesHandler).Methods(http.MethodGet)
//	return v2parentreportabs
//}

func CreateStudentV1Routes(v1student *mux.Router) *mux.Router {
	v1student.HandleFunc("/checkpupilcode/{code}", studentUser.CheckPupilCodeHandler).Methods(http.MethodPost)

	return v1student
}

func CreateTestRouter(router *mux.Router) *mux.Router {
	restrictedTest := router.PathPrefix("").Subrouter()
	restrictedTest.Use(AuthHandler)

	router.HandleFunc("/newstudent", test2.CreateStudentHandler).Methods(http.MethodPost)
	router.HandleFunc("/getstudent", test2.GetStudentHandler).Methods(http.MethodGet)

	router.HandleFunc("/newschool", test2.CreateSchoolHandler).Methods(http.MethodPost)

	restrictedTest.HandleFunc("/newannouncement", data.CreateAnnouncementHandler).Methods(http.MethodPost)

	restrictedTest.HandleFunc("/newactivity", data.CreateActivityHandler).Methods(http.MethodPost)

	return router
}

func DoFrontendStuff(router *mux.Router) {
	router.HandleFunc("/", frontend.HomePageHandler).Methods(http.MethodGet)
	router.HandleFunc("/announcements", frontend.AnnouncementsPageHandler).Methods(http.MethodGet)
}
