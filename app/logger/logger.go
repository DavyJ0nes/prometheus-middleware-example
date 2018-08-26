package logger

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

// Initialise ensures that the logger componant is configured correctly
// This should be called from main.go init() function.
//
// For production service this should probably be pulled out as a dependency
// that should be injected at a higher level (maybe within router)
func Initialise() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

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

// Warn is a wrapper for log.Warn
func Warn(v ...interface{}) {
	log.Warn(v)
}
