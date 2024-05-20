package main

import (
	"fmt"
	"github.com/daydreme/classcharts-server-mock/pkg/router"
	"net/http"
	"strconv"
)

const PORT = 4000

func main() {
	r := router.CreateMuxRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + strconv.Itoa(PORT),
	}

	fmt.Printf("Binding to :%v\n", PORT)
	srv.ListenAndServe()
}
