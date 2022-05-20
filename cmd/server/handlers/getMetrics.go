package handlers

import (
	"errors"
	"net/http"

	"github.com/achrt/metrics-collector/internal/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func (h *Handler) GetMetrics(c *gin.Context) {

	metric, status, err := h.getMetrics(c)
	if err != nil {
		log.Error(err)
		c.String(status, "err")
		return
	}
	c.JSON(status, metric)
}

func (h *Handler) getMetrics(c *gin.Context) (m *models.Metrics, status int, err error) {
	status = http.StatusOK

	if err = c.ShouldBindJSON(&m); err != nil {
		status = http.StatusBadRequest
		log.Error(err)
		return
	}

	if m.ID == "" || m.MType == "" {
		status = http.StatusNotFound
		err = errors.New("m.ID | m.MType is an empty string")
		log.Error(err)
		return
	}

	m, err = h.store.Get(m.ID)
	if err != nil {
		status = http.StatusNotFound
		log.Error(err)
	}
	return
}
