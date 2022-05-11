package storage

import (
	"errors"
	"strings"
	"sync"

	"github.com/achrt/metrics-collector/internal/domain/models"
)

type Storage struct {
	cMutex sync.RWMutex
	c      map[string]int64

	uMutex sync.RWMutex
	u      map[string]float64

	mMutex sync.RWMutex
	m      map[string]models.Metrics
}

func New() *Storage {
	return &Storage{
		c: map[string]int64{},
		u: map[string]float64{},
		m: map[string]models.Metrics{},
	}
}

func (str *Storage) Get(code string) (*models.Metrics, error) {
	str.mMutex.RLock()
	defer str.mMutex.RUnlock()

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

// TODO: rename methodes
func (str *Storage) GetMetric(code string) (float64, error) {
	str.uMutex.RLock()
	defer str.uMutex.RUnlock()

	code = strings.ToLower(code)
	if val, ok := str.u[code]; ok {
		return val, nil
	}
	return 0, errors.New("unknown metric code")
}

func (str *Storage) GetMetrics() map[string]float64 {
	return str.u
}

func (str *Storage) GetCounters() map[string]int64 {
	return str.c
}

func (str *Storage) GetCounter(code string) (int64, error) {
	str.cMutex.RLock()
	defer str.cMutex.RUnlock()

	code = strings.ToLower(code)
	if val, ok := str.c[code]; ok {
		return val, nil
	}
	return 0, errors.New("unknown counter code")
}

func (str *Storage) UpdateMetric(code string, val float64) error {
	return str.updateMetric(code, val)
}

func (str *Storage) UpdateCounter(code string, val int64) error {
	return str.updateCounter(code, val)
}

func (str *Storage) updateCounter(code string, val int64) error {
	if code == "" {
		return errors.New("code is an empty string")
	}

	str.cMutex.RLock()
	defer str.cMutex.RUnlock()

	code = strings.ToLower(code)
	str.c[code] += val
	return nil
}

func (str *Storage) updateMetric(code string, val float64) error {
	if code == "" {
		return errors.New("code is an empty string")
	}

	str.uMutex.RLock()
	defer str.uMutex.RUnlock()

	code = strings.ToLower(code)
	str.u[code] = val
	return nil
}
