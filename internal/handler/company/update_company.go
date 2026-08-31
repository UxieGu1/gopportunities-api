package company

import (
	"net/http"
	"strconv"

	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateCompanyHandler(context *gin.Context) {
	var request UpdateCompanyRequest

	if err := context.BindJSON(&request); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	idStr := context.Query("id")
	if idStr == "" {
		sendError(context, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		sendError(context, http.StatusBadRequest, "invalid id format")
		return
	}

	companyData := schemas.Company{
		Name:        request.Name,
		Description: request.Description,
		Website:     request.Website,
		Email:       request.Email,
		Location:    request.Location,
	}

	response, err := companyService.Update(uint(id), &companyData)
	if err != nil {
		sendError(context, http.StatusInternalServerError, "error updating company")
		return
	}

	sendSuccess(context, "update-company", response)
}