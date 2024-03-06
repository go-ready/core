package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type errorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func respond(c *gin.Context, status int, err error) {
	c.JSON(status, errorResponse{
		Status: status,
		Error:  err.Error(),
	})
}

// Internal sends internal server error
func Internal(c *gin.Context, err error) {
	respond(c, http.StatusInternalServerError, err)
}

// BadRequest sends bad request error
func BadRequest(c *gin.Context, err error) {
	respond(c, http.StatusBadRequest, err)
}

// NotFound sends not found error
func NotFound(c *gin.Context, err error) {
	respond(c, http.StatusNotFound, err)
}

// Forbidden sends not found error
func Forbidden(c *gin.Context, err error) {
	respond(c, http.StatusForbidden, err)
}

// UnAuthorized sends not found error
func UnAuthorized(c *gin.Context, err error) {
	respond(c, http.StatusUnauthorized, err)
}
