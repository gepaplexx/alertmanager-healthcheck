package webserver

import (
	"alertmanager_healthcheck/logging"
	"alertmanager_healthcheck/metrics"
	"encoding/json"
	"fmt"
	"net/http"
)

// IncrementEndpoint Type for the endpoint where the counters in the metrics are incrementend
// metrics: Metrics Endpoint with the counter
// logger: Logger
type IncrementEndpoint struct {
	metrics metrics.Metrics
	logger  logging.Logger
}

// NewIncrementEndpoint Instantiates the IncrementEndpoint struct
func NewIncrementEndpoint(metrics metrics.Metrics, logger logging.Logger) IncrementEndpoint {
	var ie IncrementEndpoint
	ie.metrics = metrics
	ie.logger = logger
	return ie
}

// Serve Method for the Increment Endpoint
func (web IncrementEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	req := web.ReadJSON(r)
	seq := Parse(req.Alerts)
	web.metrics.IncrementSequence(seq)
	fmt.Fprint(w, seq)
}

// ReadJSON Transforms the body of an HTTP Request into a Request Struct
func (web IncrementEndpoint) ReadJSON(r *http.Request) Request {
	decoder := json.NewDecoder(r.Body)
	var req Request
	err := decoder.Decode(&req)
	if err != nil {
		web.logger.Error(err.Error())
	}
	return req
}

// Parse Parses the Alerts-Array into an sequence of labels
func Parse(alerts []Alert) []string {
	var arr []string
	for _, alert := range alerts {
		arr = append(arr, alert.Labels.ClusterDiscriminator)
	}
	return arr
}
