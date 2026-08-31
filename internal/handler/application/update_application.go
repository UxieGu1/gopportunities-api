package application

import (
	"net/http"
	"strconv"

	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateApplicationHandler(ctx *gin.Context) {
	var request UpdateApplicationRequest

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

	idStr := ctx.Param("id")
	if idStr == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "pathParameter").Error())
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "invalid id format")
		return
	}

	application, err := applicationService.GetByID(uint(id))
	if err != nil {
		sendError(ctx, http.StatusNotFound, "application not found")
		return
	}

	if request.Status != "" {
		application.Status = request.Status
	}
	if request.Notes != "" {
		application.Notes = request.Notes
	}

	if err := applicationService.Update(application); err != nil {
		logger.Errorf("error updating application: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating application")
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

	sendSuccess(ctx, "update-application", response)
}