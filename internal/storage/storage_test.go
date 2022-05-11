package storage

import (
	"testing"

	"github.com/achrt/metrics-collector/internal/domain/models"
	"github.com/achrt/metrics-collector/internal/domain/models/health"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	counterVal int64 = 345

	counterCode = "counter-code"
	wrongCode   = "wrong-code"

	store *Storage
)

func TestMain(m *testing.M) {
	store = New()
	m.Run()
}

func TestUpdate(t *testing.T) {
	for code, val := range storeFiller {
		err := store.UpdateMetric(code, val)
		require.NoError(t, err)

		value, err := store.GetMetric(code)
		require.NoError(t, err)
		assert.Equal(t, val, value)
	}

	_, err := store.GetMetric(wrongCode)
	require.Error(t, err)
}

func TestGetCounter(t *testing.T) {
	store.UpdateCounter(counterCode, counterVal)
	val, err := store.GetCounter(counterCode)
	require.NoError(t, err)
	assert.Equal(t, counterVal, val)

	_, err = store.GetCounter(wrongCode)
	require.Error(t, err)
}

func TestSetGet(t *testing.T) {
	var val float64 = 9879879
	m := models.Metrics{
		ID:    "Mallocs",
		MType: "gauge",
		Value: &val,
	}

	store.Set(m.ID, m)
	r, err := store.Get(m.ID)
	require.NoError(t, err)
	assert.Equal(t, m, *r)

}

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
