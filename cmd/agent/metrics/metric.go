package metrics

import (
	"context"
	"runtime"
	"time"

	"github.com/achrt/metrics-collector/internal/domain/models/health"
)

type Monitor struct {
	health.HealthStat

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
		m.updateRandom()
		m.PollCount++
	}
}

func (m *Monitor) updateMetrics() {
	ms := runtime.MemStats{}
	runtime.ReadMemStats(&ms)
	m.SetMetrics(ms)
}

func (m *Monitor) updateRandom() {
	// TODO: не понятно, зачем это поле, пока пусть будет UnixMicro
	m.RandomValue = float64(time.Now().UnixMicro())
}
