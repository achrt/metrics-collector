package handlers

import (
	"errors"
	"net/http"

	"github.com/achrt/metrics-collector/internal/domain/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMetrics(c *gin.Context) {

	metric, status, err := h.getMetrics(c)
	if err != nil {
		c.String(status, err.Error())
		return
	}
	c.JSON(status, metric)
}

func (h *Handler) getMetrics(c *gin.Context) (m *models.Metrics, status int, err error) {
	status = http.StatusOK

	if err = c.ShouldBindJSON(&m); err != nil {
		status = http.StatusBadRequest
		return
	}

	if m.ID == "" || m.MType == "" {
		status = http.StatusBadRequest
		err = errors.New("m.ID | m.MType is an empty string")
		return
	}

	m, err = h.store.Get(m.ID)
	if err != nil {
		status = http.StatusBadRequest
	}
	return
}
