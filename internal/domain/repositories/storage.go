package repositories

import "github.com/achrt/metrics-collector/internal/domain/models"

type Storage interface {
	Get(code string) (*models.Metrics, error)
	Set(code string, val models.Metrics)

	UpdateMetric(code string, val float64) error
	UpdateCounter(code string, val int64)

	GetMetric(code string) (float64, error)
	GetCounter(code string) (int64, error)

	GetMetrics() map[string]float64
	GetCounters() map[string]int64
}
