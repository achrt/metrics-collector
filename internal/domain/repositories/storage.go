package repositories

import "github.com/achrt/metrics-collector/internal/domain/models/health"

type Storage interface {
	Create(h *health.HealthStat) error
	UpdateMetric(code string, val float64) error
	UpdateCounter(val int64)
}
