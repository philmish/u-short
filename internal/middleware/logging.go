package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Logger(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.SetOutput(os.Stdout)
		msg := fmt.Sprintf("[%s %s] Agent: %s", r.Method, r.URL.Path, r.UserAgent())
		log.Println(msg)
		h.ServeHTTP(w, r)
	})
}
