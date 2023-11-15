package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"worldOfLoaders/pkg/models"
	"worldOfLoaders/pkg/service"
)

type ClientInfo struct {
	Client  models.Client
	Loaders []*models.Loader
}

func (h *Handler) getClient(c *gin.Context) {
	client, loaders, err := h.service.GetClientInfo(c.GetInt("playerID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, ClientInfo{Client: client, Loaders: loaders})
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
	var params service.StartParams

	if err := c.BindJSON(&params); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid params")
		return
	}

	log.Printf("Game is starting NOW!")
	result, err := h.service.Start(c.GetInt("playerID"), params)
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
