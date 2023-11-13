package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) getClient(c *gin.Context) {

}

func (h *Handler) getAvailableTasks(c *gin.Context) {

}

//В хендлере только валидация
/*
func (h *Handler) start(c *gin.Context) {
	task, err := h.service.ChooseTask
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	loaders, err := h.service.GetLoaders
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	result, err := h.service.CompleteTask(task, loaders)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Complete": task.Name
	})
}
*/
