package handler

import (
	"github.com/UxieGu1/gopportunities-api/config"
	"github.com/UxieGu1/gopportunities-api/service"
	"gorm.io/gorm"
)


var (
	logger *config.Logger
	db *gorm.DB
	openingService *service.OpeningService
)

func InitializeHandler(){
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
	openingService = service.NewOpeningService(db)
}



