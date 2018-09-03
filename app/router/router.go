package router

import (
	"net/http"

	"github.com/davyj0nes/prometheus-example/app/logger"
	"github.com/davyj0nes/prometheus-example/app/metrics"
	"github.com/davyj0nes/prometheus-example/app/router/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"

	_ "net/http/pprof"
)

// New instantiates a new Router with handlers attached
func New() http.Handler {
	// mux := http.NewServeMux()
	mux := new(mux.Router)

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/other", handlers.OtherHandler)
	mux.Handle("/metrics", metrics.Handler())

	chain := alice.New(logger.Request, metrics.Measure).Then(mux)

	return chain
}
