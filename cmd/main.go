package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/daydreme/classcharts-server-mock/pkg/router"
	"github.com/gorilla/handlers"
)

const PORT = 4000

func main() {
	r := router.CreateMuxRouter()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	handler := handlers.CORS(headersOk, originsOk, methodsOk)(r)

	srv := &http.Server{
		Handler: handler,
		Addr:    ":" + strconv.Itoa(PORT),
	}

	fmt.Printf("Binding to :%v\n", PORT)
	srv.ListenAndServe()
}
