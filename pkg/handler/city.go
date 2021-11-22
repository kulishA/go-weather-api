package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type searchInput struct {
	City string `form:"city"`
}

func (h *Handler) Search(c *gin.Context) {
	var input searchInput

	if err := c.Bind(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	search, err := h.services.City.Search(input.City)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": search,
	})
}

type saveInput struct {
	CityId int `json:"city_id"`
}

func (h *Handler) Save(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input saveInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := h.services.City.Save(userId, input.CityId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}

func (h *Handler) GetSaved(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	result, err := h.services.City.Get(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}
