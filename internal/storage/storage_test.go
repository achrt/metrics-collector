package storage

import (
	"testing"

	"github.com/achrt/metrics-collector/internal/domain/models/health"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	store = New()
	m.Run()
}

func TestUpdate(t *testing.T) {
	for code, val := range storeFiller {
		err := store.UpdateMetric(code, val)
		require.NoError(t, err)
	}

	wrongCode := "wrong-code"
	var val float64 = 456464
	err := store.UpdateMetric(wrongCode, val)
	require.Error(t, err)
}

var store *Storage
var storeFiller = map[string]float64{
	health.Alloc:         125854,
	health.BuckHashSys:   22,
	health.Frees:         37969,
	health.GCCPUFraction: 133421,
	health.GCSys:         32552,
	health.HeapAlloc:     325353,
	health.HeapIdle:      26511,
	health.HeapInuse:     22445,
	health.HeapObjects:   3245243,
	health.HeapReleased:  145344,
	health.HeapSys:       111987,
	health.LastGC:        1234,
	health.Lookups:       444,
	health.MCacheInuse:   7896,
	health.MCacheSys:     3839796,
	health.MSpanInuse:    3686862,
	health.MSpanSys:      3763836,
	health.Mallocs:       9879879,
	health.NextGC:        789795,
	health.NumForcedGC:   98098098,
	health.NumGC:         487847847,
	health.OtherSys:      3444,
	health.PauseTotalNs:  23455,
	health.StackInuse:    34257,
	health.StackSys:      8796969,
	health.Sys:           96969,
	health.TotalAlloc:    370707,
}