package storage

import (
	"errors"
	"fmt"
	"strings"
)

type Storage struct {
	c map[string]int64
	u map[string]float64
	i map[string]interface{}
}

func New() *Storage {
	return &Storage{
		c: map[string]int64{},
		u: map[string]float64{},
	}
}

func (str *Storage) Get(code string) (string, error) {
	code = strings.ToLower(code)
	if val, ok := str.i[code]; ok {
		return fmt.Sprintf("%v", val), nil
	}
	return "", errors.New("unknown metric code")
}

func (str *Storage) Set(code string, val interface{}) {
	code = strings.ToLower(code)
	str.i[code] = val
}

func (str *Storage) GetMetric(code string) (float64, error) {
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
	if val, ok := str.c[code]; ok {
		return val, nil
	}
	return 0, errors.New("unknown counter code")
}

func (str *Storage) UpdateMetric(code string, val float64) error {
	return str.updateMetric(code, val)
}

func (str *Storage) UpdateCounter(code string, val int64) {
	str.updateCounter(code, val)
}

func (str *Storage) updateCounter(code string, val int64) {
	code = strings.ToLower(code)
	str.c[code] = val
}

func (str *Storage) updateMetric(code string, val float64) error {
	code = strings.ToLower(code)
	str.u[code] = val
	return nil
}
