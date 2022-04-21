package health

import (
	"errors"
	"runtime"
	"strconv"
)

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

func (HealthStat) GetType(code string) (val string, err error) {
	var ok bool
	if val, ok = metricTypes[code]; !ok {
		err = errors.New("uncnown metric name")
	}
	return
}

// MetricData возвращает текстовое представление метрики
func (h HealthStat) MetricData(code string) (metricType, name, value string, err error) {
	metricType, err = h.GetType(code)
	if err != nil {
		return
	}

	switch code {
	case Alloc:
		value = strconv.Itoa(int(h.Alloc))
	case BuckHashSys:
		value = strconv.Itoa(int(h.BuckHashSys))
	case Frees:
		value = strconv.Itoa(int(h.Frees))
	case GCCPUFraction:
		value = strconv.Itoa(int(h.GCCPUFraction))
	case GCSys:
		value = strconv.Itoa(int(h.GCSys))
	case HeapAlloc:
		value = strconv.Itoa(int(h.HeapAlloc))
	case HeapIdle:
		value = strconv.Itoa(int(h.HeapIdle))
	case HeapInuse:
		value = strconv.Itoa(int(h.HeapInuse))
	case HeapObjects:
		value = strconv.Itoa(int(h.HeapObjects))
	case HeapReleased:
		value = strconv.Itoa(int(h.HeapReleased))
	case HeapSys:
		value = strconv.Itoa(int(h.HeapSys))
	case LastGC:
		value = strconv.Itoa(int(h.LastGC))
	case Lookups:
		value = strconv.Itoa(int(h.Lookups))
	case MCacheInuse:
		value = strconv.Itoa(int(h.MCacheInuse))
	case MCacheSys:
		value = strconv.Itoa(int(h.MCacheSys))
	case MSpanInuse:
		value = strconv.Itoa(int(h.MSpanInuse))
	case MSpanSys:
		value = strconv.Itoa(int(h.MSpanSys))
	case Mallocs:
		value = strconv.Itoa(int(h.Mallocs))
	case NextGC:
		value = strconv.Itoa(int(h.NextGC))
	case NumForcedGC:
		value = strconv.Itoa(int(h.NumForcedGC))
	case NumGC:
		value = strconv.Itoa(int(h.NumGC))
	case OtherSys:
		value = strconv.Itoa(int(h.OtherSys))
	case PauseTotalNs:
		value = strconv.Itoa(int(h.PauseTotalNs))
	case StackInuse:
		value = strconv.Itoa(int(h.StackInuse))
	case StackSys:
		value = strconv.Itoa(int(h.StackSys))
	case Sys:
		value = strconv.Itoa(int(h.Sys))
	case TotalAlloc:
		value = strconv.Itoa(int(h.TotalAlloc))
	case PollCount:
		value = strconv.FormatInt(h.PollCount, 10)
	case RandomValue:
		value = strconv.Itoa(int(h.RandomValue))
	}

	name = code
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
}

var metricTypes = map[string]string{
	Alloc:         "gauge",
	BuckHashSys:   "gauge",
	Frees:         "gauge",
	GCCPUFraction: "gauge",
	GCSys:         "gauge",
	HeapAlloc:     "gauge",
	HeapIdle:      "gauge",
	HeapInuse:     "gauge",
	HeapObjects:   "gauge",
	HeapReleased:  "gauge",
	HeapSys:       "gauge",
	LastGC:        "gauge",
	Lookups:       "gauge",
	MCacheInuse:   "gauge",
	MCacheSys:     "gauge",
	MSpanInuse:    "gauge",
	MSpanSys:      "gauge",
	Mallocs:       "gauge",
	NextGC:        "gauge",
	NumForcedGC:   "gauge",
	NumGC:         "gauge",
	OtherSys:      "gauge",
	PauseTotalNs:  "gauge",
	StackInuse:    "gauge",
	StackSys:      "gauge",
	Sys:           "gauge",
	TotalAlloc:    "gauge",
	PollCount:     "counter",
	RandomValue:   "gauge",
}
