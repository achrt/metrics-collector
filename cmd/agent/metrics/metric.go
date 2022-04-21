package metrics

import (
	"context"
	"runtime"
	"time"

	"github.com/achrt/metrics-collector/internal/domain/models/health"
)

type Monitor struct {
	health.HealthStat

	pollCount   int
	randomValue uint64

	duration int64
}

func New(duration int64) *Monitor {
	return &Monitor{duration: duration}
}

func (m *Monitor) Run(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	interval := time.Duration(m.duration) * time.Second
	for {
		<-time.After(interval)
		m.updateMetrics()
		m.pollCount++
	}
}

func (m *Monitor) PollCount() int {
	return m.pollCount
}

func (m *Monitor) RandomValue() uint64 {
	return m.randomValue
}

func (m *Monitor) updateMetrics() {
	ms := runtime.MemStats{}
	runtime.ReadMemStats(&ms)
	m.SetMetrics(ms)
}
