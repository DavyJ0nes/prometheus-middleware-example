package main

import (
	"net/http"

	"github.com/davyj0nes/prometheus-example/logger"
	"github.com/davyj0nes/prometheus-example/router"
)

func main() {
	router := router.New()

	logger.Info("Starting Server...")
	logger.Fatal(http.ListenAndServe(":8080", router))
}
