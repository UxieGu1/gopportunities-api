package handler

import (
	"net/http"

	"github.com/UxieGu1/gopportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

// CreateOpeningHandler godoc
// @Summary Cria uma nova oportunidade
// @Tags openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Dados da oportunidade"
// @Success 200 {object} schemas.Opening
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /openings [post]
func CreateOpeningHandler(context *gin.Context){
	request := CreateOpeningRequest{}

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

	opening := schemas.Opening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: *request.Remote,
		Link: request.Link,
		Salary: request.Salary,
	}

	if err := openingService.Create(&opening); err != nil {
		logger.Errorf("Error creating opening: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(context, "create-opening", opening)
}
