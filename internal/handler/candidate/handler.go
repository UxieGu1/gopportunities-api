package candidate

import (
	"github.com/UxieGu1/gopportunities-api/internal/config"
	"github.com/UxieGu1/gopportunities-api/internal/repository"
	"github.com/UxieGu1/gopportunities-api/internal/service"
	"gorm.io/gorm"
)

var (
	logger           *config.Logger
	db               *gorm.DB
	candidateService service.CandidateService
)

func InitializeHandler() {
	logger = config.GetLogger("candidate_handler")
	db = config.GetSQLite()
	
	candidateRepo := repository.NewCandidateRepository(db)
	candidateService = service.NewCandidateService(candidateRepo)
}