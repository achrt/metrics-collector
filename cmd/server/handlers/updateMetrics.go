package handlers

import (
	"net/http"

	"github.com/achrt/metrics-collector/internal/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func (h *Handler) UpdateMetrics(c *gin.Context) {
	status, err := h.updateMetrics(c)
	if err != nil {
		log.Print(err)
	}
	c.JSON(status, struct{}{})
}

func (h *Handler) updateMetrics(c *gin.Context) (status int, err error) {
	status = http.StatusOK

	var m models.Metrics
	if err = c.ShouldBindJSON(&m); err != nil {
		status = http.StatusBadRequest
		log.Error(err)
		return
	}

	if err = h.store.Set(m.ID, m); err != nil {
		log.Error(err)
		status = http.StatusBadRequest
	} else {
		if m.Delta != nil {
			log.Infof("[updated] %s, %s, %d", m.ID, m.MType, *m.Delta)
		}
		if m.Value != nil {
			log.Infof("[updated] %s, %s, %v", m.ID, m.MType, *m.Value)
		}
	}

	return
}
