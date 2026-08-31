package application

import (
	"net/http"

	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func ListApplicationsHandler(ctx *gin.Context) {
	applications, err := applicationService.GetAll()
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing applications")
		return
	}

	responses := make([]schemas.ApplicationResponse, 0, len(applications))
	for _, app := range applications {
		responses = append(responses, schemas.ApplicationResponse{
			ID:          app.ID,
			CandidateID: app.CandidateID,
			OpeningID:   app.OpeningID,
			Status:      app.Status,
			Notes:       app.Notes,
			CreatedAt:   app.CreatedAt,
			UpdatedAt:   app.UpdatedAt,
		})
	}

	sendSuccess(ctx, "list-applications", responses)
}