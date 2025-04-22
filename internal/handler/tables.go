package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tachkovka/internal/utils"
)

func (h *Handler) GetTableByName(c *gin.Context) {
	name := c.Query("name")

	records, err := h.services.GetTable(name)

	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, records)
}
