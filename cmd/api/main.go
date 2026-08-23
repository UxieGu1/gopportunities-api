// @title Gopportunities API
// @version 1.0
// @description API para gerenciar oportunidades de emprego.
// @host localhost:8081
// @BasePath /api/v1
package main

import (
	"github.com/UxieGu1/gopportunities-api/internal/config"
	"github.com/UxieGu1/gopportunities-api/internal/handler"
	"github.com/UxieGu1/gopportunities-api/internal/router"
	"github.com/joho/godotenv"
)

var (
	logger *config.Logger
)

func main() {
	_ = godotenv.Load()

	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	handler.InitializeHandler()

	router.Initialize()
}
