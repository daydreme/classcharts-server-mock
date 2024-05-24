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
	r.HandleFunc("/apiv2student/ping", handlers.UserHandler).Methods(http.MethodPost)

	return r
}
