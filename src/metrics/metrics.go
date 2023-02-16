package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics Type for exposing and accessing the Metrics Endpoint
type Metrics struct {
	CounterVec *prometheus.CounterVsec
}

// NewMetrics Creates a new Metrics Instance
func NewMetrics(name string, help string, label string) Metrics {
	var metrics Metrics
	metrics.SetCounterVec(name, help, label)
	return metrics
}

// SetCounterVec Sets the CounterVec wrapped by this instance
func (metrics *Metrics) SetCounterVec(name string, help string, label string) {
	metrics.CounterVec = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: name,  // "alertmanager_status",
		Help: help}, //"The status of the alertmanager",},
		[]string{label}) //"cluster_discriminator"})
}

// Increment Increments the Counter with the label
func (metrics *Metrics) Increment(label string) {
	metrics.CounterVec.WithLabelValues(label).Inc()
}

// IncrementSequence Increments according to a shard of labels
func (metrics *Metrics) IncrementSequence(labels []string) {
	for _, label := range labels {
		metrics.Increment(label)
	}
}
