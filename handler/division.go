package handler

import (
	"dataon/division"
	"net/http"

	"github.com/gin-gonic/gin"
)

type divisionHandler struct {
	service division.Service
}

func NewDivisionHandler(service division.Service) *divisionHandler {
	return &divisionHandler{service}
}

func (h *divisionHandler) GetDivisions(c *gin.Context) {
	divisions, err := h.service.GetDivisions()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, divisions)
}
