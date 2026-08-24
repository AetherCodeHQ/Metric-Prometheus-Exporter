package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"
)

func main() {
	port := "9100"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Fprintf(w, "# HELP go_goroutines Number of running goroutines\n")
		fmt.Fprintf(w, "# TYPE go_goroutines gauge\n")
		fmt.Fprintf(w, "go_goroutines %d\n", runtime.NumGoroutine())
		fmt.Fprintf(w, "# HELP go_mem_alloc_bytes Memory allocated\n")
		fmt.Fprintf(w, "# TYPE go_mem_alloc_bytes gauge\n")
		fmt.Fprintf(w, "go_mem_alloc_bytes %d\n", m.Alloc)
		fmt.Fprintf(w, "# HELP go_uptime_seconds Uptime\n")
		fmt.Fprintf(w, "# TYPE go_uptime_seconds counter\n")
		fmt.Fprintf(w, "go_uptime_seconds %d\n", time.Since(startTime).Seconds())
	})
	fmt.Printf("prometheus exporter on :%s/metrics\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var startTime = time.Now()
