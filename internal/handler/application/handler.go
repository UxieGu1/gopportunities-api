package application

import (
	"github.com/UxieGu1/gopportunities-api/internal/config"
	"github.com/UxieGu1/gopportunities-api/internal/repository"
	"github.com/UxieGu1/gopportunities-api/internal/service"
	"gorm.io/gorm"
)

var (
	logger             *config.Logger
	db                 *gorm.DB
	applicationService service.ApplicationService
)

func InitializeHandler() {
	logger = config.GetLogger("application_handler")
	db = config.GetSQLite()
	
	applicationRepo := repository.NewApplicationRepository(db)
	applicationService = service.NewApplicationService(applicationRepo)
}