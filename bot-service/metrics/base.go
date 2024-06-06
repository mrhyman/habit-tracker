package metrics

import "github.com/prometheus/client_golang/prometheus"

func MustRegister[TCollector prometheus.Collector](cs TCollector) TCollector {
	prometheus.MustRegister(cs)
	return cs
}
