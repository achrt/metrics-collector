package repositories

import "github.com/achrt/metrics-collector/internal/domain/models"

type Storage interface {
	Get(code string) (*models.Metrics, error)
	Set(code string, val models.Metrics) error

	PrintMetrics() map[string]string // возвращает текстовое представление имеющихся метрик

	Load() error // загрузка метрик в in-memory из файла
	Cast() error // загрузка метрик в файл из in-memory
}
