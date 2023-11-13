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

	auth := router.Group("/")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)
		auth.POST("/tasks", h.createTasks)
		auth.GET("/players", h.getPlayers)
	}

	loaders := router.Group("/loaders", h.userIdentity)

	{
		loaders.GET("/me", h.getLoader)
		loaders.GET("/tasks", h.getCompletedTasks)
	}

	clients := router.Group("/clients", h.userIdentity)
	{
		clients.GET("/me", h.getClient)
		clients.GET("/tasks", h.getAvailableTasks)
		//clients.POST("/start", h.start)
	}
	return router
}
