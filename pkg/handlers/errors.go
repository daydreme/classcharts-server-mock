package handlers

import (
	"fmt"
	"net/http"

	"github.com/daydreme/classcharts-server-mock/pkg/responses"
)

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
