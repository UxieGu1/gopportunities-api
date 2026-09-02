package router

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/UxieGu1/gopportunities-api/internal/config"
	"github.com/UxieGu1/gopportunities-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	if envMode := os.Getenv("GIN_MODE"); envMode != "" {
		gin.SetMode(envMode)
	}

	router := gin.Default()
	router.Use(middleware.CorsMiddleware())
	logger := config.GetLogger("http")
	router.Use(func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = fmt.Sprintf("req-%d", time.Now().UnixNano())
		}

		c.Set("request_id", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)

		start := time.Now()
		c.Next()

		fields := map[string]string{
			"request_id": requestID,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"status":     strconv.Itoa(c.Writer.Status()),
			"latency_ms": strconv.FormatInt(int64(time.Since(start)/time.Millisecond), 10),
		}

		if c.Writer.Status() >= 400 {
			logger.ErrorWithFields("request_failed", fields)
			return
		}

		logger.InfoWithFields("request_completed", fields)
	})

	initializeRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
