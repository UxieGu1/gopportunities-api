package candidate

import (
	"net/http"
	"strconv"

	"github.com/UxieGu1/gopportunities-api/internal/middleware"
	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateCandidateHandler(context *gin.Context) {
	var request UpdateCandidateRequest

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

	idStr := context.Param("id")
	if idStr == "" {
		sendError(context, http.StatusBadRequest, errParamIsRequired("id", "pathParameter").Error())
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		sendError(context, http.StatusBadRequest, "invalid id format")
		return
	}

	userID := middleware.GetUserID(context)
	if middleware.GetUserRole(context) == "ADMIN" {
		userID = 0
	}
	candidate, err := candidateService.GetByID(uint(id))
	if userID != 0 {
		candidate, err = candidateService.GetByIDForUser(uint(id), userID)
	}
	if err != nil {
		sendError(context, http.StatusNotFound, "candidate not found")
		return
	}

	if request.Name != "" {
		candidate.Name = request.Name
	}
	if request.Linkedin != "" {
		candidate.Linkedin = request.Linkedin
	}
	if request.ResumeURL != "" {
		candidate.ResumeURL = request.ResumeURL
	}
	if request.Skills != "" {
		candidate.Skills = request.Skills
	}

	if err := candidateService.Update(candidate); err != nil {
		logger.Errorf("error updating candidate: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "error updating candidate")
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

	sendSuccess(context, "update-candidate", response)
}
