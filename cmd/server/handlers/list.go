package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) List(c *gin.Context) {
	// TODO: тут не особо красиво принтуется html
	// нужно разобраться с использованием шаблонов в gin
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"list": h.list(c),
	})
}

func (h *Handler) list(c *gin.Context) string {
	counters := h.store.GetCounters()
	metrics := h.store.GetMetrics()
	list := []string{}

	for code, val := range counters {
		list = append(list, fmt.Sprintf("%s: %d", code, val))
	}

	for code, val := range metrics {
		list = append(list, fmt.Sprintf("%s: %v", code, val))
	}

	return strings.Join(list, "\n")
}
