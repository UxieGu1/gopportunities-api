package router


import (
	"log"

	"github.com/gin-gonic/gin"
)

func Initialize() {
  router := gin.Default()

  initilizeRoutes(router)

  router.Run(":8080")

  if err := router.Run(); err != nil {
    log.Fatalf("failed to run server: %v", err)
  }
}