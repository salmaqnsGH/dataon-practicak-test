package handler

import (
	"dataon/company"
	"net/http"
	"strconv"

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

func (h *companyHandler) CreateCompany(c *gin.Context) {
	var input company.CreateCompanyInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	newCompany, err := h.service.CreateCompany(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, newCompany)
}

func (h *companyHandler) GetCompanyByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	subDivisions, err := h.service.GetCompanyByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, subDivisions)
}
