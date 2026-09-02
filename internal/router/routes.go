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
			protected.GET("/users", middleware.RoleRequired("ADMIN"), user.ListUsersHandler)
			protected.GET("/users/:id", middleware.RoleRequired("ADMIN"), user.ShowUserHandler)
			protected.PUT("/users/:id", middleware.RoleRequired("ADMIN"), user.UpdateUserHandler)
			protected.DELETE("/users/:id", middleware.RoleRequired("ADMIN"), user.DeleteUserHandler)

			protected.GET("/openings", middleware.RoleRequired("COMPANY", "CANDIDATE", "ADMIN"), opening.ListOpeningsHandler)
			protected.GET("/openings/:id", middleware.RoleRequired("COMPANY", "CANDIDATE", "ADMIN"), opening.ShowOpeningHandler)
			protected.POST("/openings", middleware.RoleRequired("COMPANY", "ADMIN"), opening.CreateOpeningHandler)
			protected.PUT("/openings/:id", middleware.RoleRequired("COMPANY", "ADMIN"), opening.UpdateOpeningHandler)
			protected.DELETE("/openings/:id", middleware.RoleRequired("COMPANY", "ADMIN"), opening.DeleteOpeningHandler)

			protected.GET("/companies", middleware.RoleRequired("COMPANY", "CANDIDATE", "ADMIN"), company.ListCompaniesHandler)
			protected.GET("/companies/:id", middleware.RoleRequired("COMPANY", "CANDIDATE", "ADMIN"), company.ShowCompanyHandler)
			protected.POST("/companies", middleware.RoleRequired("COMPANY", "ADMIN"), company.CreateCompanyHandler)
			protected.PUT("/companies/:id", middleware.RoleRequired("COMPANY", "ADMIN"), company.UpdateCompanyHandler)
			protected.DELETE("/companies/:id", middleware.RoleRequired("COMPANY", "ADMIN"), company.DeleteCompanyHandler)

			protected.GET("/candidates", middleware.RoleRequired("CANDIDATE", "ADMIN"), candidate.ListCandidatesHandler)
			protected.GET("/candidates/:id", middleware.RoleRequired("CANDIDATE", "ADMIN"), candidate.ShowCandidateHandler)
			protected.POST("/candidates", middleware.RoleRequired("CANDIDATE", "ADMIN"), candidate.CreateCandidateHandler)
			protected.PUT("/candidates/:id", middleware.RoleRequired("CANDIDATE", "ADMIN"), candidate.UpdateCandidateHandler)
			protected.DELETE("/candidates/:id", middleware.RoleRequired("CANDIDATE", "ADMIN"), candidate.DeleteCandidateHandler)

			protected.GET("/applications", middleware.RoleRequired("CANDIDATE", "ADMIN"), application.ListApplicationsHandler)
			protected.GET("/applications/:id", middleware.RoleRequired("CANDIDATE", "ADMIN"), application.ShowApplicationHandler)
			protected.POST("/applications", middleware.RoleRequired("CANDIDATE"), application.CreateApplicationHandler)
			protected.PUT("/applications/:id", middleware.RoleRequired("ADMIN"), application.UpdateApplicationHandler)
			protected.DELETE("/applications/:id", middleware.RoleRequired("ADMIN"), application.DeleteApplicationHandler)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
