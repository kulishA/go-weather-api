package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kulishA/go-weather-api/domains"
	"github.com/kulishA/go-weather-api/helper"
	"net/http"
)

func (h *Handler) SingUp(c *gin.Context) {
	var input domains.User

	if err := c.BindJSON(&input); err != nil{
		helper.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.Authorization.CreateUser(input)

	if err != nil {
		helper.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SingIn(c *gin.Context) {}
