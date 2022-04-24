package repositories

type Storage interface {
	UpdateMetric(code string, val float64) error
	UpdateCounter(code string, val int64)
	GetMetric(code string) (float64, error)
	GetCounter(code string) (int64, error)
	GetMetrics() map[string]float64
	GetCounters() map[string]int64
}
