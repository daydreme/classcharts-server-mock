package router

import (
	"github.com/daydreme/classcharts-server-mock/pkg/global"
	"github.com/daydreme/classcharts-server-mock/pkg/global/models/responses"
	parentUser "github.com/daydreme/classcharts-server-mock/pkg/parent/handlers/user"
	studentData "github.com/daydreme/classcharts-server-mock/pkg/student/handlers/data"
	studentUser "github.com/daydreme/classcharts-server-mock/pkg/student/handlers/user"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateMuxRouter() *mux.Router {
	// StrictSlash(true) is used to make sure that the router will automatically redirect requests with a trailing slash to the equivalent URL without a trailing slash.
	// Honestly, this is more of a preference thing. I don't think it's necessary to have this, but it's good to have it anyway. Some dev out there might forget that this doesn't use trailing slashes, and end up spending 2 hours debugging why their code isn't working.
	// This is just a safety net to save some unobservant people :)
	router := mux.NewRouter().StrictSlash(true)
	router.Use(global.ErrorHandler)
	router.Use(global.RequestHandler)

	CreateStudentRoutes(router.PathPrefix("/apiv2student").Subrouter(), true)
	CreateStudentV1Routes(router.PathPrefix("/student").Subrouter())

	CreateParentRoutes(router.PathPrefix("/apiv2parent").Subrouter())
	//CreateParentReportAbsenceRoutes(router.PathPrefix("/apiv2parentreportabsence").Subrouter())

	CreateTestRouter(router.PathPrefix("/test").Subrouter())

	return router
}

func CreateStudentRoutes(v2student *mux.Router, includeExtras bool) *mux.Router {
	if includeExtras {
		v2student.HandleFunc("/hasdob", studentUser.HasDOBHandler).Methods(http.MethodPost)
		v2student.HandleFunc("/login", studentUser.LoginHandler).Methods(http.MethodPost)
		v2student.HandleFunc("/ping", studentUser.StudentUserHandler).Methods(http.MethodPost)
		v2student.HandleFunc("/getcode", studentUser.GetCodeHandler).Methods(http.MethodPost)
		v2student.HandleFunc("/logout", studentUser.LogoutHandler).Methods(http.MethodPost)
	}

	v2student.HandleFunc("/behaviour/{studentId}", studentData.GetBehaviourHandler).Methods(http.MethodGet)
	v2student.HandleFunc("/activity/{studentId}", studentData.GetActivityHandler).Methods(http.MethodGet)

	v2student.HandleFunc("/announcements/{studentId}", studentData.GetAnnouncementHandler).Methods(http.MethodGet)

	v2student.HandleFunc("/addconcern", studentData.AddConcernHandler).Methods(http.MethodPost)

	v2student.HandleFunc("/getacademicreports", studentData.ListAcademicReportsHandler).Methods(http.MethodGet)
	v2student.HandleFunc("/getacademicreport/{id}", studentData.GetAcademicReportHandler).Methods(http.MethodGet)

	v2student.HandleFunc("/getpupilreportcards", studentData.ListOnReportCardsHandler).Methods(http.MethodPost)
	v2student.HandleFunc("/getpupilreportcard/{id}", studentData.GetOnReportCardHandler).Methods(http.MethodGet)
	v2student.HandleFunc("/getpupilreportcardsummarycomment/{id}", studentData.GetOnReportCardSummaryCommentHandler).Methods(http.MethodGet)
	v2student.HandleFunc("/getpupilreportcardtarget/{id}", studentData.GetOnReportCardTargetHandler).Methods(http.MethodGet)

	v2student.HandleFunc("/timetable/{studentId}", studentData.TimetableHandler).Methods(http.MethodGet)

	v2student.HandleFunc("/rewards/{studentId}", studentData.GetRewardHandler).Methods(http.MethodGet)
	v2student.HandleFunc("/purchase/{itemId}", studentData.GetPurchaseHandler).Methods(http.MethodPost)

	return v2student
}

func CreateParentRoutes(v2parent *mux.Router) *mux.Router {
	v2parent.HandleFunc("/login", parentUser.LoginHandler).Methods(http.MethodPost)
	v2parent.HandleFunc("/ping", parentUser.ParentUserHandler).Methods(http.MethodPost)
	v2parent.HandleFunc("/logout", parentUser.LogoutHandler).Methods(http.MethodPost)

	v2parent.HandleFunc("/pupils", parentUser.GetPupilsHandler).Methods(http.MethodGet)
	v2parent.HandleFunc("/announcements", studentData.GetAnnouncementHandler).Methods(http.MethodGet)

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
	router.HandleFunc("/newstudent", func(w http.ResponseWriter, r *http.Request) {
		student := global.CreateStudent(global.StudentDB{
			Id:        global.GetNextID(),
			Name:      r.FormValue("name"),
			FirstName: r.FormValue("first_name"),
			LastName:  r.FormValue("last_name"),
			DOB:       "2010-01-01",
			Code:      r.FormValue("code"),
		})

		response := responses.NewSuccessfulResponse(student)
		response.Write(w)
	}).Methods(http.MethodPost)

	return router
}
