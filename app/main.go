package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/davyj0nes/prometheus-example/app/logger"
	"github.com/davyj0nes/prometheus-example/app/metrics"
	"github.com/davyj0nes/prometheus-example/app/router"
)

// init is run before main is called.
// using it here to manage configuration of dependencies.
// should use a proper dependency management solution in future.
func init() {
	logger.Initialise()
	metrics.Initialise()
}

func main() {
	port := flag.String("port", "8080", "port number to use with Web Service")
	flag.Parse()

	logger.Info("Starting Service...")
	rtr := router.New()

	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + *port,
		Handler: rtr,
	}

	go func() {
		logger.Fatal(srv.ListenAndServe())
	}()

	logger.Info("The service is ready to go...")

	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		logger.Warn("Got SIGINT...")
	case syscall.SIGTERM:
		logger.Warn("Got SIGTERM...")
	}

	logger.Info("The service is shutting down...")
	srv.Shutdown(context.Background())
	logger.Info("Done")
}
