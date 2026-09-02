package opening

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListOpeningsHandler godoc
// @Summary Lista todas as oportunidades
// @Tags openings
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} map[string]interface{}
// @Router /openings [get]
func ListOpeningsHandler(context *gin.Context) {
	openings, err := openingService.List()
	if err != nil {
		sendError(context, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(context, "list-openings", openings)
}
