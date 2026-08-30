package company

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteCompanyHandler(ctx *gin.Context) {
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

	if err := companyService.Delete(uint(id)); err != nil {
		sendError(ctx, http.StatusInternalServerError, "error deleting company")
		return
	}

	sendSuccess(ctx, "delete-company", nil)
}