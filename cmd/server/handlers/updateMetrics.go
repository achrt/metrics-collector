package handlers

import (
	"log"
	"net/http"

	"github.com/achrt/metrics-collector/internal/domain/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateMetrics(c *gin.Context) {
	status, err := h.updateMetrics(c)
	if err != nil {
		log.Println(err)
	}
	c.JSON(status, struct{}{})
}

func (h *Handler) updateMetrics(c *gin.Context) (status int, err error) {
	status = http.StatusOK

	var m models.Metrics
	if err = c.ShouldBindJSON(&m); err != nil {
		status = http.StatusBadRequest
		return
	}

	if err = h.store.Set(m.ID, m); err != nil {
		status = http.StatusBadRequest
	}

	return
}
