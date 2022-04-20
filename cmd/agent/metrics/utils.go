package metrics

import (
	"errors"
	"strconv"
)

// MetricCodes возвращает доступные метрики
func (m *Monitor) MetricCodes() []string {
	return metricCodes
}

// MetricData возвращает текстовое представление метрики
func (m *Monitor) MetricData(code string) (mType, name, value string, err error) {
	var ok bool

	if mType, ok = metricTypes[code]; !ok {
		err = errors.New("uncnown metric name")
	}

	name = code

	switch code {
	case Alloc:
		value = strconv.Itoa(int(m.Alloc))
	case BuckHashSys:
		value = strconv.Itoa(int(m.BuckHashSys))
	case Frees:
		value = strconv.Itoa(int(m.Frees))
	case GCCPUFraction:
		value = strconv.Itoa(int(m.GCCPUFraction))
	case GCSys:
		value = strconv.Itoa(int(m.GCSys))
	case HeapAlloc:
		value = strconv.Itoa(int(m.HeapAlloc))
	case HeapIdle:
		value = strconv.Itoa(int(m.HeapIdle))
	case HeapInuse:
		value = strconv.Itoa(int(m.HeapInuse))
	case HeapObjects:
		value = strconv.Itoa(int(m.HeapObjects))
	case HeapReleased:
		value = strconv.Itoa(int(m.HeapReleased))
	case HeapSys:
		value = strconv.Itoa(int(m.HeapSys))
	case LastGC:
		value = strconv.Itoa(int(m.LastGC))
	case Lookups:
		value = strconv.Itoa(int(m.Lookups))
	case MCacheInuse:
		value = strconv.Itoa(int(m.MCacheInuse))
	case MCacheSys:
		value = strconv.Itoa(int(m.MCacheSys))
	case MSpanInuse:
		value = strconv.Itoa(int(m.MSpanInuse))
	case MSpanSys:
		value = strconv.Itoa(int(m.MSpanSys))
	case Mallocs:
		value = strconv.Itoa(int(m.Mallocs))
	case NextGC:
		value = strconv.Itoa(int(m.NextGC))
	case NumForcedGC:
		value = strconv.Itoa(int(m.NumForcedGC))
	case NumGC:
		value = strconv.Itoa(int(m.NumGC))
	case OtherSys:
		value = strconv.Itoa(int(m.OtherSys))
	case PauseTotalNs:
		value = strconv.Itoa(int(m.PauseTotalNs))
	case StackInuse:
		value = strconv.Itoa(int(m.StackInuse))
	case StackSys:
		value = strconv.Itoa(int(m.StackSys))
	case Sys:
		value = strconv.Itoa(int(m.Sys))
	case TotalAlloc:
		value = strconv.Itoa(int(m.TotalAlloc))
	case PollCount:
		value = strconv.Itoa(m.PollCount())
	case RandomValue:
		value = strconv.Itoa(int(m.RandomValue()))
	}
	return
}

const (
	Alloc         = "Alloc"
	BuckHashSys   = "BuckHashSys"
	Frees         = "Frees"
	GCCPUFraction = "GCCPUFraction"
	GCSys         = "GCSys"
	HeapAlloc     = "HeapAlloc"
	HeapIdle      = "HeapIdle"
	HeapInuse     = "HeapInuse"
	HeapObjects   = "HeapObjects"
	HeapReleased  = "HeapReleased"
	HeapSys       = "HeapSys"
	LastGC        = "LastGC"
	Lookups       = "Lookups"
	MCacheInuse   = "MCacheInuse"
	MCacheSys     = "MCacheSys"
	MSpanInuse    = "MSpanInuse"
	MSpanSys      = "MSpanSys"
	Mallocs       = "Mallocs"
	NextGC        = "NextGC"
	NumForcedGC   = "NumForcedGC"
	NumGC         = "NumGC"
	OtherSys      = "OtherSys"
	PauseTotalNs  = "PauseTotalNs"
	StackInuse    = "StackInuse"
	StackSys      = "StackSys"
	Sys           = "Sys"
	TotalAlloc    = "TotalAlloc"
	PollCount     = "PollCount"
	RandomValue   = "RandomValue"
)

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
}
