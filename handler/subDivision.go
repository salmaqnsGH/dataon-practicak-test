package handler

import (
	subdivision "dataon/subdivision"
	"net/http"

	"github.com/gin-gonic/gin"
)

type subDivisionHandler struct {
	service subdivision.Service
}

func NewSubDivisionHandler(service subdivision.Service) *subDivisionHandler {
	return &subDivisionHandler{service}
}

func (h *subDivisionHandler) GetSubDivisions(c *gin.Context) {
	subDivisions, err := h.service.GetSubDivisions()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, subDivisions)
}

func (h *subDivisionHandler) CreateSubDivision(c *gin.Context) {
	var divisionID subdivision.DivisionIDInput
	err := c.ShouldBindUri(&divisionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var input subdivision.CreateSubDivisionInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	newDivision, err := h.service.CreateSubDivision(input, divisionID.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, newDivision)
}

func (h *subDivisionHandler) DeleteSubDivision(c *gin.Context) {
	var subDivisionID subdivision.SubDivisionIDInput
	err := c.ShouldBindUri(&subDivisionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.DeleteSubDivision(subDivisionID.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Subdivision is deleted successfully")
}
