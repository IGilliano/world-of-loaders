package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type StartParams struct {
	Task    int   `json:"task"`
	Loaders []int `json:"loaders"`
}

func (h *Handler) getClient(c *gin.Context) {
	playerID, ok := c.Get("playerID")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "ID is empty")
		return
	}

	client, err := h.service.GetClient(playerID.(int))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, client)
}

func (h *Handler) getAvailableTasks(c *gin.Context) {
	tasks, err := h.service.GetClientTasks()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) start(c *gin.Context) {
	var params StartParams
	playerID, ok := c.Get("playerID")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "ID is empty")
		return
	}

	if err := c.BindJSON(&params); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid params")
		return
	}

	log.Printf("Game is starting NOW!")
	result, err := h.service.Start(playerID.(int), params.Task, params.Loaders)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !result {
		c.JSON(http.StatusOK, "You lost! Buy battlepass for another try")
	} else {
		c.JSON(http.StatusOK, "Task completed! Congratulations!")
	}
}
