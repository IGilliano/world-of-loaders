package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Error string `json:"error"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	err := errors.New(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{err.Error()})
}
