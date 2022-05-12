package storage

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/achrt/metrics-collector/internal/domain/models"
	"github.com/achrt/metrics-collector/internal/domain/models/health"
)

type Storage struct {
	mMutex *sync.RWMutex
	m      map[string]models.Metrics
}

func New() *Storage {
	return &Storage{
		m:      map[string]models.Metrics{},
		mMutex: &sync.RWMutex{},
	}
}

func (str *Storage) Init() {
	v := float64(0.0)
	d := int64(0)
	h := health.HealthStat{}
	for _, code := range h.MetricCodes() {
		t, _ := h.GetType(code)
		m := models.Metrics{
			ID: code,
			MType: t,
		}
		if t == models.TypeCounter {
			m.Delta = &d
		}
		if t == models.TypeGauge {
			m.Value = &v
		}

		str.m[code] = m
	}
}

func (str *Storage) Get(code string) (*models.Metrics, error) {
	code = strings.ToLower(code)
	if val, ok := str.m[code]; ok {
		return &val, nil
	}
	return nil, errors.New("unknown metric code")
}

func (str *Storage) Set(code string, val models.Metrics) error {
	if code == "" {
		return errors.New("code is an empty string")
	}

	str.mMutex.RLock()
	defer str.mMutex.RUnlock()

	code = strings.ToLower(code)

	if val.MType == models.TypeCounter {
		if val.Delta == nil {
			return errors.New("val.Delta can not be nil")
		}
		if str.m[code].Delta == nil {
			str.m[code] = val
			return nil
		}
		*str.m[code].Delta += *val.Delta
	}

	if val.MType == models.TypeGauge {
		if val.Value == nil {
			return errors.New("val.Value can not be nil")
		}
		str.m[code] = val
	}
	return nil
}

func (s Storage) PrintMetrics() map[string]string {
	res := map[string]string{}
	for code, val := range s.m {
		if val.Delta != nil {
			res[code] = fmt.Sprintf("%v", *val.Delta)
		}
		if val.Value != nil {
			res[code] = fmt.Sprintf("%v", *val.Value)
		}
	}
	return res
}
