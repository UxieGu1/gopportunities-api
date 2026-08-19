package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(context *gin.Context){

	context.JSON(http.StatusOK, gin.H{
				"message" : "GET Openings",
			})
}