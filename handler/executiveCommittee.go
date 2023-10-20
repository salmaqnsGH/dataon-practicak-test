package handler

import (
	"dataon/executiveCommittee"
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
