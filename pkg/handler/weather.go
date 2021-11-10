package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CurrentWeather(c *gin.Context) {
	//Check in DB if we have response today
	//If not, call api and save to DB then return
	//Response to User
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	weather, err := h.services.Weather.Get(userId, "Kiev")

	fmt.Println(weather)

	c.JSON(http.StatusOK, "Hello there")
}
