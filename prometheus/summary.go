package main

import "github.com/prometheus/client_golang/prometheus"

var HttpDurations = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_durations_seconds",
		Objectives: map[float64]float64{
			0.5: 0.05, 0.9: 0.01, 0.99: 0.001,
		},
	},
	[]string{"path"},
)

func init() {
	prometheus.MustRegister(HttpDurations)
}
