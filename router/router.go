package router

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
  router := gin.Default()

  initilizeRoutes(router)

  port := os.Getenv("PORT")
  if port == "" {
    port = "8081"
  }

  if err := router.Run(":" + port); err != nil {
    log.Fatalf("failed to run server: %v", err)
  }
}