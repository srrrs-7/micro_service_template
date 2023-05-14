package log

import (
	"log"
	"net/http"
)

func NewLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("receive request: %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
