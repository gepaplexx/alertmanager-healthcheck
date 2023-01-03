package main

import (
	"alertmanager_healthcheck/metrics"
	"alertmanager_healthcheck/webserver"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"alertmanager_healthcheck/logging"
)

// Entrypoint for the Application
func main() {
    logger := logging.NewLogger()
	mux := CreateMux(logger)

    port := ":2112"

	logger.Info("Starting Alertmanager Health Check service on port " + port)
	http.ListenAndServe(port, mux)
}

// Creates the Mux that Serves /inc and /metrics
func CreateMux(logger logging.Logger) *http.ServeMux {
	mux := http.NewServeMux()
    web := webserver.NewIncrementEndpoint(CreateMetrics(), logging.NewLogger())
    mux.Handle("/inc", web)
	mux.Handle("/metrics", promhttp.Handler())
	return mux
}

// Creates the Metrics Counter Vector "alertmanager_status"
// That differentiates by the label "gepardec_cluster"c
func CreateMetrics() metrics.Metrics {
	var metrics metrics.Metrics
        metrics.SetCounterVec(
                "alertmanager_status",
                "The status of the alertmanager",
                "gepardec_cluster",
        )
	return metrics
}