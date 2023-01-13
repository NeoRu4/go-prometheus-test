package main

import (
	"fmt"
	"go-prometheus/internal/metrics"
	randomRest "go-prometheus/internal/random-rest"
)

func main() {

	fmt.Println("Starting..")

	metricsInstance := metrics.NewPrometheusInstance("/metrics", ":2112")
	metricsInstance.StartPrometheusHandler()
	defer metricsInstance.WaitSync()

	queryService := randomRest.NewQueryService()
	for i := 0; i < 500; i++ {
		go queryService.QueryOne(i)
	}
}
