package handler

import "github.com/prometheus/client_golang/prometheus"
import "main/metrics"

var (
	AdultUserInc = metrics.MustRegister(prometheus.NewCounter(prometheus.CounterOpts{
		Name: "adult_users_total_session",
		Help: "Total number created adult users in the last session",
	}))
)
