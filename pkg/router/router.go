package router

import (
	"github.com/daydreme/classcharts-server-mock/pkg/global"
	"github.com/daydreme/classcharts-server-mock/pkg/global/models/responses"
	user2 "github.com/daydreme/classcharts-server-mock/pkg/parent/handlers/user"
	"github.com/daydreme/classcharts-server-mock/pkg/student/handlers/data"
	"github.com/daydreme/classcharts-server-mock/pkg/student/handlers/user"
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
	CreateParentRoutes(router.PathPrefix("/apiv2parent").Subrouter())

	testRouter := router.PathPrefix("/test").Subrouter()
	testRouter.HandleFunc("/newstudent", func(w http.ResponseWriter, r *http.Request) {
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

func CreateStudentRoutes(v2student *mux.Router, includeExtras bool) *mux.Router {
	if includeExtras {
		v2student.HandleFunc("/hasdob", user.HasDOBHandler).Methods(http.MethodPost)
		v2student.HandleFunc("/login", user.LoginHandler).Methods(http.MethodPost)
		v2student.HandleFunc("/ping", user.StudentUserHandler).Methods(http.MethodPost)
		v2student.HandleFunc("/getcode", user.GetCodeHandler).Methods(http.MethodPost)
		v2student.HandleFunc("/logout", user.LogoutHandler).Methods(http.MethodPost)
	}

	v2student.HandleFunc("/behaviour/{student}", data.GetBehaviourHandler).Methods(http.MethodGet)
	v2student.HandleFunc("/activity/{student}", data.GetActivityHandler).Methods(http.MethodGet)
	v2student.HandleFunc("/announcements/{student}", data.GetAnnouncementHandler).Methods(http.MethodGet)

	v2student.HandleFunc("/rewards/{student}", data.GetRewardHandler).Methods(http.MethodGet)
	v2student.HandleFunc("/purchase/{itemID}", data.GetPurchaseHandler).Methods(http.MethodPost)

	return v2student
}

func CreateParentRoutes(v2parent *mux.Router) *mux.Router {
	v2parent.HandleFunc("/login", user2.LoginHandler).Methods(http.MethodPost)
	v2parent.HandleFunc("/ping", user2.ParentUserHandler).Methods(http.MethodPost)
	v2parent.HandleFunc("/logout", user2.LogoutHandler).Methods(http.MethodPost)

	v2parent.HandleFunc("/pupils", user2.GetPupilsHandler).Methods(http.MethodGet)
	v2parent.HandleFunc("/announcements", data.GetAnnouncementHandler).Methods(http.MethodGet)

	CreateStudentRoutes(v2parent, false) // Creates all the /apiv2parent/behaviour, /apiv2parent/activity, etc. routes

	return v2parent
}

func CreateStudentV1Routes(v1student *mux.Router) *mux.Router {
	return v1student
}
