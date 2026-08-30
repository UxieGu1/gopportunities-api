package company

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowCompanyHandler(ctx *gin.Context) {
	idStr := ctx.Query("id")
	if idStr == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "invalid id format")
		return
	}

	response, err := companyService.GetById(uint(id))
	if err != nil {
		sendError(ctx, http.StatusNotFound, "company not found")
		return
	}

	sendSuccess(ctx, "show-company", response)
}