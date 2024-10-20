package handlers

import (
	"fmt"
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
