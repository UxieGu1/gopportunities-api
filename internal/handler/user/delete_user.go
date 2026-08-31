package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(context *gin.Context) {
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

	if err := userService.Delete(uint(id)); err != nil {
		sendError(context, http.StatusInternalServerError, "error deleting user")
		return
	}

	sendSuccess(context, "delete-user", nil)
}