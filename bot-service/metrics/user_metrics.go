package metrics

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"log/slog"
	"main/internal/repo/database"
	"time"
)

var (
	AdultUserCounter = MustRegister(prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "adult_users_total",
		Help: "Total number of adult users in the database.",
	}))
)

func RecordMetrics(repo *database.UserRepositoryImpl, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c, err := repo.AdultUserMetric()
			if err != nil {
				slog.ErrorContext(
					context.Background(),
					"get metric error",
					slog.String("err", err.Error()),
				)
			}
			AdultUserCounter.Set(float64(c))
		}
	}
}
