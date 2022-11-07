package main

import "github.com/prometheus/client_golang/prometheus"

var QueueGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "queue_num_total",
	},
	[]string{"name"},
)

func init() {
	prometheus.MustRegister(QueueGauge)
}
