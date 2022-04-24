package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/achrt/metrics-collector/internal/domain/models/health"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
		return
	}

	status, err := h.update(r)
	if err != nil {
		http.Error(w, err.Error(), status)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
}

func (h *Handler) update(r *http.Request) (status int, err error) {
	status = http.StatusOK

	params := strings.Split(r.URL.Path, "/")

	if len(params) < 5 {
		err = errors.New("wrong request")
		status = http.StatusNotFound
		return
	}

	if params[2] != health.TypeGauge && params[2] != health.TypeCounter {
		err = errors.New("wrong metric type")
		status = http.StatusNotImplemented
		return
	}

	if params[2] == health.TypeCounter && params[3] == health.PollCount {
		var value int64
		value, err = strconv.ParseInt(params[4], 10, 64)
		if err != nil {
			status = http.StatusBadRequest
			return
		}
		h.store.UpdateCounter(value)
		return
	}

	var value float64
	value, err = strconv.ParseFloat(params[4], 64)
	if err != nil {
		status = http.StatusBadRequest
		return
	}

	// TODO: не понятно, нужно ли сохранять неопознанную метрику
	if !health.IsExists(params[3]) {
		return
	}

	if err = h.store.UpdateMetric(params[3], value); err != nil {
		status = http.StatusInternalServerError
		return
	}
	return
}
