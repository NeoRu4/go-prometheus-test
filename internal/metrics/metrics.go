package metrics

import "github.com/prometheus/client_golang/prometheus"

func LatencyOfQuerySummary(metricName string) *prometheus.HistogramVec {
	latencyHi := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: metricName,
			Help: "Продолжительность выполнения запроса",
		},
		[]string{"method"},
	)

	prometheus.MustRegister(latencyHi)

	return latencyHi
}

func CountOfSuccess(metricName string) prometheus.Counter {
	success := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: metricName,
			Help: "Количество запросов",
		},
	)

	return success
}
