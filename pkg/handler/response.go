package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	err := errors.New(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{err.Error()})
}
