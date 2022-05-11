package health

import (
	"errors"
	"runtime"
	"strconv"
	"strings"

	"github.com/achrt/metrics-collector/internal/domain/models"
)

func IsExists(code string) bool {
	code = strings.ToLower(code)
	for _, c := range metricCodes {
		if c == code {
			return true
		}
	}
	return false
}

func (h *HealthStat) SetMetrics(r runtime.MemStats) {
	h.Alloc = float64(r.Alloc)
	h.BuckHashSys = float64(r.BuckHashSys)
	h.Frees = float64(r.Frees)
	h.GCCPUFraction = float64(r.Frees)
	h.GCSys = float64(r.GCSys)
	h.HeapAlloc = float64(r.HeapAlloc)
	h.HeapIdle = float64(r.HeapIdle)
	h.HeapInuse = float64(r.HeapInuse)
	h.HeapObjects = float64(r.HeapObjects)
	h.HeapReleased = float64(r.HeapReleased)
	h.HeapSys = float64(r.HeapSys)
	h.LastGC = float64(r.LastGC)
	h.Lookups = float64(r.Lookups)
	h.MCacheInuse = float64(r.MCacheInuse)
	h.MCacheSys = float64(r.MCacheSys)
	h.MSpanInuse = float64(r.MSpanInuse)
	h.MSpanSys = float64(r.MSpanSys)
	h.Mallocs = float64(r.Mallocs)
	h.NextGC = float64(r.NextGC)
	h.NumForcedGC = float64(r.NumForcedGC)
	h.NumGC = float64(r.NumGC)
	h.OtherSys = float64(r.OtherSys)
	h.PauseTotalNs = float64(r.PauseTotalNs)
	h.StackInuse = float64(r.StackInuse)
	h.StackSys = float64(r.StackSys)
	h.Sys = float64(r.Sys)
	h.TotalAlloc = float64(r.TotalAlloc)
}

// MetricCodes возвращает доступные метрики
func (HealthStat) MetricCodes() []string {
	return metricCodes
}

func (HealthStat) GetType(code string) (val string, err error) {
	var ok bool
	if val, ok = metricTypes[code]; !ok {
		err = errors.New("uncnown metric name")
	}
	return
}

// MetricData возвращает текстовое представление метрики
func (h HealthStat) MetricData(code string) (metricType, value string, err error) {
	metricType, err = h.GetType(code)
	if err != nil {
		return
	}

	switch code {
	case PollCount:
		value = strconv.FormatInt(h.PollCount, 10)
	case RandomValue:
		value = strconv.Itoa(int(h.RandomValue))
	default:
		v, _ := h.Metric(code)
		value = strconv.Itoa(int(v))
	}

	return
}

func (h HealthStat) MetricDataModel(code string) (m models.Metrics, err error) {
	m.MType, err = h.GetType(code)
	if err != nil {
		return
	}

	m.ID = code

	switch code {
	case PollCount:
		m.Delta = &h.PollCount
	case RandomValue:
		m.Value = &h.RandomValue
	default:
		v, _ := h.Metric(code)
		m.Value = &v
	}

	return
}

func (h HealthStat) Metric(code string) (value float64, err error) {
	if _, ok := metricTypes[code]; !ok {
		err = errors.New("uncnown metric name")
		return
	}

	switch code {
	case Alloc:
		value = h.Alloc
	case BuckHashSys:
		value = h.BuckHashSys
	case Frees:
		value = h.Frees
	case GCCPUFraction:
		value = h.GCCPUFraction
	case GCSys:
		value = h.GCSys
	case HeapAlloc:
		value = h.HeapAlloc
	case HeapIdle:
		value = h.HeapIdle
	case HeapInuse:
		value = h.HeapInuse
	case HeapObjects:
		value = h.HeapObjects
	case HeapReleased:
		value = h.HeapReleased
	case HeapSys:
		value = h.HeapSys
	case LastGC:
		value = h.LastGC
	case Lookups:
		value = h.Lookups
	case MCacheInuse:
		value = h.MCacheInuse
	case MCacheSys:
		value = h.MCacheSys
	case MSpanInuse:
		value = h.MSpanInuse
	case MSpanSys:
		value = h.MSpanSys
	case Mallocs:
		value = h.Mallocs
	case NextGC:
		value = h.NextGC
	case NumForcedGC:
		value = h.NumForcedGC
	case NumGC:
		value = h.NumGC
	case OtherSys:
		value = h.OtherSys
	case PauseTotalNs:
		value = h.PauseTotalNs
	case StackInuse:
		value = h.StackInuse
	case StackSys:
		value = h.StackSys
	case Sys:
		value = h.Sys
	case TotalAlloc:
		value = h.TotalAlloc
	}
	return
}

var metricCodes = []string{
	Alloc,
	BuckHashSys,
	Frees,
	GCCPUFraction,
	GCSys,
	HeapAlloc,
	HeapIdle,
	HeapInuse,
	HeapObjects,
	HeapReleased,
	HeapSys,
	LastGC,
	Lookups,
	MCacheInuse,
	MCacheSys,
	MSpanInuse,
	MSpanSys,
	Mallocs,
	NextGC,
	NumForcedGC,
	NumGC,
	OtherSys,
	PauseTotalNs,
	StackInuse,
	StackSys,
	Sys,
	TotalAlloc,
	PollCount,
	RandomValue,
}

var metricTypes = map[string]string{
	Alloc:         models.TypeGauge,
	BuckHashSys:   models.TypeGauge,
	Frees:         models.TypeGauge,
	GCCPUFraction: models.TypeGauge,
	GCSys:         models.TypeGauge,
	HeapAlloc:     models.TypeGauge,
	HeapIdle:      models.TypeGauge,
	HeapInuse:     models.TypeGauge,
	HeapObjects:   models.TypeGauge,
	HeapReleased:  models.TypeGauge,
	HeapSys:       models.TypeGauge,
	LastGC:        models.TypeGauge,
	Lookups:       models.TypeGauge,
	MCacheInuse:   models.TypeGauge,
	MCacheSys:     models.TypeGauge,
	MSpanInuse:    models.TypeGauge,
	MSpanSys:      models.TypeGauge,
	Mallocs:       models.TypeGauge,
	NextGC:        models.TypeGauge,
	NumForcedGC:   models.TypeGauge,
	NumGC:         models.TypeGauge,
	OtherSys:      models.TypeGauge,
	PauseTotalNs:  models.TypeGauge,
	StackInuse:    models.TypeGauge,
	StackSys:      models.TypeGauge,
	Sys:           models.TypeGauge,
	TotalAlloc:    models.TypeGauge,
	PollCount:     models.TypeCounter,
	RandomValue:   models.TypeGauge,
}
