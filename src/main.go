package main

import (
	"alertmanager_health/metrics"
	"alertmanager_health/webserver"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"alertmanager_health/logging"
)

// Entrypoint for the Application
func main() {
	mux := CreateMux()
	http.ListenAndServe(":2112", mux)
}

// Creates the Mux that Serves /inc and /metrics
func CreateMux() *http.ServeMux {
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