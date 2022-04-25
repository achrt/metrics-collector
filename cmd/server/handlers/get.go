package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/achrt/metrics-collector/internal/domain/models/health"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Get(c *gin.Context) {

	value, status, err := h.get(c)
	if err != nil {
		c.String(status, err.Error())
		return
	}
	c.String(status, value)
}

func (h *Handler) get(c *gin.Context) (value string, status int, err error) {
	mType := c.Param("type")
	code := c.Param("name")

	status = http.StatusOK

	if mType != health.TypeCounter && mType != health.TypeGauge {
		status = http.StatusNotFound
		err = errors.New(http.StatusText(status))
		return
	}

	if mType == health.TypeCounter {
		var val int64
		val, err = h.store.GetCounter(code)
		if err != nil {
			status = http.StatusNotFound
			err = errors.New(http.StatusText(status))
			return
		}
		value = fmt.Sprintf("%v", val)
		return
	}

	var val float64
	val, err = h.store.GetMetric(code)
	if err != nil {
		status = http.StatusNotFound
		err = errors.New(http.StatusText(status))
		return
	}
	value = fmt.Sprintf("%v", val)
	return
}
