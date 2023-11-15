package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getLoader(c *gin.Context) {
	loader, err := h.service.GetLoader(c.GetInt("playerID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, loader)
}

func (h *Handler) getCompletedTasks(c *gin.Context) {
	tasks, err := h.service.GetLoaderTasks(c.GetInt("playerID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, tasks)
}
