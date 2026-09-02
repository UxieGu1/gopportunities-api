package application

import (
	"net/http"

	"github.com/UxieGu1/gopportunities-api/internal/middleware"
	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func CreateApplicationHandler(ctx *gin.Context) {
	var request CreateApplicationRequest

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

	candidate, err := candidateService.GetByUserID(middleware.GetUserID(ctx))
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "crie seu perfil de candidato antes de se candidatar")
		return
	}

	application := schemas.Application{
		CandidateID: candidate.ID,
		OpeningID:   request.OpeningID,
		Notes:       request.Notes,
	}

	if err := applicationService.Create(&application); err != nil {
		logger.Errorf("error creating application: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating application on database")
		return
	}

	response := schemas.ApplicationResponse{
		ID:          application.ID,
		CandidateID: application.CandidateID,
		OpeningID:   application.OpeningID,
		Status:      application.Status,
		Notes:       application.Notes,
		CreatedAt:   application.CreatedAt,
		UpdatedAt:   application.UpdatedAt,
	}

	sendSuccess(ctx, "create-application", response)
}
