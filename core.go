package core

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ready/core/pkg/health"
)

// CoreHandler should be used to construct the application
type CoreHandler struct {
	health *health.Health
}

// NewCoreHandler returns a new instance of CoreHandler
func NewCoreHandler(health *health.Health) *CoreHandler {
	return &CoreHandler{health}
}

// RegisterHandlers registers all the handlers
func (h *CoreHandler) RegisterHandlers(ginEngine *gin.Engine) {
	h.health.RegisterRoutes(ginEngine)
}
