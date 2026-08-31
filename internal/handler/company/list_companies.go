package company

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCompaniesHandler(context *gin.Context) {
	responses, err := companyService.GetAll()
	if err != nil {
		sendError(context, http.StatusInternalServerError, "error listing companies")
		return
	}

	sendSuccess(context, "list-companies", responses)
}