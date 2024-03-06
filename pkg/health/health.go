package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Health ...
type Health struct {
}

// NewHealth returns a new instance of Health
func NewHealth() *Health {
	return &Health{}
}

// GetHealth return health of the server
func (h *Health) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
