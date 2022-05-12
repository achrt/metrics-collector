package health

import (
	"testing"

	"github.com/achrt/metrics-collector/internal/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMetricData(t *testing.T) {
	var testData float64 = 8000000
	var testDataVal = "8000000"
	hs := &HealthStat{
		Alloc: testData,
	}
	wrongName := "wrong-metric-name"
	_, _, err := hs.MetricData(wrongName)
	require.Error(t, err)

	mType, val, err := hs.MetricData(Alloc)
	require.NoError(t, err)
	assert.Equal(t, testDataVal, val)
	assert.Equal(t, models.TypeGauge, mType)
}
