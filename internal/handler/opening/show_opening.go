package opening

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowOpeningHandler godoc
// @Summary Busca uma oportunidade por id
// @Tags openings
// @Produce json
// @Param id path int true "ID da oportunidade"
// @Success 200 {object} schemas.Opening
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /openings/{id} [get]
func ShowOpeningHandler(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		sendError(context, http.StatusBadRequest, errParamIsRequired("id", "pathParameter").Error())
		return
	}
	opening, err := openingService.GetByID(id)
	if err != nil {
		sendError(context, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(context, "show-opening", opening)
}
