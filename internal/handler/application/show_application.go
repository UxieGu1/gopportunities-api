package application

import (
	"net/http"
	"strconv"

	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func ShowApplicationHandler(ctx *gin.Context) {
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

	response := schemas.ApplicationResponse{
		ID:          application.ID,
		CandidateID: application.CandidateID,
		OpeningID:   application.OpeningID,
		Status:      application.Status,
		Notes:       application.Notes,
		CreatedAt:   application.CreatedAt,
		UpdatedAt:   application.UpdatedAt,
	}

	sendSuccess(ctx, "show-application", response)
}