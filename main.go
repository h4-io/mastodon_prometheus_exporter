package main

import (
	"log"
	"net/http"
	"os"

	"mastodon_prometheus_exporter/collector"

	"github.com/VictoriaMetrics/metrics"
)

func main() {
	instance := os.Getenv("MASTODON_INSTANCE")
	if instance == "" {
		log.Fatal("MASTODON_INSTANCE is not defined")
	}
	port := os.Getenv("EXPORTER_PORT")
	if port == "" {
		port = "8080"
	}

	mastodonCollector := collector.InitMastodonStats(instance)

	log.Println("start listening on " + port)
	http.HandleFunc("/metrics", func(w http.ResponseWriter, req *http.Request) {
		log.Println("new metrics query: start collecting")
		collector.GetMastodonStats(mastodonCollector)
		metrics.WritePrometheus(w, true)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
