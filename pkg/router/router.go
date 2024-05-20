package router

import (
	"github.com/daydreme/classcharts-server-mock/pkg/handlers"
	"github.com/gorilla/mux"
)

func CreateMuxRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(handlers.ErrorHandler)
	r.HandleFunc("/apiv2student/hasdob", handlers.HasDOBHandler).Methods("GET")

	return r
}
