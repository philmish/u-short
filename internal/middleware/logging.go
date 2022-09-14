package middleware

import (
	"log"
	"net/http"
	"os"
)

func Logger(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.SetOutput(os.Stdout)
		log.Println(r.UserAgent())
		h.ServeHTTP(w, r)
	})
}
