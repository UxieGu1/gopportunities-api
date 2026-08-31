package candidate

import (
	"net/http"

	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func CreateCandidateHandler(context *gin.Context) {
	var request CreateCandidateRequest

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

	candidate := schemas.Candidate{
		UserID:    request.UserID,
		Name:      request.Name,
		Linkedin:  request.Linkedin,
		ResumeURL: request.ResumeURL,
		Skills:    request.Skills,
	}

	if err := candidateService.Create(&candidate); err != nil {
		logger.Errorf("error creating candidate: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "error creating candidate on database")
		return
	}

	response := schemas.CandidateResponse{
		ID:        candidate.ID,
		UserID:    candidate.UserID,
		Name:      candidate.Name,
		Linkedin:  candidate.Linkedin,
		ResumeURL: candidate.ResumeURL,
		Skills:    candidate.Skills,
		CreatedAt: candidate.CreatedAt,
		UpdatedAt: candidate.UpdatedAt,
	}

	sendSuccess(context, "create-candidate", response)
}