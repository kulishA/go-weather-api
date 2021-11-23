package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type getInput struct {
	CityId int `form:"city_id"`
}

func (h *Handler) GetById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	var input getInput

	if err := c.Bind(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	weather, err := h.services.Weather.Get(userId, input.CityId)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": weather,
	})
}

func (h *Handler) GetAll(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	weather, err := h.services.Weather.GetAll(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": weather,
	})
}
