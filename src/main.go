package main

import (
	"alertmanager_healthcheck/logging"
	"alertmanager_healthcheck/metrics"
	"alertmanager_healthcheck/webserver"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// Entrypoint for the Application
func main() {
	logger := logging.NewLogger()
	mux := CreateMux(logger)

	port := ":2112"

	logger.Info("Starting Alertmanager Health Check service on port " + port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Println(err)
	}
}

// CreateMux Creates the Mux that Serves /inc and /metrics
func CreateMux(logger logging.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	web := webserver.NewIncrementEndpoint(CreateMetrics(), logging.NewLogger())
	mux.Handle("/alertmanager", web)
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", webserver.NewRootEndpoint())
	return mux
}

// CreateMetrics Creates the Metrics Counter Vector "alertmanager_status"
// That differentiates by the label "gepardec_cluster"c
func CreateMetrics() metrics.Metrics {
	var metric metrics.Metrics
	metric.SetCounterVec(
		"alertmanagerhealthcheck_received_health_checks",
		"The amount of health checks that got received from an alertmanager.",
		"cluster",
	)
	return metric
}
