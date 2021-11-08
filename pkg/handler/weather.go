package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CurrentWeather(c *gin.Context) {
	//Check in DB if we have response today
	//If not, call api and save to DB then return
	//Response to User
	c.JSON(http.StatusOK, "Hello there")
}
