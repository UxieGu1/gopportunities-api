package company

import (
	"github.com/UxieGu1/gopportunities-api/internal/config"
	"github.com/UxieGu1/gopportunities-api/internal/service/company"
	"gorm.io/gorm"
)

var (
	logger         *config.Logger
	db             *gorm.DB
	companyService service.CompanyService 
)

func InitializeHandler() {
	logger = config.GetLogger("company_handler")
	db = config.GetSQLite()
	// Inicialize o repositório e injete no service
	// companyRepo := repository.NewCompanyRepository(db)
	// companyService = service.NewCompanyService(companyRepo)
}