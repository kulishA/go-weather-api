package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type searchQuery struct {
	City string `form:"city"`
}

func (h *Handler) Search(c *gin.Context) {
	var input searchQuery

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

func (h *Handler) Save(c *gin.Context) {
	c.JSON(http.StatusOK, "Save")
}

func (h *Handler) GetSaved(c *gin.Context) {
	c.JSON(http.StatusOK, "GetSaved")
}
