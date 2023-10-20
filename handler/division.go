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

func (h *divisionHandler) CreateDivision(c *gin.Context) {
	var executiveCommitteeID division.ExecutiveCommitteeIDInput
	err := c.ShouldBindUri(&executiveCommitteeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var input division.CreateDivisionInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	newDivision, err := h.service.CreateDivision(input, executiveCommitteeID.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, newDivision)
}
