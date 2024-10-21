package global

import (
	"fmt"
	"github.com/daydreme/classcharts-server-mock/pkg/global/models/responses"
	"net/http"
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
				response := responses.NewErrorfulResponse(fmt.Sprintf("%v", err))

				w.WriteHeader(http.StatusInternalServerError)
				response.Write(w)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
