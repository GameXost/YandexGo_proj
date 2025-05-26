package prometh

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	RideAcceptedCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "driver_ride_accepted_total",
		Help: "Total number of accepted rides",
	})

	RideCompletedCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "driver_ride_completed_total",
		Help: "Total number of completed rides",
	})

	RideCanceledCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "driver_ride_canceled_total",
		Help: "Total number of canceled rides",
	})
)

// InitPrometheus запускает endpoint /prometh
func InitPrometheus(addr string) {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		_ = http.ListenAndServe(addr, nil)
	}()
}
