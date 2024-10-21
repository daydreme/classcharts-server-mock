package main

import (
	"fmt"
	"github.com/daydreme/classcharts-server-mock/pkg/global"
	"net/http"
	"strconv"

	"github.com/daydreme/classcharts-server-mock/pkg/router"
	muxHandlers "github.com/gorilla/handlers"
)

const PORT = 4000

func main() {
	r := router.CreateMuxRouter()

	headersOk := muxHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := muxHandlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := muxHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	handler := muxHandlers.CORS(headersOk, originsOk, methodsOk)(r)

	srv := &http.Server{
		Handler: handler,
		Addr:    ":" + strconv.Itoa(PORT),
	}

	global.InitDB()

	fmt.Printf("Binding to :%v\n", PORT)
	srv.ListenAndServe()
}
