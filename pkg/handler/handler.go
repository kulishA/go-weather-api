package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kulishA/go-weather-api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sing-up", h.SingUp)
		auth.POST("sing-in", h.SingIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		weather := api.Group("/weather")
		{
			weather.GET("/", h.GetById)
			weather.GET("/all", h.GetAll)
		}

		city := api.Group("/city")
		{
			city.GET("/search", h.Search)
			city.POST("/", h.Save)
			city.GET("/", h.GetSaved)
		}
	}

	return router
}
