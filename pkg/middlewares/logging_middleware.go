package middlewares

import (
	"log"
	"net/http"
)

func LogMiddleware() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL.Path)
			log.Println("User-Agent: ", r.Header.Get("User-Agent"))
			log.Println(r.Header["Content-Type"])

			hf(w, r)
		}
	}
}
