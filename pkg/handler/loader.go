package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getLoader(c *gin.Context) {
	playerID, ok := c.Get("playerID")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "ID is empty")
		return
	}

	loader, err := h.service.GetLoader(playerID.(int))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, loader)
}

func (h *Handler) getCompletedTasks(c *gin.Context) {
	playerID, ok := c.Get("playerID")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "ID is empty")
		return
	}

	tasks, err := h.service.GetLoaderTasks(playerID.(int))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, tasks)
}
