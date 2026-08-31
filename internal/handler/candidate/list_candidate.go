package candidate

import (
	"net/http"

	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func ListCandidatesHandler(context *gin.Context) {
	candidates, err := candidateService.GetAll()
	if err != nil {
		sendError(context, http.StatusInternalServerError, "error listing candidates")
		return
	}

	responses := make([]schemas.CandidateResponse, 0, len(candidates))
	for _, c := range candidates {
		responses = append(responses, schemas.CandidateResponse{
			ID:        c.ID,
			UserID:    c.UserID,
			Name:      c.Name,
			Linkedin:  c.Linkedin,
			ResumeURL: c.ResumeURL,
			Skills:    c.Skills,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}

	sendSuccess(context, "list-candidates", responses)
}