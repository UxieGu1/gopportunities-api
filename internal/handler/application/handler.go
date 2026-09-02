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
	candidateService   service.CandidateService
)

func InitializeHandler() {
	logger = config.GetLogger("application_handler")
	db = config.GetSQLite()

	applicationRepo := repository.NewApplicationRepository(db)
	applicationService = service.NewApplicationService(applicationRepo)
	candidateRepo := repository.NewCandidateRepository(db)
	candidateService = service.NewCandidateService(candidateRepo)
}
