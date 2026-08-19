package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func ListOpeningsHandler(context *gin.Context){

	context.JSON(http.StatusOK, gin.H{
				"message" : "GET Openings",
			})
}
func CreateOpeningHandler(context *gin.Context){

	context.JSON(http.StatusOK, gin.H{
				"message" : "GET Openings",
			})
}
func ShowOpeningHandler(context *gin.Context){

	context.JSON(http.StatusOK, gin.H{
				"message" : "GET Openings",
			})
}
func UpdateOpeningHandler(context *gin.Context){

	context.JSON(http.StatusOK, gin.H{
				"message" : "GET Openings",
			})
}
func DeleteOpeningHandler(context *gin.Context){

	context.JSON(http.StatusOK, gin.H{
				"message" : "GET Openings",
			})
}
