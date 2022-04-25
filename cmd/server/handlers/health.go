package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Health(c *gin.Context) {
	health := struct {
		Health string `json:"health"`
	}{Health: "ok"}

	c.JSON(http.StatusOK, health)
}
