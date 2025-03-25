package main

import "github.com/prometheus/client_golang/prometheus"

var AccessCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "api_request_total",
	},
	[]string{"method", "path"},
)

func init() {
	prometheus.MustRegister(AccessCounter)
}
