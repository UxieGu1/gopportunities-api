package candidate

import (
	"net/http"
	"strconv"

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

	if err := candidateService.Delete(uint(id)); err != nil {
		sendError(context, http.StatusInternalServerError, "error deleting candidate")
		return
	}

	sendSuccess(context, "delete-candidate", nil)
}