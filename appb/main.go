package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	klog "k8s.io/klog/v2"
)

func main() {
	var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
	flag.Parse()
	// prometheus config
	histogram := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "http_durations_histogram_seconds",
		Help:    "Latency distrubution for HTTP calls to downstream service",
		Buckets: prometheus.DefBuckets,
	})
	prometheus.MustRegister(histogram)
	// generate random data
	go func() {
		for {
			callService(histogram)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	http.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func callService(histogram prometheus.Histogram) {

	val := rand.Float64()
	klog.Infof("observed sent for: %v", val)
	histogram.Observe(val)

}
