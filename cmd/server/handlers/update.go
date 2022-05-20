package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/achrt/metrics-collector/internal/domain/models"
	"github.com/gin-gonic/gin"

	"github.com/labstack/gommon/log"
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

	if mType != models.TypeCounter && mType != models.TypeGauge {
		status = http.StatusNotImplemented
		err = errors.New(http.StatusText(status))
		log.Error(err)
		return
	}

	if mType == models.TypeCounter {
		var value int64
		value, err = strconv.ParseInt(rawValue, 10, 64)
		if err != nil {
			status = http.StatusBadRequest
			log.Error(err)
			return
		}
		m := models.Metrics{
			ID:    code,
			MType: mType,
			Delta: &value,
		}
		if err = h.store.Set(code, m); err != nil {
			status = http.StatusInternalServerError
			log.Error(err)
		} else {
			log.Infof("[updated] %s, %s, %d", m.ID, m.MType, *m.Delta)
		}
		return
	}

	var value float64
	value, err = strconv.ParseFloat(rawValue, 64)
	if err != nil {
		status = http.StatusBadRequest
		log.Error(err)
		return
	}

	m := models.Metrics{
		ID:    code,
		MType: mType,
		Value: &value,
	}

	if err = h.store.Set(code, m); err != nil {
		status = http.StatusInternalServerError
		log.Error(err)
	} else {
		log.Infof("[updated] %s, %s, %d", m.ID, m.MType, *m.Value)
	}
	return
}
