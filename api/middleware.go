package api

import (
	"net/http"
	"urlShortener/utils"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.Log.Info(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
