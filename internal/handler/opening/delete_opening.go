package opening

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteOpeningHandler godoc
// @Summary Remove uma oportunidade por id
// @Tags openings
// @Produce json
// @Param id path int true "ID da oportunidade"
// @Success 200 {object} DeleteOpeningResponse	
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /openings/{id} [delete]
func DeleteOpeningHandler(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		sendError(context, http.StatusBadRequest, errParamIsRequired("id", "pathParameter").Error())
		return
	}

	opening, err := openingService.GetByID(id)
	if err != nil {
		sendError(context, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if err := openingService.Delete(&opening); err != nil {
		sendError(context, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	sendSuccess(context, "delete-opening", opening)
}
