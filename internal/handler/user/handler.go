package user

import (
	"github.com/UxieGu1/gopportunities-api/internal/config"
	"github.com/UxieGu1/gopportunities-api/internal/repository"
	"github.com/UxieGu1/gopportunities-api/internal/service"
	"gorm.io/gorm"
)

var (
	logger      *config.Logger
	db          *gorm.DB
	userService service.UserService 
)

func InitializeHandler() {
	logger = config.GetLogger("user_handler")
	db = config.GetSQLite()
	userRepo := repository.NewUserRepository(db)
	userService = service.NewUserService(userRepo)
}