package candidate

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


type CreateCandidateResponse struct {
	Message string                    `json:"message"`
	Data    schemas.CandidateResponse `json:"data"`
}

type DeleteCandidateResponse struct {
	Message string                    `json:"message"`
	Data    schemas.CandidateResponse `json:"data"`
}

type ShowCandidateResponse struct {
	Message string                    `json:"message"`
	Data    schemas.CandidateResponse `json:"data"`
}

type ListCandidatesResponse struct {
	Message string                      `json:"message"`
	Data    []schemas.CandidateResponse `json:"data"`
}

type UpdateCandidateResponse struct {
	Message string                    `json:"message"`
	Data    schemas.CandidateResponse `json:"data"`
}