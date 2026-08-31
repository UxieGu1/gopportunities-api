package company

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowCompanyHandler(context *gin.Context) {
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

	response, err := companyService.GetByID(uint(id))
	if err != nil {
		sendError(context, http.StatusNotFound, "company not found")
		return
	}

	sendSuccess(context, "show-company", response)
}