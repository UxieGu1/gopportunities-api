package candidate

import (
	"net/http"
	"strconv"

	"github.com/UxieGu1/gopportunities-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func DeleteCandidateHandler(context *gin.Context) {
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
	if middleware.GetUserRole(context) != "ADMIN" {
		candidate, err := candidateService.GetByIDForUser(uint(id), userID)
		if err != nil {
			sendError(context, http.StatusNotFound, "candidate not found")
			return
		}
		_ = candidate
	}

	if err := candidateService.Delete(uint(id)); err != nil {
		sendError(context, http.StatusInternalServerError, "error deleting candidate")
		return
	}

	sendSuccess(context, "delete-candidate", nil)
}
