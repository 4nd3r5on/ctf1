package chi_utils

import (
	"net/http"

	chim "github.com/go-chi/chi/v5/middleware"
)

// TODO: add configuration

// Adds headers related to a CORS to every responce
func EnableCors(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")

		ww := chim.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)
	}

	return http.HandlerFunc(fn)
}
