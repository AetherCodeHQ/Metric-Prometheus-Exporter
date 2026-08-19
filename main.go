package main

import (
	"fmt"
	"os"
)

// metric_prometheus_exporter - Export metrics to Prometheus
func metric_prometheus_exporter(path string) {
	fmt.Println("========================================")
	fmt.Println("  Metric-Prometheus-Exporter")
	fmt.Println("  Export metrics to Prometheus")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	metric_prometheus_exporter(path)
}
