package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/achrt/metrics-collector/internal/domain/models/health"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(c *gin.Context) {
	status, err := h.update(c)
	if err != nil {
		c.String(status, err.Error())
		return
	}
	c.JSON(status, nil)
}

func (h *Handler) update(c *gin.Context) (status int, err error) {
	status = http.StatusOK

	mType := c.Param("type")
	code := c.Param("code")
	rawValue := c.Param("value")

	if mType != health.TypeCounter && mType != health.TypeGauge {
		status = http.StatusNotImplemented
		err = errors.New(http.StatusText(status))
		return
	}

	if mType == health.TypeCounter {
		var value int64
		value, err = strconv.ParseInt(rawValue, 10, 64)
		if err != nil {
			status = http.StatusBadRequest
			return
		}
		err = h.store.UpdateCounter(code, value)
		return
	}

	var value float64
	value, err = strconv.ParseFloat(rawValue, 64)
	if err != nil {
		status = http.StatusBadRequest
		return
	}

	if err = h.store.UpdateMetric(code, value); err != nil {
		status = http.StatusInternalServerError
		return
	}
	return
}
