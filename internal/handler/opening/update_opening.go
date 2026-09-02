package opening

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateOpeningHandler godoc
// @Summary Atualiza uma oportunidade
// @Tags openings
// @Accept json
// @Produce json
// @Param id path int true "ID da oportunidade"
// @Param request body UpdateOpeningRequest true "Dados atualizados da oportunidade"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /openings/{id} [put]
func UpdateOpeningHandler(context *gin.Context) {
	request := UpdateOpeningRequest{}

	if err := context.BindJSON(&request); err != nil {
		logger.Errorf("json binding error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

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

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.CompanyID > 0 {
		opening.CompanyID = request.CompanyID
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := openingService.Update(&opening); err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "error updating opening")
		return
	}
	sendSuccess(context, "update-opening", opening)
}
