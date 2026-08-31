package candidate

import (
	"net/http"
	"strconv"

	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func ShowCandidateHandler(context *gin.Context) {
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

	candidate, err := candidateService.GetByID(uint(id))
	if err != nil {
		sendError(context, http.StatusNotFound, "candidate not found")
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

	sendSuccess(context, "show-candidate", response)
}