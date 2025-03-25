package main

import "github.com/prometheus/client_golang/prometheus"

var HttpDurationHistogram = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_durations_histogram_seconds",
		Buckets: []float64{0.2, 0.5, 1, 2, 5, 10, 30},
	},
	[]string{"path"},
)

func init() {
	prometheus.MustRegister(HttpDurationHistogram)
}
