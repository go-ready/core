package core

import "github.com/gin-gonic/gin"

type Handler interface {
	RegisterRoutes(*gin.Engine)
}
