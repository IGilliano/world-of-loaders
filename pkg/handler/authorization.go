package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
	"worldOfLoaders/pkg/repository/repo_models"
)

const (
	authorizationHeader = "Authorization"
)

type Login struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Class    string `json:"class"`
}

func (h *Handler) register(c *gin.Context) {
	var player repo_models.Player

	if err := c.BindJSON(&player); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.service.Authorization.CreatePlayer(player)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		`id`: id,
	})

}

func (h *Handler) login(c *gin.Context) {
	var input Login

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}

func (h *Handler) getPlayers(c *gin.Context) {
	user, err := h.service.Authorization.GetPlayers()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	playerId, playerClass, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("playerID", playerId)
	c.Set("playerClass", playerClass)
}

func (h *Handler) createTasks(c *gin.Context) {
	numString := c.GetHeader("number")
	num, err := strconv.Atoi(numString)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	tasksId, err := h.service.CreateTasks(num)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	log.Println(tasksId)
	c.JSON(http.StatusOK, "Tasks created")

}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("playerID")
	if !ok {
		return 0, errors.New("user not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}
	return idInt, nil
}
