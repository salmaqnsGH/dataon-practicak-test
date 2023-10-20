package handler

import (
	"dataon/company"
	"net/http"

	"github.com/gin-gonic/gin"
)

type companyHandler struct {
	service company.Service
}

func NewCompanyHandler(service company.Service) *companyHandler {
	return &companyHandler{service}
}

func (h *companyHandler) GetCompanies(c *gin.Context) {
	companies, err := h.service.GetCompanies()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, companies)
}
