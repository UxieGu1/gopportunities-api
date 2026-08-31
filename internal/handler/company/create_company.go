package company

import (
	"net/http"

	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func CreateCompanyHandler(context *gin.Context) {
	var request CreateCompanyRequest

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

	company := schemas.Company{
		Name:        request.Name,
		Description: request.Description,
		Website:     request.Website,
		Email:       request.Email,
		Location:    request.Location,
	}

	response, err := companyService.Create(&company)
	if err != nil {
		logger.Errorf("error creating company: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "error creating company on database")
		return
	}

	sendSuccess(context, "create-company", response)
}