package application

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

type CreateApplicationResponse struct {
	Message string                      `json:"message"`
	Data    schemas.ApplicationResponse `json:"data"`
}

type DeleteApplicationResponse struct {
	Message string                      `json:"message"`
	Data    schemas.ApplicationResponse `json:"data"`
}

type ShowApplicationResponse struct {
	Message string                      `json:"message"`
	Data    schemas.ApplicationResponse `json:"data"`
}

type ListApplicationsResponse struct {
	Message string                        `json:"message"`
	Data    []schemas.ApplicationResponse `json:"data"`
}

type UpdateApplicationResponse struct {
	Message string                      `json:"message"`
	Data    schemas.ApplicationResponse `json:"data"`
}