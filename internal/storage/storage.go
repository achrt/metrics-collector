package storage

import (
	"errors"

	"github.com/achrt/metrics-collector/internal/domain/models/health"
)

// TODO: пока не ясно, как организовывать хранение, пусть будет просто модель
type Storage struct {
	s *health.HealthStat
}

func New() *Storage {
	return &Storage{s: &health.HealthStat{}}
}

func (str *Storage) Create(h *health.HealthStat) error {
	str.s = h
	return nil
}

func (str *Storage) UpdateMetric(code string, val float64) error {
	return str.updateMetric(code, val)
}

func (str *Storage) UpdateCounter(val int64) {
	str.s.PollCount = val
}

func (str *Storage) updateMetric(code string, val float64) error {

	switch code {
	case health.Alloc:
		str.s.Alloc = val
	case health.BuckHashSys:
		str.s.BuckHashSys = val
	case health.Frees:
		str.s.Frees = val
	case health.GCCPUFraction:
		str.s.GCCPUFraction = val
	case health.GCSys:
		str.s.GCSys = val
	case health.HeapAlloc:
		str.s.HeapAlloc = val
	case health.HeapIdle:
		str.s.HeapIdle = val
	case health.HeapInuse:
		str.s.HeapInuse = val
	case health.HeapObjects:
		str.s.HeapObjects = val
	case health.HeapReleased:
		str.s.HeapReleased = val
	case health.HeapSys:
		str.s.HeapSys = val
	case health.LastGC:
		str.s.LastGC = val
	case health.Lookups:
		str.s.Lookups = val
	case health.MCacheInuse:
		str.s.MCacheInuse = val
	case health.MCacheSys:
		str.s.MCacheSys = val
	case health.MSpanInuse:
		str.s.MSpanInuse = val
	case health.MSpanSys:
		str.s.MSpanSys = val
	case health.Mallocs:
		str.s.Mallocs = val
	case health.NextGC:
		str.s.NextGC = val
	case health.NumForcedGC:
		str.s.NumForcedGC = val
	case health.NumGC:
		str.s.NumGC = val
	case health.OtherSys:
		str.s.OtherSys = val
	case health.PauseTotalNs:
		str.s.PauseTotalNs = val
	case health.StackInuse:
		str.s.StackInuse = val
	case health.StackSys:
		str.s.StackSys = val
	case health.Sys:
		str.s.Sys = val
	case health.TotalAlloc:
		str.s.TotalAlloc = val
	default:
		return errors.New("uncnown metric name")
	}
	return nil
}
