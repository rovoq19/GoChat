package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rovoq19/GoChat/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())

	chats := router.Group("/chat")
	{
		chats.POST("/", h.createChat)
		chats.GET("/", h.getAllChats)
		chats.GET("/:id", h.getChatById)
		chats.PUT("/:id", h.updateChat)
		chats.DELETE("/:id", h.deleteChat)
	}
	router.POST("/message", h.SendMessage)

	return router
}
