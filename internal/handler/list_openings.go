package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListOpeningsHandler godoc
// @Summary Lista todas as oportunidades
// @Tags openings
// @Produce json
// @Success 200 {array} schemas.Opening
// @Failure 500 {object} map[string]interface{}
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings, err := openingService.List()
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
