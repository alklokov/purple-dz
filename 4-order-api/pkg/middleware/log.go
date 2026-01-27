package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		start := time.Now()
		next.ServeHTTP(wrapper, r)
		log.SetFormatter(&log.JSONFormatter{})
		log.Println(wrapper.StatusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
