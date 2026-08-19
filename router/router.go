package router


import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
  router := gin.Default()

  router.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })


  if err := router.Run(); err != nil {
    log.Fatalf("failed to run server: %v", err)
  }
}