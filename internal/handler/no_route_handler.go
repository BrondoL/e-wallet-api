package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) NoRoute(c *gin.Context) {
	response := utils.ErrorResponse("page not found", http.StatusNotFound, "page not found")
	c.JSON(http.StatusNotFound, response)
}
