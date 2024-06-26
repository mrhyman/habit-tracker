package createuser

import (
	"github.com/prometheus/client_golang/prometheus"

	"main/metrics"
)

var (
	adultUserInc = metrics.MustRegister(prometheus.NewCounter(prometheus.CounterOpts{
		Name: "adult_users_total_session",
		Help: "Total number created adult users in the last session",
	}))
)
