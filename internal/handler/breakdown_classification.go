package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tachkovka/internal/utils"
)

func (h *Handler) GetClassification(c *gin.Context) {
	classifications, err := h.services.GetClassifications()

	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, classifications)
}
