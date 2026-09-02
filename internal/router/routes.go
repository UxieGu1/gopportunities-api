package router

import (
	docs "github.com/UxieGu1/gopportunities-api/docs"
	"github.com/UxieGu1/gopportunities-api/internal/handler/application"
	"github.com/UxieGu1/gopportunities-api/internal/handler/candidate"
	"github.com/UxieGu1/gopportunities-api/internal/handler/company"
	"github.com/UxieGu1/gopportunities-api/internal/handler/opening"
	"github.com/UxieGu1/gopportunities-api/internal/handler/user"
	"github.com/UxieGu1/gopportunities-api/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	opening.InitializeHandler()
	company.InitializeHandler()
	user.InitializeHandler()
	candidate.InitializeHandler()
	application.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.POST("/login", user.LoginUserHandler)
		v1.POST("/register", user.RegisterUserHandler)

		protected := v1.Group("/")
		protected.Use(middleware.AuthRequired())
		{
			protected.GET("/users", user.ListUsersHandler)
			protected.GET("/users/:id", user.ShowUserHandler)
			protected.PUT("/users/:id", user.UpdateUserHandler)
			protected.DELETE("/users/:id", user.DeleteUserHandler)

			protected.GET("/openings", opening.ListOpeningsHandler)
			protected.GET("/openings/:id", opening.ShowOpeningHandler)
			protected.POST("/openings", middleware.RoleRequired("COMPANY", "ADMIN"), opening.CreateOpeningHandler)
			protected.PUT("/openings/:id", middleware.RoleRequired("COMPANY", "ADMIN"), opening.UpdateOpeningHandler)
			protected.DELETE("/openings/:id", middleware.RoleRequired("COMPANY", "ADMIN"), opening.DeleteOpeningHandler)
        
			protected.GET("/companies", company.ListCompaniesHandler)
			protected.GET("/companies/:id", company.ShowCompanyHandler)
			protected.POST("/companies", middleware.RoleRequired("COMPANY", "ADMIN"), company.CreateCompanyHandler)
			protected.PUT("/companies/:id", middleware.RoleRequired("COMPANY", "ADMIN"), company.UpdateCompanyHandler)
			protected.DELETE("/companies/:id", middleware.RoleRequired("COMPANY", "ADMIN"), company.DeleteCompanyHandler)

			protected.GET("/candidates", candidate.ListCandidatesHandler)
			protected.GET("/candidates/:id", candidate.ShowCandidateHandler)
			protected.POST("/candidates", middleware.RoleRequired("CANDIDATE", "ADMIN"), candidate.CreateCandidateHandler)
			protected.PUT("/candidates/:id", middleware.RoleRequired("CANDIDATE", "ADMIN"), candidate.UpdateCandidateHandler)
			protected.DELETE("/candidates/:id", middleware.RoleRequired("CANDIDATE", "ADMIN"), candidate.DeleteCandidateHandler)

			protected.GET("/applications", application.ListApplicationsHandler)
			protected.GET("/applications/:id", application.ShowApplicationHandler)
			protected.POST("/applications", middleware.RoleRequired("CANDIDATE"), application.CreateApplicationHandler)
			protected.PUT("/applications/:id", middleware.RoleRequired("COMPANY", "ADMIN"), application.UpdateApplicationHandler)
			protected.DELETE("/applications/:id", middleware.RoleRequired("ADMIN"), application.DeleteApplicationHandler)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}