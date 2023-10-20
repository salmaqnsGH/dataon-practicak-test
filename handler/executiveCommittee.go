package handler

import (
	"dataon/executiveCommittee"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type executiveCommitteeHandler struct {
	service executiveCommittee.Service
}

func NewExecutiveCommitteeHandler(service executiveCommittee.Service) *executiveCommitteeHandler {
	return &executiveCommitteeHandler{service}
}

func (h *executiveCommitteeHandler) GetExecutiveCommitties(c *gin.Context) {
	executiveCommitties, err := h.service.GetExecutivecommitties()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, executiveCommitties)
}

func (h *executiveCommitteeHandler) CreateExecutiveCommittee(c *gin.Context) {
	var companyID executiveCommittee.CompanyIDInput
	err := c.ShouldBindUri(&companyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var input executiveCommittee.CreateExecutiveCommitteeInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	fmt.Println(input)

	newExecutiveCommittee, err := h.service.CreateExecutivecommittee(input, companyID.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, newExecutiveCommittee)
}
