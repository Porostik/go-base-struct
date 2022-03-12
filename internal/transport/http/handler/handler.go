package handler

import (
	"architecture/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	userHandler := NewUserHandler(h.services.User)

	api := router.Group("/api")
	{
		userHandler.Init(api)
	}
}
