package health

import "github.com/gin-gonic/gin"

// RegisterRoutes registers all the routes for health
func (h *Health) RegisterRoutes(ginEngine *gin.Engine) {
	healthRouter := ginEngine.Group("/health")
	{
		healthRouter.GET("/", h.GetHealth)
	}
}
