package opening

import (
	"github.com/UxieGu1/gopportunities-api/internal/config"
	"github.com/UxieGu1/gopportunities-api/internal/repository"
	"github.com/UxieGu1/gopportunities-api/internal/service"
	"gorm.io/gorm"
)

var (
	logger         *config.Logger
	db             *gorm.DB
	openingService service.OpeningService 
)

func InitializeHandler() {
	logger = config.GetLogger("opening_handler")
	db = config.GetSQLite()
	
	openingRepo := repository.NewOpeningRepository(db)
	openingService = service.NewOpeningService(openingRepo)
}