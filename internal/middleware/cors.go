package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     configuredOrigins(),
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Request-ID"},
		ExposeHeaders:    []string{"Content-Length", "X-Request-ID"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func configuredOrigins() []string {
	origins := strings.Split(os.Getenv("CORS_ORIGINS"), ",")
	for index := range origins {
		origins[index] = strings.TrimSpace(origins[index])
	}
	if len(origins) == 1 && origins[0] == "" {
		return []string{"http://localhost:3000", "http://localhost:5173"}
	}
	return origins
}
