package company

import (
	"fmt"
	"net/http"

	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(context *gin.Context, code int, message string) {
	context.Header("Content-type", "application/json")
	context.JSON(code, gin.H{
		"message":   message,
		"errorCode": code,
	})
}

func sendSuccess(context *gin.Context, op string, data interface{}) {
	context.Header("Content-type", "application/json")
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s sucessfully", op),
		"data":    data,
	})
}

type CreateCompanyResponse struct {
	Message string                  `json:"message"`
	Data    schemas.CompanyResponse `json:"data"`
}

type DeleteCompanyResponse struct {
	Message string                  `json:"message"`
	Data    schemas.CompanyResponse `json:"data"`
}

type ShowCompanyResponse struct {
	Message string                  `json:"message"`
	Data    schemas.CompanyResponse `json:"data"`
}

type ListCompaniesResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.CompanyResponse `json:"data"`
}

type UpdateCompanyResponse struct {
	Message string                  `json:"message"`
	Data    schemas.CompanyResponse `json:"data"`
}