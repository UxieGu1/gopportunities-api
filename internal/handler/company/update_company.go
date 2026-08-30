package company

import (
	"net/http"
	"strconv"

	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateCompanyHandler(ctx *gin.Context) {
	var request UpdateCompanyRequest

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	idStr := ctx.Query("id")
	if idStr == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "invalid id format")
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
		sendError(ctx, http.StatusInternalServerError, "error updating company")
		return
	}

	sendSuccess(ctx, "update-company", response)
}