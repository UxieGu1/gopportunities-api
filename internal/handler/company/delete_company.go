package company

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteCompanyHandler(context *gin.Context) {
	idStr := context.Query("id")
	if idStr == "" {
		sendError(context, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		sendError(context, http.StatusBadRequest, "invalid id format")
		return
	}

	if err := companyService.Delete(uint(id)); err != nil {
		sendError(context, http.StatusInternalServerError, "error deleting company")
		return
	}

	sendSuccess(context, "delete-company", nil)
}