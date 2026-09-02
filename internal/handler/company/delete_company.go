package company

import (
	"net/http"
	"strconv"

	"github.com/UxieGu1/gopportunities-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func DeleteCompanyHandler(context *gin.Context) {
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

	var deleteErr error
	if middleware.GetUserRole(context) == "ADMIN" {
		deleteErr = companyService.Delete(uint(id))
	} else {
		deleteErr = companyService.DeleteForUser(uint(id), middleware.GetUserID(context))
	}
	if deleteErr != nil {
		sendError(context, http.StatusNotFound, deleteErr.Error())
		return
	}

	sendSuccess(context, "delete-company", nil)
}
