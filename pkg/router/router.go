package router

import (
	"github.com/daydreme/classcharts-server-mock/pkg/handlers"
	parentUser "github.com/daydreme/classcharts-server-mock/pkg/handlers/parent/user"
	"net/http"

	studentData "github.com/daydreme/classcharts-server-mock/pkg/handlers/student/data"
	studentUser "github.com/daydreme/classcharts-server-mock/pkg/handlers/student/user"

	"github.com/gorilla/mux"
)

func CreateMuxRouter() *mux.Router {
	// StrictSlash(true) is used to make sure that the router will automatically redirect requests with a trailing slash to the equivalent URL without a trailing slash.
	// Honestly, this is more of a preference thing. I don't think it's necessary to have this, but it's good to have it anyway. Some dev out there might forget that this doesn't use trailing slashes, and end up spending 2 hours debugging why their code isn't working.
	// This is just a safety net to save some unobservant people :)
	router := mux.NewRouter().StrictSlash(true)
	router.Use(handlers.ErrorHandler)
	router.Use(handlers.RequestHandler)

	CreateStudentRoutes(router.PathPrefix("/apiv2student").Subrouter(), true)
	CreateParentRoutes(router.PathPrefix("/apiv2parent").Subrouter())

	return router
}

func CreateStudentRoutes(v2student *mux.Router, includeExtras bool) *mux.Router {
	//v1student := router.PathPrefix("/student").Subrouter()

	if includeExtras {
		v2student.HandleFunc("/hasdob", studentUser.HasDOBHandler).Methods(http.MethodPost)
		v2student.HandleFunc("/login", studentUser.LoginHandler).Methods(http.MethodPost)
		v2student.HandleFunc("/ping", studentUser.StudentUserHandler).Methods(http.MethodPost)
		v2student.HandleFunc("/getcode", studentUser.GetCodeHandler).Methods(http.MethodPost)
		v2student.HandleFunc("/logout", studentUser.LogoutHandler).Methods(http.MethodPost)
	}

	v2student.HandleFunc("/behaviour/{student}", studentData.GetBehaviourHandler).Methods(http.MethodGet)
	v2student.HandleFunc("/behaviour", studentData.GetBehaviourHandler).Methods(http.MethodGet)

	v2student.HandleFunc("/activity/{student}", studentData.GetActivityHandler).Methods(http.MethodGet)
	v2student.HandleFunc("/activity", studentData.GetActivityHandler).Methods(http.MethodGet)

	v2student.HandleFunc("/announcements/{student}", studentData.GetAnnouncementHandler).Methods(http.MethodGet)
	v2student.HandleFunc("/announcements", studentData.GetAnnouncementHandler).Methods(http.MethodGet)

	v2student.HandleFunc("/rewards/{student}", studentData.GetRewardHandler).Methods(http.MethodGet)
	v2student.HandleFunc("/rewards", studentData.GetRewardHandler).Methods(http.MethodGet)

	v2student.HandleFunc("/purchase/{itemID}", studentData.GetPurchaseHandler).Methods(http.MethodPost)

	return v2student
}

func CreateParentRoutes(v2parent *mux.Router) *mux.Router {
	v2parent.HandleFunc("/login", parentUser.LoginHandler).Methods(http.MethodPost)
	v2parent.HandleFunc("/ping", parentUser.ParentUserHandler).Methods(http.MethodPost)
	v2parent.HandleFunc("/logout", parentUser.LogoutHandler).Methods(http.MethodPost)

	v2parent.HandleFunc("/pupils", parentUser.GetPupilsHandler).Methods(http.MethodGet)

	CreateStudentRoutes(v2parent, false) // Creates all the /apiv2parent/behaviour, /apiv2parent/activity, etc. routes

	return v2parent
}
