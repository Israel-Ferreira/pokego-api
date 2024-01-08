package middlewares

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func JsonMiddleware() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			hf(w, r)
		}
	}
}
