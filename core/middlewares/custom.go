package middlewares

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func CustomLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Logging inside the middleware ", r.Method)
		next.ServeHTTP(w, r)
	})
}
