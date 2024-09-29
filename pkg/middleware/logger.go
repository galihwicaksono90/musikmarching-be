package middleware

import (
	"fmt"
	"net/http"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s \n", r.Method)

		next.ServeHTTP(w, r)
	})
}
