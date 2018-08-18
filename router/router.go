package router

import (
	"net/http"

	"github.com/davyj0nes/prometheus-example/logger"
	"github.com/davyj0nes/prometheus-example/metrics"
	"github.com/davyj0nes/prometheus-example/router/handlers"
	"github.com/justinas/alice"
)

// New instantiates a new Router with handlers attached
func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/other", handlers.OtherHandler)
	mux.Handle("/metrics", metrics.Handler())

	chain := alice.New(logger.Request, metrics.MeasureRequest).Then(mux)

	return chain
}
