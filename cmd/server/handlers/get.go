package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/achrt/metrics-collector/internal/domain/models"
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

	if mType != models.TypeCounter && mType != models.TypeGauge {
		status = http.StatusNotFound
		err = errors.New(http.StatusText(status))
		return
	}

	m, err := h.store.Get(code)
	if err != nil {
		status = http.StatusNotFound
		return
	}

	if mType == models.TypeCounter && m.Delta != nil {
		value = fmt.Sprintf("%v", *m.Delta)
		return
	}

	if m.Value != nil {
		value = fmt.Sprintf("%v", *m.Value)
	}

	return
}
