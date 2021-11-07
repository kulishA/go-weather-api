package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kulishA/go-weather-api/pkg/service"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		Services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sing-up", h.SingUp)
		auth.POST("sing-in", h.SingIn)
	}

	return router
}
