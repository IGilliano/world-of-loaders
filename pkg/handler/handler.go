package handler

import (
	"github.com/gin-gonic/gin"
	"worldOfLoaders/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/tasks", h.createTasks)

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)
		auth.GET("/players", h.getPlayers)
	}

	loaders := router.Group("/loaders", h.validateLoader)

	{
		loaders.GET("/me", h.getLoader)
		loaders.GET("/tasks", h.getCompletedTasks)
	}

	clients := router.Group("/clients", h.validateClient)
	{
		clients.GET("/me", h.getClient)
		clients.GET("/tasks", h.getAvailableTasks)
		clients.POST("/start", h.start)
	}

	return router
}
