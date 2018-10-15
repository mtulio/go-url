package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func sendMetrics(metrics []Metric) {
	metricLocation := config.Location

	metricHTTPTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gourl_test_http_time_ms",
		Help: "GoURL test - HTTP Time Taken milliseconds.",
	})
	metricDNSTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gourl_test_dns_time_ms",
		Help: "GoURL test - DNS Time Taken milliseconds.",
	})
	registry := prometheus.NewRegistry()
	registry.MustRegister(metricHTTPTime, metricDNSTime)

	for _, m := range metrics {
		pusher := push.New(config.MetricPush, "gourl").
			Grouping("location", metricLocation).
			Gatherer(registry)
		pusher.Grouping("server", m.HTTPServer)
		pusher.Grouping("host", m.HTTPHost)
		metricHTTPTime.Add(float64(m.HTTPTimeTaken))

		if m.DNSTimeTaken != 0 {
			metricDNSTime.Add(float64(m.DNSTimeTaken))
		}
		pusher.Push()
	}
}
