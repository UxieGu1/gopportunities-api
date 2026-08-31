package application

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteApplicationHandler(ctx *gin.Context) {
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

	if err := applicationService.Delete(uint(id)); err != nil {
		sendError(ctx, http.StatusInternalServerError, "error deleting application")
		return
	}

	sendSuccess(ctx, "delete-application", nil)
}