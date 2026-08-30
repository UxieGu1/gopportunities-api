package company

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCompaniesHandler(ctx *gin.Context) {
	responses, err := companyService.GetAll()
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing companies")
		return
	}

	sendSuccess(ctx, "list-companies", responses)
}