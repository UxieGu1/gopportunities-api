package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)



func sendError(context *gin.Context, code int, message string){
	context.Header("Content-type", "application/json")
	context.JSON(code, gin.H{
		"message": message,
		"errorCode" : code,
	})
}


func sendSucess(context *gin.Context, op string, data interface{}){
	context.Header("Content-type", "application/json")
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s sucessfully", op),
		"data" : data,
	})
}