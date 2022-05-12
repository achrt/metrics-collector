package handlers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) List(c *gin.Context) {
	// TODO: есть проблемы с загрузкой шаблона во время билда в автотестах
	// времено возврат html без использования фреймфорка

	// c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 	"list": h.list(c),
	// })
	
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(c.Writer, h.list(c))
}

func (h *Handler) list(c *gin.Context) string {
	metrics := h.store.PrintMetrics()
	list := []string{}

	for code, val := range metrics {
		list = append(list, fmt.Sprintf("<li>%s: %s</li>", code, val))
	}

	return fmt.Sprintf("<html><body><ul>%s</ul></body></html>", strings.Join(list, "\n"))
}
