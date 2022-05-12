package repositories

import "github.com/achrt/metrics-collector/internal/domain/models"

type Storage interface {
	Init()
	
	Get(code string) (*models.Metrics, error)
	Set(code string, val models.Metrics) error

	PrintMetrics() map[string]string
}
