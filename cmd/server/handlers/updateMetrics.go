package handlers

import (
	"net/http"

	"github.com/achrt/metrics-collector/internal/domain/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateMetrics(c *gin.Context) {
	status, err := h.updateMetrics(c)
	if err != nil {
		c.String(status, err.Error())
		return
	}
	c.JSON(status, nil)
}

func (h *Handler) updateMetrics(c *gin.Context) (status int, err error) {
	status = http.StatusOK

	var m models.Metrics
	if err = c.ShouldBindJSON(&m); err != nil {
		status = http.StatusBadRequest
		return
	}

	h.store.Set(m.ID, m)

	return status, nil
}
