package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	weather, err := h.services.Weather.Get(userId, "Kiev")

	fmt.Println(weather)

	c.JSON(http.StatusOK, "Hello there")
}

func (h *Handler) GetAll(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	weather, err := h.services.Weather.Get(userId, "Kiev")

	fmt.Println(weather)

	c.JSON(http.StatusOK, "Hello there")
}
