package handlers

import (
	"github.com/achrt/metrics-collector/cmd/server/application"
	"github.com/achrt/metrics-collector/internal/domain/repositories"
	"github.com/achrt/metrics-collector/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store  repositories.Storage
	router *gin.Engine
}

func New(app *application.App) *Handler {
	return &Handler{
		store:  app.Store,
		router: app.Router,
	}
}

func (h *Handler) Router() {
	// TODO: проблемы с загрузкой шаблона в автотестах, не находится путь до директории
	// h.router.LoadHTMLGlob("./handlers/templates/*")

	h.router.RedirectFixedPath = false
	h.router.RedirectTrailingSlash = false

	h.router.Use(gin.Logger())
	h.router.Use(gin.Recovery())
	h.router.Use(middlewares.GzipDecompress(), middlewares.GzipCompress())

	h.router.GET("/", h.List)
	h.router.GET("/health", h.Health)
	h.router.GET("/value/:type/:name", h.Get)

	h.router.POST("/value/", h.GetMetrics)
	h.router.POST("/update/", h.UpdateMetrics)
	h.router.POST("/update/:type/:code/:value", h.Update)
}
