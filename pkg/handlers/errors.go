package handlers

import (
	"fmt"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, fmt.Sprintf("Error: %+v", err), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
