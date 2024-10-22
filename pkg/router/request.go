package router

import (
	"fmt"
	"net/http"

	"github.com/CommunityCharts/CCModels/shared"
)

func RequestHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// Log the request for debugging purposes
			fmt.Printf("%s %s %s\n", r.Method, r.URL.Path, r.Proto)
		}()

		next.ServeHTTP(w, r)
	})
}

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				response := shared.NewErrorfulResponse(fmt.Sprintf("%v", err))

				w.WriteHeader(http.StatusInternalServerError)
				response.Write(w)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
