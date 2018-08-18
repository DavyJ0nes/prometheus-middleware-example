package logger

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Request is a simple middleware logger
func Request(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.WithFields(log.Fields{
			"path":        req.URL.String(),
			"remote_addr": req.RemoteAddr,
		}).Info("Request Received")

		next.ServeHTTP(w, req)
	})
}

// Fatal is a wrapper for log.Fatal
func Fatal(v ...interface{}) {
	log.Fatal(v)
}

// Info is a wrapper for log.Info
func Info(v ...interface{}) {
	log.Info(v)
}
