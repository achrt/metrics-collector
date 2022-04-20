package metrics

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

type Monitor struct {
	runtime.MemStats

	pollCount   int
	randomValue uint64

	duration int64
}

func New(duration int64) *Monitor {
	return &Monitor{duration: duration}
}

func (m *Monitor) Run(ctx context.Context, cancel context.CancelFunc) {
	defer fmt.Println("Canceled Monitor")
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
	runtime.ReadMemStats(&m.MemStats)
}
