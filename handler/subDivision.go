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
